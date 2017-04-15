package main

import (
	"time"

	"github.com/cenkalti/backoff"
	"github.com/urfave/cli"
	"tantalic.com/anybar"
	"tantalic.com/cistatus"
)

const anybarLaunchWaitDuration = 1500 * time.Millisecond

func runloop(conf config) error {
	logger := conf.Logger()

	anybarClient := anybar.Client{
		Hostname: conf.AnyBarHost,
		Port:     conf.AnyBarPort,
	}

	logger.Printf("starting cistatusanybar version %s\n", cistatus.Version)

	if conf.StartAnyBar {
		logger.Println("starting anybar app")
		err := anybarClient.Start()
		if err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
		time.Sleep(anybarLaunchWaitDuration)
	}

	logger.Printf("setting status to %s until first status is received\n", anybar.Question)
	anybarClient.Set(anybar.Question)

	summaryChan := make(chan cistatus.Summary)
	go func() {
		for {
			status := <-summaryChan

			logger.Printf("received status: %s\n", status)
			anybarClient.Set(anybar.Style(status.Color))
		}
	}()

	client := cistatus.Client{
		Hostname: conf.StatusHost,
		Port:     conf.StatusPort,
	}

	operation := func() error {
		logger.Printf("subscribing to watch status (%s)\n", client.Hostname)
		err := client.Watch(summaryChan)
		if err != nil {
			logger.Printf("error watching status:%v\n", err)
			anybarClient.Set(anybar.Question)
		}
		return err
	}

	expBackoff := backoff.ExponentialBackOff{
		InitialInterval:     1 * time.Second,
		RandomizationFactor: 0.3,
		Multiplier:          1.5,
		MaxInterval:         5 * time.Minute,
		MaxElapsedTime:      0,
		Clock:               backoff.SystemClock,
	}
	err := backoff.Retry(operation, &expBackoff)
	return err
}

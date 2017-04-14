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

	logger.Logf("starting cistatusanybar version %s", cistatus.Version)

	if conf.StartAnyBar {
		logger.Log("starting anybar app")
		err := anybarClient.Start()
		if err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
		time.Sleep(anybarLaunchWaitDuration)
	}

	logger.Logf("setting status to %s until first status is received", anybar.Question)
	anybarClient.Set(anybar.Question)

	summaryChan := make(chan cistatus.Summary)
	go func() {
		for {
			status := <-summaryChan

			logger.Logf("received status: %s", status)
			anybarClient.Set(anybar.Style(status.Color))
		}
	}()

	client := cistatus.Client{
		Hostname: conf.StatusHost,
		Port:     conf.StatusPort,
	}

	operation := func() error {
		logger.Logf("subscribing to watch status (%s)", client.Hostname)
		err := client.Watch(summaryChan)
		if err != nil {
			logger.Logf("error watching status:%v", err)
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

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/urfave/cli"

	"gobot.io/x/gobot"
	"tantalic.com/cistatus"
)

func main() {
	app := cli.App{
		Name:         filepath.Base(os.Args[0]),
		HelpName:     filepath.Base(os.Args[0]),
		Usage:        "changes the ",
		UsageText:    "cistatuslight [options] [ci-status-server]",
		Version:      cistatus.Version,
		BashComplete: cli.DefaultAppComplete,
		Writer:       os.Stdout,
	}

	var conf config
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "p,port",
			EnvVar:      "PORT",
			Usage:       "http(s) port to connect to the ci status server on",
			Destination: &conf.StatusPort,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "print status information",
			EnvVar:      "VERBOSE",
			Destination: &conf.Verbose,
		},
		cli.IntFlag{
			Name:        "r,red-pin",
			EnvVar:      "RED_PIN",
			Usage:       "gpio for red pin",
			Destination: &conf.RedPin,
		},
		cli.IntFlag{
			Name:        "y,yello-pin",
			EnvVar:      "YELLOW_PIN",
			Usage:       "gpio for yellow pin",
			Destination: &conf.YellowPin,
		},
		cli.IntFlag{
			Name:        "g,green-pin",
			EnvVar:      "GREEN_PIN",
			Usage:       "gpio for green pin",
			Destination: &conf.GreenPin,
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) != 1 {
			err := cli.NewExitError("status server hostname argument must be provided", 1)
			printUsage(c)
			return err
		}

		conf.StatusHost = c.Args().Get(0)
		return run(conf)
	}

	app.Run(os.Args)
}

func printUsage(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "usage: %s\n", c.App.UsageText)
}

func run(conf config) error {
	summaryChan := make(chan cistatus.Summary)

	go func() {
		client := conf.CIStatusClient()

		operation := func() error {
			log.Println("Subscribing to watch status")
			err := client.Watch(summaryChan)
			if err != nil {
				log.Printf("Error watching status:%v\n", err)
				summaryChan <- cistatus.Summary{
					Color: cistatus.Unknown,
				}
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
		backoff.Retry(operation, &expBackoff)
	}()

	robot := NewStatusLightRobot(conf, summaryChan)
	return robot.Start()
}

func NewStatusLightRobot(c config, summaryChan chan cistatus.Summary) *gobot.Robot {
	pi := c.Adapter()
	red := c.RedPinDriver()
	yellow := c.YellowPinDriver()
	green := c.GreenPinDriver()

	work := func() {
		for {
			summary := <-summaryChan
			log.Printf("Received status update: %s\n", summary.Color)

			if summary.Color == cistatus.Red {
				red.Off()
			} else {
				red.On()
			}

			if summary.Color == cistatus.Yellow {
				yellow.Off()
			} else {
				yellow.On()
			}

			if summary.Color == cistatus.Green {
				green.Off()
			} else {
				green.On()
			}
		}
	}

	robot := gobot.NewRobot("cistatuslight",
		[]gobot.Connection{pi},
		[]gobot.Device{
			red,
			yellow,
			green,
		},
		work,
	)

	return robot
}

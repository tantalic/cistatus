package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"tantalic.com/cistatus"
)

func main() {
	app := cli.App{
		Name:         filepath.Base(os.Args[0]),
		HelpName:     filepath.Base(os.Args[0]),
		Usage:        "updates the anybar status icon based on the status of the continuous integration server",
		UsageText:    "cistatusanybar [options] [ci-hostname]",
		Version:      cistatus.Version,
		BashComplete: cli.DefaultAppComplete,
		Writer:       os.Stdout,
	}

	var conf config
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "status-port",
			EnvVar:      "STATUS_PORT",
			Usage:       "tcp (http) port to connect to the status server on",
			Destination: &conf.StatusPort,
		},
		cli.StringFlag{
			Name:        "anybar-host",
			EnvVar:      "ANYBAR_HOST",
			Usage:       "host that is running the anybar app",
			Destination: &conf.AnyBarHost,
		},
		cli.IntFlag{
			Name:        "anybar-port",
			Value:       1739,
			EnvVar:      "ANYBAR_PORT",
			Usage:       "udp port the anybar app is listening on",
			Destination: &conf.AnyBarPort,
		},
		cli.BoolFlag{
			Name:        "start-anybar",
			EnvVar:      "START_ANYBAR",
			Usage:       "start (or restart) the anybar app",
			Destination: &conf.StartAnyBar,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "print status information",
			EnvVar:      "VERBOSE",
			Destination: &conf.Verbose,
		},
		cli.BoolFlag{
			Name:        "install-launch-agent",
			Usage:       "install macos launchd agent to launch on startup and keep running",
			Destination: &conf.InstallLaunchAgent,
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) != 1 {
			err := cli.NewExitError("status server hostname argument must be provided", 1)
			printUsage(c)
			return err
		}

		conf.StatusHost = c.Args().Get(0)

		if conf.InstallLaunchAgent {
			return installLaunchAgent(conf)
		}

		return runloop(conf)
	}

	app.Run(os.Args)
}

func printUsage(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "usage: %s\n", c.App.UsageText)
}

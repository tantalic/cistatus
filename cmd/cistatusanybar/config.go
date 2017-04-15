package main

import (
	"os"

	"tantalic.com/cistatus"
)

type config struct {
	StatusHost         string
	StatusPort         int
	AnyBarHost         string
	AnyBarPort         int
	StartAnyBar        bool
	Verbose            bool
	InstallLaunchAgent bool
}

func (c config) Logger() cistatus.Logger {
	var logger cistatus.Logger
	logger = cistatus.NewNullLogger()

	if c.Verbose {
		logger = cistatus.NewVerboseLogger(os.Stdout)
	}

	return logger
}

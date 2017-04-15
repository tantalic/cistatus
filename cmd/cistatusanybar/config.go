package main

import (
	"io/ioutil"
	"log"
	"os"
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

func (c config) Logger() *log.Logger {
	if !c.Verbose {
		return log.New(ioutil.Discard, "", 0)
	}

	return log.New(os.Stdout, "", log.LstdFlags)
}

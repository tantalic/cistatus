package main

import (
	"strconv"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"

	"tantalic.com/cistatus"
)

const (
	DefaultRedPin    = 15
	DefaultYellowPin = 12
	DefaultGreenPin  = 11
)

type config struct {
	StatusHost string
	StatusPort int

	Verbose bool

	RedPin    int
	YellowPin int
	GreenPin  int
	adaptor   gobot.Connection
}

func (c config) CIStatusClient() *cistatus.Client {
	return &cistatus.Client{
		Hostname: c.StatusHost,
		Port:     c.StatusPort,
	}
}

func (c config) Adapter() gobot.Connection {
	if c.adaptor == nil {
		c.adaptor = raspi.NewAdaptor()
	}

	return c.adaptor
}

func (c config) RedPinDriver() *gpio.DirectPinDriver {
	return c.newPinDriver(c.RedPin, DefaultRedPin)
}

func (c config) YellowPinDriver() *gpio.DirectPinDriver {
	return c.newPinDriver(c.YellowPin, DefaultYellowPin)
}

func (c config) GreenPinDriver() *gpio.DirectPinDriver {
	return c.newPinDriver(c.GreenPin, DefaultGreenPin)
}

func (c config) newPinDriver(pin, defaultPin int) *gpio.DirectPinDriver {
	if pin == 0 {
		pin = defaultPin
	}

	return gpio.NewDirectPinDriver(c.Adapter(), strconv.Itoa(pin))
}

package dragonboard

import (
	"errors"
	"testing"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/gobottest"
	"gobot.io/x/gobot/sysfs"
)

// make sure that this Adaptor fullfills all the required interfaces
var _ gobot.Adaptor = (*Adaptor)(nil)
var _ gpio.DigitalReader = (*Adaptor)(nil)
var _ gpio.DigitalWriter = (*Adaptor)(nil)
var _ i2c.Connector = (*Adaptor)(nil)

type NullReadWriteCloser struct {
	contents []byte
}

func (n *NullReadWriteCloser) SetAddress(int) error {
	return nil
}

func (n *NullReadWriteCloser) Write(b []byte) (int, error) {
	n.contents = make([]byte, len(b))
	copy(n.contents[:], b[:])

	return len(b), nil
}

func (n *NullReadWriteCloser) Read(b []byte) (int, error) {
	copy(b, n.contents)
	return len(b), nil
}

var closeErr error

func (n *NullReadWriteCloser) Close() error {
	return closeErr
}

func initTestDragonBoardAdaptor(t *testing.T) *Adaptor {
	a := NewAdaptor()
	if err := a.Connect(); err != nil {
		t.Error(err)
	}
	return a
}

func TestDragonBoardAdaptorDigitalIO(t *testing.T) {
	a := initTestDragonBoardAdaptor(t)
	fs := sysfs.NewMockFilesystem([]string{
		"/sys/class/gpio/export",
		"/sys/class/gpio/unexport",
		"/sys/class/gpio/gpio36/value",
		"/sys/class/gpio/gpio36/direction",
		"/sys/class/gpio/gpio12/value",
		"/sys/class/gpio/gpio12/direction",
	})

	sysfs.SetFilesystem(fs)

	_ = a.DigitalWrite("GPIO_B", 1)
	gobottest.Assert(t, fs.Files["/sys/class/gpio/gpio12/value"].Contents, "1")

	fs.Files["/sys/class/gpio/gpio36/value"].Contents = "1"
	i, _ := a.DigitalRead("GPIO_A")
	gobottest.Assert(t, i, 1)

	gobottest.Assert(t, a.DigitalWrite("GPIO_M", 1), errors.New("Not a valid pin"))
}

func TestDragonBoardAdaptorI2c(t *testing.T) {
	a := initTestDragonBoardAdaptor(t)
	fs := sysfs.NewMockFilesystem([]string{
		"/dev/i2c-1",
	})
	sysfs.SetFilesystem(fs)
	sysfs.SetSyscall(&sysfs.MockSyscall{})

	con, err := a.GetConnection(0xff, 1)
	gobottest.Assert(t, err, nil)

	_, _ = con.Write([]byte{0x00, 0x01})
	data := []byte{42, 42}
	_, _ = con.Read(data)
	gobottest.Assert(t, data, []byte{0x00, 0x01})

	gobottest.Assert(t, a.Finalize(), nil)
}

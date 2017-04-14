package cistatus

import (
	"io"
	"log"
)

type Logger interface {
	Log(a ...interface{})
	Logf(format string, a ...interface{})
}

type VerboseLogger struct {
	Logger *log.Logger
}

func NewVerboseLogger(out io.Writer) VerboseLogger {
	return VerboseLogger{
		Logger: log.New(out, "", log.LstdFlags),
	}
}

func (l VerboseLogger) Log(a ...interface{}) {
	l.Logger.Println(a...)
}

func (l VerboseLogger) Logf(format string, a ...interface{}) {
	l.Logger.Printf(format+"\n", a...)
}

type NullLogger struct{}

func (l NullLogger) Log(a ...interface{}) {
	// No-op
}

func (l NullLogger) Logf(format string, a ...interface{}) {
	// No-op
}

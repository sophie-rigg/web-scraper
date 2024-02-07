package logging

import (
	"github.com/sirupsen/logrus"
)

type LogLevel string

func (l *LogLevel) Set(s string) error {
	level, err := logrus.ParseLevel(s)
	if err != nil {
		return err
	}
	// Set the global log level from flag
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	*l = LogLevel(s)
	return nil
}

func (l *LogLevel) Value() string {
	return l.String()
}

func (l *LogLevel) String() string {
	return string(*l)
}

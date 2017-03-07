package main

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"os"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("error", errors.New("log an error")).Warn("error setting log level, using debug as default")

}

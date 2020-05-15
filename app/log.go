package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

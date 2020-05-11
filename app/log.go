package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.NewEntry(logrus.New())
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

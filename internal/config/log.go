package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func InitLogger() error {
	dir, _ := os.Getwd()
	logDir := fmt.Sprintf("%s/logs", dir)

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			return err
		}
	}

	logFile := fmt.Sprintf("%s/logs_%d.log", logDir, time.Now().Unix())
	log, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	logrus.SetOutput(log)

	return nil
}

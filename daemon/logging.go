package daemon

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func StartLoggingToFile(logFileName string) *os.File {
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	return file
}

package log

import (
	"fmt"
	"os"

	"github.com/MSFT/internal/cfg"
	log "github.com/sirupsen/logrus"
)

var ContextLogger *log.Entry

func InitLogger(cfg *cfg.Config) {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	ContextLogger = log.WithFields(log.Fields{
		"service": cfg.Current_service,
	})

	if cfg.Logging_in_file {
		logger_file, err := os.OpenFile(fmt.Sprintf("logger/%v.log", cfg.Current_service), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("failed to create or opening the logger file:\n" + err.Error())
		}
		defer logger_file.Close()
		log.SetOutput(logger_file)
	}
}

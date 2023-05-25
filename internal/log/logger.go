package log

import (
	"fmt"
	"io"
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

	logger_file, err := os.OpenFile(fmt.Sprintf("logger/%v.log", cfg.Current_service), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to create or opening the logger file:\n" + err.Error())
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logger_file))
}

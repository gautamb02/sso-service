package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gautamb02/sso-service/confreader"
)

var Log *log.Logger

func InitLogger(cfg confreader.LoggerConfig) error {
	logFile := cfg.File
	dir := filepath.Dir(logFile)
	if dir != "." {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Info("Logger initialized")
	return nil
}

// Level-specific helpers

func Info(format string, v ...any) {
	Log.Output(2, "[INFO] "+fmt.Sprintf(format, v...))
}

func Warn(format string, v ...any) {
	Log.Output(2, "[WARN] "+fmt.Sprintf(format, v...))
}

func Error(format string, v ...any) {
	Log.Output(2, "[ERROR] "+fmt.Sprintf(format, v...))
}

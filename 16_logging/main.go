package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	// prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type myFormatter struct {
	log.TextFormatter
}

func (f *myFormatter) Format(entry *log.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 31 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func main() {
	f, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	logger := &log.Logger{
		Out:   io.MultiWriter(os.Stderr, f),
		Level: log.InfoLevel,
		Formatter: &myFormatter{log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
		},
		// formatter := &myFormatter{
		// 	log.TextFormatter{
		// 		FullTimestamp:          true,
		// 		TimestampFormat:        "2006-01-02 15:04:05",
		// 		ForceColors:            true,
		// 		DisableLevelTruncation: true,
		// 	},
		// }
		// log.SetLevel(log.DebugLevel)
		// log.SetOutput(io.MultiWriter(os.Stdout, logFile))
		// log.SetFormatter(formatter)

	}
	logger.Info("Info message")
	logger.Warning("Warning message")
	// fmt.Printf("\x1b[%dm%s\x1b[0m line2\n", 31, "Hiii")
}

package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func InitLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuration
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// Custom format for the file path and line number
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logger.SetFormatter(formatter)
	return logger
}

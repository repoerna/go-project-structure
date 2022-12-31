package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)

	// logrus.WithField("app", os.Getenv("APP"))

	logrus.WithFields(
		logrus.Fields{
			"app": os.Getenv("APP"),
		},
	)

	textFormatter := &logrus.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuration
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
		PadLevelText: true,
	}

	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}

	if os.Getenv("ENV") == "production" {
		logrus.SetFormatter(jsonFormatter)

	} else {
		logrus.SetFormatter(textFormatter)
		logrus.SetLevel(logrus.DebugLevel)
	}

}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

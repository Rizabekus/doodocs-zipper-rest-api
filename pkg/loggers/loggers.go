package loggers

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var (
	Log1 *logrus.Logger
	Log2 *logrus.Logger
	Log3 *logrus.Logger
	f    *os.File
)

func InitLoggers() {
	f, err := os.OpenFile("logs/all.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	Log1 = &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%] [%time%] (%file%:%line%) - %msg%\n",
		},
	}

	Log2 = &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.ErrorLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%] [%method%] [%time%] (%file%:%line%) - %msg% | Path: %path% | Status Code: %status_code% | IP Address: %ip_address% | Content Type: %content_type% | User Agent: %user_agent% | Error: %error% \n",
		},
	}

	Log3 = &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%] [%method%] [%time%] (%file%:%line%) - %msg% | Path: %path% | Status Code: %status_code% | IP Address: %ip_address% | Content Type: %content_type% | User Agent: %user_agent% \n",
		},
	}
}

func InfoLog(file string, line int, msg string) {
	tmp := strings.Split(file, ":")
	var fsplit string
	if len(tmp) == 1 {
		fsplit = tmp[0]
	} else {
		fsplit = tmp[1]
	}

	split := strings.Split(fsplit, "/")
	file = split[len(split)-1]
	Log1.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Info(msg)
}

func DebugLog(file string, line int, method string, path string, status_code int, ip_address string, content_type string, user_agent string, msg string) {
	split := strings.Split(strings.Split(file, ":")[1], "/")
	log.Println(file)
	file = split[len(split)-1]
	content_type = strings.Split(content_type, ";")[0]
	Log3.WithFields(logrus.Fields{
		"file":         file,
		"line":         line,
		"method":       method,
		"path":         path,
		"status_code":  status_code,
		"ip_address":   ip_address,
		"content_type": content_type,
		"user_agent":   user_agent,
	}).Debug(msg)
}

func ErrorLog(file string, line int, method string, path string, status_code int, ip_address string, content_type string, user_agent string, err string, msg string) {
	split := strings.Split(strings.Split(file, ":")[1], "/")
	file = split[len(split)-1]
	content_type = strings.Split(content_type, ";")[0]
	Log2.WithFields(logrus.Fields{
		"file":         file,
		"line":         line,
		"method":       method,
		"path":         path,
		"status_code":  status_code,
		"ip_address":   ip_address,
		"content_type": content_type,
		"user_agent":   user_agent,
		"error":        err,
	}).Error(msg)
}

func CloseLogFile() error {
	if f != nil {
		err := f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

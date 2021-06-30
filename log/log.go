package log

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

func Info(message interface{}) {
	fmt.Printf("%v [ INFO] - %v", time.Now().Format("2006-01-02 15:04:05.000"), message)
}

func Infof(format string, a ...interface{}) {
	fmt.Printf("%v [ INFO] - %v", time.Now().Format("2006-01-02 15:04:05.000"), fmt.Sprintf(format, a...))
}

func Error(message interface{}) {
	fmt.Printf("%v [ INFO] - %v", time.Now().Format("2006-01-02 15:04:05.000"), message)
}

func Errorf(format string, a ...interface{}) {
	fmt.Printf("%v [ INFO] - %v", time.Now().Format("2006-01-02 15:04:05.000"), fmt.Sprintf(format, a...))
}

func Debug(message interface{}) {
	fmt.Printf("%v [ INFO] - %v", time.Now().Format("2006-01-02 15:04:05.000"), message)
}

func Debugf(format string, a ...interface{}) {
	fmt.Printf("%v [ INFO] - %v", time.Now().Format("2006-01-02 15:04:05.000"), fmt.Sprintf(format, a...))
}

func Panic(message ...interface{}) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Exception: ", log.Lshortfile)
	)
	logger.Print(message...)
	fmt.Printf("%v [PANIC] - %v", time.Now().Format("2006-01-02 15:04:05.000"), &buf)
}

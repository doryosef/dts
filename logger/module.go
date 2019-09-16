package logger

import (
    "log"
    "os"
    "sync"
)

type logger struct {
    filename string
    *log.Logger
}

var myLogger *logger

var once sync.Once

func GetInstance() *logger {
    once.Do(func() {
        myLogger = createLogger("dts.log")
    })
    return myLogger
}

func createLogger(fname string) *logger {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
    return &logger{
        filename: fname,
        Logger:   log.New(file,"", log.Ldate | log.Ltime),
    }
}
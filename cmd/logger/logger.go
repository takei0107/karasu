package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "[INFO]", log.LstdFlags)
	Error = log.New(os.Stderr, "[ERROR]", log.LstdFlags)
}

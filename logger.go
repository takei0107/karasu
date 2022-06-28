package karasu

import (
	"log"
	"os"
)

var (
	info_logger *log.Logger
	err_logger  *log.Logger
)

func init() {
	info_logger = log.New(os.Stdout, "[INFO]", log.LstdFlags)
	err_logger = log.New(os.Stderr, "[ERROR]", log.LstdFlags)
}

func log_info(message string) {
	info_logger.Println(message)
}

func log_error(message string, err error) {
	err_logger.Printf("%s\n%s", message, err.Error())
}

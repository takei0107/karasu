package main

import (
	"github.com/takei0107/karasu"
	"github.com/takei0107/karasu/cmd/logger"
)

func main() {
	if err := karasu.Run(); err != nil {
		logger.Error.Fatal(err)
	}
}

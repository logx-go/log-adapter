package main

import (
	"log"
	"os"

	"github.com/logx-go/contract/pkg/logx"
	"github.com/logx-go/log-adapter/pkg/logadapter"
)

func main() {
	// we are using the standard adapter from https://pkg.go.dev/log
	logger := logadapter.New(log.New(os.Stdout, "", 0))

	logSomething(logger)
}

func logSomething(logger logx.Logger) {
	logger.Info("This is log message")
}

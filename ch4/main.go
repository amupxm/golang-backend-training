package main

import (
	arg "github.com/amupxm/golang-backend-training/c4/args"
	logger "github.com/amupxm/xmus-logger"
)

func main() {
	argLogLever := arg.LogLevel
	LogOptions := logger.LoggerOptions{
		LogLevel: logger.LogLevel(argLogLever),
		Verbose:  true,
		Std:      true,
	}

	log := logger.CreateLogger(&LogOptions)
	log.Inform("Hello World!\n")
	log.End()
}

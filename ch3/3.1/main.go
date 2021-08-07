package main

import (
	"flag"

	logger "github.com/amupxm/golang-backend-training/c3/3.1/logger/srv"
)

func main() {
	flag.Parse()
	logger := logger.CreateLogger(
		&logger.LoggerOptions{
			LogLevel:    0,
			Verbose:     false,
			File:        true,
			FilePath:    "string",
			Std:         true,
			UseCollores: true,
		},
	)
	logger.Informln(1)
	logger.Highlightln(2)
	logger.Log(1)
	logger.Logln(1, ",2323")
	logger.LogF("%s \n%d\n", "mmd", 1)
	pref := logger.Prefix("here", "there")
	pref.GetCaller().Log(2221)
	pref.Warnln("warn")

	logger.End()
}

package main

func main() {
	logger := CreateLogger(
		&LoggerOptions{
			LogLevel:    0,
			Verbose:     true,
			File:        true,
			FilePath:    "string",
			Std:         true,
			UseCollores: true,
		},
	)
	// logger.Begin()
	logger.Log(1)
	logger.Logln(1, ",2323")
	logger.LogF("%s \n%d\n", "mmd", 1)
	pref := logger.Prefix("here", "there")
	pref.GetCaller().Log(2221)
	pref.Warnln("warn")

	logger.End()
}
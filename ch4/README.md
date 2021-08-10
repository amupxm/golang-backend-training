### Chapter 4️⃣
 
You can check logger repo at [HERE](https://github.com/amupxm/xmus-logger/) and docs at [golang-docs](https://pkg.go.dev/github.com/amupxm/xmus-logger@v0.0.0-20210809175243-b2862ebd67e5)

you can use logger as it describes in main.go 

```go

LogOptions := logger.LoggerOptions{
	LogLevel: logger.LogLevel(6),
	Verbose:  true,
	Std:      true,

log := logger.CreateLogger(&LogOptions)
log.Inform("Hello World!\n")
log.End()


```
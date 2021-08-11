package main

import (
	"net/http"

	"github.com/amupxm/golang-backend-training/ch13/13.2/src/middleware"
	logger "github.com/amupxm/xmus-logger"
)

type (
	logController struct{}
	LogController interface {
		GetLog(w http.ResponseWriter, r *http.Request)
	}
)

var log logger.Logger

func main() {
	log = logger.CreateLogger(
		&logger.LoggerOptions{
			LogLevel: logger.Inform,
			Std:      true,
		},
	)
	http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Informln("New request received")
			log.InformF("BeforUpgrade : Request method is ( %s )\n", r.Method)
			middleware.NewMiddleWare().LogUserAgent(
				w, r, NewLogController().GetLog,
			)
		}))
}

func NewLogController() LogController {
	return &logController{}
}

func (lc *logController) GetLog(w http.ResponseWriter, r *http.Request) {
	log.Prefix("Handler", "GetLog").InformF("Request M-Method is ( %s )\n", r.Method)

	w.Write([]byte("hello"))
}

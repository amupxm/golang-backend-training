package main

import (
	"net/http"

	logger "github.com/amupxm/xmus-logger"
)

type (
	logController struct{}
	LogController interface {
		GetLog(w http.ResponseWriter, r *http.Request)
	}
	middleWare struct{}
	MiddleWare interface {
		SetXMethodHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, method string)
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
			NewMiddleWare().SetXMethodHeader(w, r, NewLogController().GetLog, "get")
		}))
}

func NewLogController() LogController {
	return &logController{}
}
func NewMiddleWare() MiddleWare {
	return &middleWare{}
}

func (m *middleWare) SetXMethodHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, method string) {
	log.Prefix("Middlewar", "X-MethodUpgrader").InformF("Request M-Method is ( %s )\n", r.Header.Get("X-Method"))
	if r.Header.Get("X-Method") == "" {
		r.Header.Set("X-Method", method)
	}
	next(w, r)
}

func (lc *logController) GetLog(w http.ResponseWriter, r *http.Request) {
	log.Prefix("Handler", "GetLog").InformF("Request M-Method is ( %s )\n", r.Header.Get("X-Method"))

	w.Write([]byte("hello"))
}

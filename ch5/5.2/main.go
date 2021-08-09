package main

import (
	"net/http"

	logger "github.com/amupxm/xmus-logger/srv"
)

func main() {
	log := logger.CreateLogger(
		&logger.LoggerOptions{
			LogLevel: 5,
			Std:      true,
			Verbose:  false,
		},
	)
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		params := r.URL.Query()
		log.Informln("req params ", params)

		if params.Get("name") == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Name is required"))
			return
		}
		log.Informln("Hello " + params.Get("name"))

		w.Write([]byte("Hello " + params.Get("name")))
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

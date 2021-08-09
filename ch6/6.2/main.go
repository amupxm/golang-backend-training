package main

import (
	"encoding/json"
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
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "Method not allowed",
				},
			)
			w.Write(errMsg)
			return
		}
		params := r.URL.Query()
		log.Informln("req params ", params)

		if params.Get("name") == "" {
			w.WriteHeader(http.StatusBadRequest)
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "Name is required",
				},
			)
			w.Write([]byte(errMsg))
			return
		}
		log.Informln("Hello " + params.Get("name"))
		msg, _ := json.Marshal(
			map[string]string{
				"msg": "Hello " + params.Get("name"),
			},
		)
		w.Write(msg)
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

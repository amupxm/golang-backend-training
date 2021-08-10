package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	logger "github.com/amupxm/xmus-logger"
)

func main() {
	loggerOptions := logger.LoggerOptions{
		LogLevel: 6,
	}
	log := logger.CreateLogger(&loggerOptions)
	log.Informln("Starting server")
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "Method not allowed",
				},
			)
			w.Header().Set("Content-Type", "application/json")
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
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(errMsg))
			return
		}
		response := struct {
			Msg string `json:"message"`
		}{
			Msg: fmt.Sprintf("Hello %s", params.Get("name")),
		}
		resAsBytes, _ := json.Marshal(response)
		log.InformF("response is %s\n", response.Msg)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resAsBytes)
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

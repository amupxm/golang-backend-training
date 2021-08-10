package main

import (
	"encoding/json"
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
		log.InformF("🆕 new request received")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Warn("request  %s rejected", r.Method)

			w.Write([]byte("Method not allowed"))
			return
		}
		d, _ := json.Marshal(
			map[string]string{
				"msg": "Hello, world!",
			},
		)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

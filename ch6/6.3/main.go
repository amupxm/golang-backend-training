package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	logger "github.com/amupxm/xmus-logger/srv"
)

func main() {

	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log := logger.CreateLogger(
			&logger.LoggerOptions{
				LogLevel: 5,
				Std:      true,
				Verbose:  false,
			},
		)
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

		if params.Get("x") == "" {
			w.WriteHeader(http.StatusBadRequest)
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "x is required",
				},
			)
			log.Informln("err : req not contains x ")

			w.Write([]byte(errMsg))
			return
		}

		if params.Get("y") == "" {
			w.WriteHeader(http.StatusBadRequest)
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "y is required",
				},
			)
			log.Informln("err : req not contains y ")

			w.Write([]byte(errMsg))
			return
		}
		theX, err := strconv.ParseInt(params.Get("x"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "x is not int",
				},
			)
			w.Write([]byte(errMsg))
			return
		}
		theY, err := strconv.ParseInt(params.Get("y"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMsg, _ := json.Marshal(
				map[string]string{
					"error": "y is not int",
				},
			)
			w.Write([]byte(errMsg))
			return
		}
		msg, _ := json.Marshal(
			map[string]string{
				"result": fmt.Sprint(theX + theY),
			},
		)
		w.Write(msg)

		log.End()
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	logger "github.com/amupxm/xmus-logger"
)

var log logger.Logger

func main() {

	log = logger.CreateLogger(
		&logger.LoggerOptions{
			LogLevel: 5,
			Std:      true,
		},
	)
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			ResponseError("Method not allowed", http.StatusMethodNotAllowed, w)

			return
		}
		params := r.URL.Query()
		y, err := ExtractParams("y", &params)
		if err != nil {
			log.Errorln(err)
			ResponseError("missing numberic y in params", http.StatusBadRequest, w)
			return
		}
		x, err := ExtractParams("x", &params)
		if err != nil {
			log.Errorln(err)
			ResponseError("missing numberic x in params ", http.StatusBadRequest, w)
			return
		}

		msg, _ := json.Marshal(
			map[string]string{
				"result": fmt.Sprint(x + y),
			},
		)
		w.Header().Set("Content-Type", "application/json")

		w.Write(msg)

		log.End()
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

func ResponseError(msg string, code int, w http.ResponseWriter) {
	errRes := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	b, _ := json.Marshal(errRes)
	log.Prefix("RESPONSE", "ERROR").Warnln(errRes.Message)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return
}

func ExtractParams(paramKey string, params *url.Values) (int, error) {
	value, err := strconv.ParseInt(params.Get(paramKey), 10, 64)
	if err != nil {
		return 0, err
	}
	res := int(value)
	return res, nil
}

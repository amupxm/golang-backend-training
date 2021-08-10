package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	logger "github.com/amupxm/xmus-logger"
	"github.com/reiver/go-simplehttp"
	_ "github.com/reiver/go-simplehttp/driver/json"
)

var log logger.Logger

func main() {

	httpHandler := http.NewServeMux()
	shttp, err := simplehttp.Load("json")
	if err != nil {
		panic(err)
	}

	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log = logger.CreateLogger(
			&logger.LoggerOptions{
				LogLevel: 5,
				Std:      true,
			},
		)
		if r.Method != "GET" {
			shttp.MethodNotAllowed(w)
			return
		}
		params := r.URL.Query()
		y, err := ExtractParams("y", &params)
		if err != nil {
			log.Errorln(err)
			shttp.BadRequest(w, struct {
				Message string `key:"message"`
			}{Message: "missing numberic y in params"})
			return
		}
		x, err := ExtractParams("x", &params)
		if err != nil {
			log.Errorln(err)
			shttp.BadRequest(w, struct {
				Message string `key:"message"`
			}{Message: "missing numberic x in params"})
			return
		}

		shttp.OK(w, struct {
			Result string `key:"result"`
		}{Result: fmt.Sprint(x + y)})
		log.End()
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

func ExtractParams(paramKey string, params *url.Values) (int, error) {
	value, err := strconv.ParseInt(params.Get(paramKey), 10, 64)
	if err != nil {
		return 0, err
	}
	res := int(value)
	return res, nil
}

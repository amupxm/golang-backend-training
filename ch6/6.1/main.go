package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		d, _ := json.Marshal(
			map[string]string{
				"msg": "Hello, world!",
			},
		)
		w.Write(d)
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

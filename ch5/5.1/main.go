package main

import "net/http"

func main() {

	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		w.Write([]byte("Hello, world!"))
		return
	})
	http.ListenAndServe(":8080", httpHandler)
}

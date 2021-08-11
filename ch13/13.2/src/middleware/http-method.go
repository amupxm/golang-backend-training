package middleware

import (
	"log"
	"net/http"
	"strings"
)

type (
	middleWare struct{}
	MiddleWare interface {
		SetXMethodHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, method string)
		MethodUpgrader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		LogUserAgent(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		isValidMethod(method string) bool
	}
)

func NewMiddleWare() MiddleWare {
	return &middleWare{}
}

func (m *middleWare) isValidMethod(method string) bool {
	validMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
	for _, v := range validMethods {
		if v == method {
			return true
		}
	}
	return false
}

func (m *middleWare) MethodUpgrader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get("X-HTTP-Method-Override") != "" {
		method := r.Header.Get("X-HTTP-Method-Override")
		if !m.isValidMethod(method) {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		r.Method = strings.TrimSpace(method)
	}
	next(w, r)

}

func (m *middleWare) SetXMethodHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, method string) {
	if r.Header.Get("X-Method") == "" {
		r.Header.Set("X-Method", method)
	}
	next(w, r)
}

func (m *middleWare) LogUserAgent(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get("User-Agent") != "" {
		log.Printf("New Request from : %s\n", r.Header.Get("User-Agent"))
	}
	next(w, r)
}

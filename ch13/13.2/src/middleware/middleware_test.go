package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amupxm/golang-backend-training/ch13/13.2/src/middleware"
)

func TestUpgrader(t *testing.T) {
	testTable := []struct {
		Method       string
		CustomMethod string
		OK           bool
	}{
		{"GET", "GET", true},
		{"GET", "POST", true},
		{"POST", "ARE", false},
		{"POST", "", true},
	}
	for testIndex, test := range testTable {
		req, err := http.NewRequest(test.Method, "/", nil)
		req.Header.Set("X-HTTP-Method-Override", test.CustomMethod)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		innerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(r.Method))
		})
		mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middleware.NewMiddleWare().MethodUpgrader(w, r, innerHandler)
		})

		mainHandler.ServeHTTP(rr, req)
		res := rr.Body.String()
		if (res != test.CustomMethod) == test.OK {
			if test.CustomMethod == "" && (test.Method == res) {
				continue
			}
			t.Errorf("Test %d: Expected %s, got %s", testIndex, test.CustomMethod, rr.Body.String())
			continue
		} else {
			t.Logf("Test %d: OK", testIndex)

		}
	}
}

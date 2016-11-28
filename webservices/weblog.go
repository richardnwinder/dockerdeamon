// Package webservices supplies common services fpr web http requests
package webservices

import (
	"log"
	"net/http"
	"time"
)

// LogHandler initialises authentication webservice
func Log( next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
		log.Printf("[%s] %q\n", r.Method, r.URL.String())
    next(w,r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }
  return http.HandlerFunc(fn)
}

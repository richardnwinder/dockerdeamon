// Package webservices supplies common services fpr web http requests
package webservices

import (
	"fmt"
	"log"
	"net/http"
)

// Recover catches panics and logs the error and maintains the application  process running
func Recovery(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
    defer func() {
			err := recover()
	    if err != nil {
	      log.Printf("panic: %+v", err)
	      http.Error(w, http.StatusText(500), 500)
	    }
		}()
		fmt.Print("Recovery handler\n")
    next(w, r)
  }
  return http.HandlerFunc(fn)
}

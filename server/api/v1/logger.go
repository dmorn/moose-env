package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Logger used to log request information and will be used to check authentication state of the requests
func Logger(inner http.Handler, name string, auth bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//here we should authenticate some calls, to be defined
		if auth {
			fmt.Printf("Authentication required!\n")
		}

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

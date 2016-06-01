package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/freehaha/token-auth"
	"github.com/freehaha/token-auth/memory"
)

var tokenauth *tauth.TokenAuth
var memStore = memstore.New("salty")

func init() {

	bearerGetter := tauth.BearerGetter{}
	tokenauth = tauth.NewTokenAuth(nil, nil, memStore, &bearerGetter)

}

//MMiddleware used to log request information and will be used to check authentication state of the requests
func MMiddleware(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		if strings.Compare("Login", name) == 0 {
			inner.ServeHTTP(w, r)
		} else {
			token := tauth.Get(r)
			if token != nil {
				//here we should authenticate some calls, to be defined
				inner.ServeHTTP(w, r)
			}
		}

		log.Printf(
			"%s\t%s\t\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/freehaha/token-auth"
	"github.com/freehaha/token-auth/memory"
)

var tokenauth *tauth.TokenAuth
var memStore = memstore.New("salty")

//var memStore tauth.MemoryTokenStore

func init() {
	//memStore = memstore.New("salty")
	var bearerGetter = tauth.BearerGetter{}
	tokenauth = tauth.NewTokenAuth(nil, nil, memStore, &bearerGetter)
}

//MMiddleware used to log request information and will be used to check authentication state of the requests
func MAuthMiddleware(inner http.Handler, name string) http.Handler {
	return tokenauth.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		token := tauth.Get(r)
		fmt.Println(token)
		inner.ServeHTTP(w, r)

		mLog(start, name, r)
	})
}

func MMiddleware(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		inner.ServeHTTP(w, r)
		mLog(start, name, r)
	})
}

func mLog(start time.Time, name string, r *http.Request) {
	log.Printf(
		"%s\t%s\t\t%s\t%s",
		r.Method,
		r.RequestURI,
		name,
		time.Since(start),
	)
}

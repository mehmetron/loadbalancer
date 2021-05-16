package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	go healthCheck(w)

	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/", loadBalancer())

	http.ListenAndServe(":8080", mux)

}

func loadBalancer() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Info: ", r.Host, r.Method, r.Proto)

		parsedUrl := weightedRandom()
		fmt.Println("57 ", parsedUrl)

		proxy := httputil.NewSingleHostReverseProxy(parsedUrl)
		proxy.ServeHTTP(w, r)
	}
}

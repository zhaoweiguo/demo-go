// from: http://www.gorillatoolkit.org/pkg/mux
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


/*
简单说明
*/
func main() {
	r := mux.NewRouter()
	// Only matches if domain is "www.domain.com".
	r.Host("www.domain.com")
	// Matches a dynamic subdomain.
	r.Host("{subdomain:[a-z]+}.domain.com")

	// match path prefixes
	r.PathPrefix("/products/")

	//http method、schema、header、query value
	r.Methods("GET", "POST")
	r.Schemes("https")
	r.Headers("X-Requested-With", "XMLHttpRequest")
	r.Queries("key", "value")

	//use a custom matcher function:
	r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
        return r.ProtoMajor == 0
    })

	// combine several matchers in a single route:
	r.HandleFunc("/products", ProductsHandler).
		Host("www.domain.com").
		Methods("GET").
		Schemes("http")

	// subrouter
	s := r.Host("www.domain.com").Subrouter()
	s.HandleFunc("/products/", ProductsHandler)
	s.HandleFunc("/products/{key}", ProductHandler)
	s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	// When a subrouter has a path prefix, the inner routes use it as base for their paths:
	s := r.PathPrefix("/products").Subrouter()
	// "/products/"
	s.HandleFunc("/", ProductsHandler)
	// "/products/{key}/"
	s.HandleFunc("/{key}/", ProductHandler)
	// "/products/{key}/details"
	s.HandleFunc("/{key}/details", ProductDetailsHandler)


	//  registered URLs.
	// 用的時候再看

	http.Handle("/", r)
	addr := ":7800"
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
	}
	server.ListenAndServe()
}




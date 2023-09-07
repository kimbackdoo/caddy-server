package main

import (
	"fmt"
	"net/http"
)

func main() {
	allowedDomains := []string{"www.caddy-example.o-r.kr", "www.caddy-fail.kro.kr"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

		requestDomain := r.URL.Query().Get("domain")
		if requestDomain == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Empty domain is not allowed")
			return
		}

		for _, domain := range allowedDomains {
			if domain == requestDomain {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "OK")
				fmt.Println("Domain is allowed: ", requestDomain)
				return
			}
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Domain is not allowed: ", requestDomain)
	})

	fmt.Println("🚀 Starting TLS Check server on port 5555")
	http.ListenAndServe(":5555", nil)
}
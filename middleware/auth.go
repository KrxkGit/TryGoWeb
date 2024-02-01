package middleware

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}
	username, password, ok := r.BasicAuth()
	fmt.Println(ok)
	if ok && username == "Krxk" && password == "Krxk" {
		am.Next.ServeHTTP(w, r)
	} else {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		//w.WriteHeader(http.StatusUnauthorized)
	}
}

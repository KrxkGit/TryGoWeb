package controller

import "net/http"

func registerRedirectRoute() {
	http.HandleFunc("/redirect", handleRedirect)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	/* 重定向 */
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(http.StatusFound)
}

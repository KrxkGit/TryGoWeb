package controller

import (
	"log"
	"net/http"
)

func registerRootRoute() {
	//http.Handle("/", http.FileServer(http.Dir("resources")))
	http.HandleFunc("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("ROOT"))
	if pusher, ok := w.(http.Pusher); ok {
		/*服务器推送在新浏览器可能被废除*/
		log.Println("pusher run")
		err := pusher.Push("/resources/styles/index.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
	if r.URL.Path == "" {
		r.URL.Path = "index.html"
	}
	http.ServeFile(w, r, "resources"+r.URL.Path)
}

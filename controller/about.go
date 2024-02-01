package controller

import (
	json2 "encoding/json"
	"net/http"
)

func registerAboutRoute() {
	mh := myHandler{}
	http.Handle("/about", &mh)
}

type myHandler struct{}
type Post struct {
	User    string
	Threads []string
}

/*实现 ServeHTTP 接口*/
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query["id"][0]
	w.Write([]byte("about:" + id + "\n"))

	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Krxk",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json2.Marshal(post)
	w.Write(json)
}

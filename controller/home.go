package controller

import (
	"fmt"
	"net/http"
)

func registerHomeRoute() {
	http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.URL.Path)
	fmt.Fprintln(w, r.URL.RawQuery) /*查询字符串*/
	fmt.Fprintln(w, r.URL.Query())  /*解析查询字符串为map*/
	fmt.Fprintln(w, r.URL.Fragment)

	w.WriteHeader(http.StatusOK) /*返回状态码*/

	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body)) /*若无指定，前512字节用于自动推测内容类型*/

}

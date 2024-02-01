package controller

import (
	"fmt"
	"net/http"
)

func registerFormRoute() {
	http.HandleFunc("/form", handleForm)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)     /*包括 url 包含的数据， 表单数据靠前， url 靠后*/
	fmt.Fprintln(w, r.PostForm) /*只包含表单提供的数据*/

	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)

	fmt.Fprintln(w, r.FormValue("first_name"))     /*自动解析，且返回第一个值*/
	fmt.Fprintln(w, r.PostFormValue("first_name")) /*自动解析，且返回第一个值(只解析表单数据)*/
}

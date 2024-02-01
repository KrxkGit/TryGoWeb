package controller

import "net/http"

func registerRootRoute() {
	http.Handle("/", http.FileServer(http.Dir("resources")))
}

/*func handleRoot(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("ROOT"))
	//http.ServeFile(w, r, http.Dir("resources")) // 从文件系统返回文件作为响应
}*/

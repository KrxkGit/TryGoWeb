package main

import (
	"github.com/KrxkGit/TryGoWeb/controller"
	_ "github.com/KrxkGit/TryGoWeb/mapper"
	"github.com/KrxkGit/TryGoWeb/middleware"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: new(middleware.AuthMiddleware), /* nil 则使用 DefaultServeMux */
	}
	controller.RegisterRoutes() /*注册路由*/
	server.ListenAndServe()
}

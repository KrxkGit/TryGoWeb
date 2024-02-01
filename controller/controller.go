package controller

import "github.com/KrxkGit/TryGoWeb/controller/dbOper"

// RegisterRoutes ...
func RegisterRoutes() {
	registerRootRoute()
	registerHomeRoute()
	registerAboutRoute()
	registerFormRoute()
	registerHandleUploadRoute()
	registerRedirectRoute()
	dbOper.RegisterDbOperRoute()
}

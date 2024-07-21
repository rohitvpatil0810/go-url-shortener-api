package routes

import "net/http"

func RegisterRoutes(prefix string, router *http.ServeMux) {
	RegisterUserRoutes(prefix+"/users", router)
}

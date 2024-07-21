package routes

import "net/http"

func RegisterUserRoutes(prefix string, router *http.ServeMux) {
	router.HandleFunc(path(http.MethodPost, prefix, "/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p>POST /users/</p>"))
	})
	router.HandleFunc(path(http.MethodPost, prefix, "/login/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p>POST /users/login/</p>"))
	})
	router.HandleFunc(path(http.MethodPost, prefix, "/logout/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p>POST /users/logout/</p>"))
	})
}

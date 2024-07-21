package routes

import (
	"fmt"
	"net/http"
)

func path(method string, prefix string, path string) string {
	preparedPath := method + " " + prefix + path
	fmt.Println(preparedPath)
	return preparedPath
}

func RegisterV1Routes() *http.ServeMux {
	router := http.NewServeMux()

	RegisterRoutes("/v1", router)

	return router
}

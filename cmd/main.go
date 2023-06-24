package main

import (
	"net/http"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes"
)

func main() {
	err := http.ListenAndServe(config.Host+":"+config.Port, routes.UsersInit())
	if err != nil {
		panic(err)
	}
}

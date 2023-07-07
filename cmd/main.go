package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes/packages"
	"strux_api/internal/rest_api/routes/users"
)

func main() {
	err := http.ListenAndServe(config.Host+":"+config.Port, InitRoutes())
	if err != nil {
		panic(err)
	}
}

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/create-user/", users.CreateUserService)
	r.Get("/user-exist/", users.UserExistService)
	r.Delete("/user-delete/", users.UserDeleteService)
	r.Put("/user-password-update/", users.UserPasswordUpdateService)
	r.Get("/user-log-in/", users.UserLogInService)

	r.Post("/upload-pkg/", packages.UploadPkgService)
	r.Get("/exists-package/", packages.PackageExistsService)
	r.Get("/download-package/", packages.DownloadPackage)
	r.Get("/package-versions/", packages.ShowVersions)
	return r
}

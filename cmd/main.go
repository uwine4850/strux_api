package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/internal/rest_api/routes/packages"
	"github.com/uwine4850/strux_api/internal/rest_api/routes/users"
	"net/http"
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

package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"trueconf-userapi/services/user_service"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()
	userService := user_service.NewUserService()

	r.Get("/", userService.GetUsers)
	r.Post("/", userService.CreateUser)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", userService.GetUser)
		r.Patch("/", userService.UpdateUser)
		r.Delete("/", userService.DeleteUser)
	})

	return r
}

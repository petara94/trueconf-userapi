package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
	"trueconf-userapi/pkg/app"
	v1 "trueconf-userapi/routers/api/v1"
)

func ApiRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		app.SendResponse(w, r, app.CODE_SUCCESS, app.MESSAGE_WELCOME, nil)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Mount("/users", v1.UserRouter())
		})
	})

	return r
}

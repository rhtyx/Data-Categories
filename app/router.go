package app

import (
	"Data-Category/controller"
	"Data-Category/exception"

	"github.com/go-chi/chi/v5"
)

func NewRouter(cc controller.CategoryController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(exception.ErrorHandler)

	r.Route("/api/categories", func(r chi.Router) {

		r.Get("/", cc.FindAll)
		r.Post("/", cc.Create)
		r.Delete("/", cc.DeleteAll)

		r.Route("/{categoryId}", func(r chi.Router) {
			r.Get("/", cc.FindById)
			r.Put("/", cc.UpdateById)
			r.Delete("/", cc.DeleteById)
		})
	})

	return r
}

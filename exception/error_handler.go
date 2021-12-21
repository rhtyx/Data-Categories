package exception

import (
	"Data-Category/helper"
	"Data-Category/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				var webResponse web.WebResponse
				w.Header().Set("Content-Type", "application/json")

				if exception, ok := rvr.(NotFoundError); ok {
					w.WriteHeader(http.StatusNotFound)
					webResponse = web.WebResponse{
						Code:   http.StatusNotFound,
						Status: "Not Found",
						Data:   exception.Error,
					}
				} else if exception, ok := rvr.(validator.ValidationErrors); ok {
					w.WriteHeader(http.StatusBadRequest)
					webResponse = web.WebResponse{
						Code:   http.StatusBadRequest,
						Status: "Bad Request",
						Data:   exception.Error(),
					}
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					webResponse = web.WebResponse{
						Code:   http.StatusInternalServerError,
						Status: "Internal Server Error",
						Data:   rvr,
					}
				}

				helper.WriteToResponseBody(w, webResponse)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

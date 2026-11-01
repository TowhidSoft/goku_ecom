package review

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /reviews",
		manager.With(
			http.HandlerFunc(h.GetReviews),
		),
	)
}

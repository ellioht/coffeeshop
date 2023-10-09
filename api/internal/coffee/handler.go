package coffee

import (
	"fmt"
	"github.com/go-chi/chi"
)

type Handler struct {
	s *Service
}

func NewHandler(r *chi.Mux, s *Service) *Handler {
	h := &Handler{
		s: s,
	}
	h.SetupRoutes(r)
	return h
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	fmt.Println("Setting up coffee routes")

	r.Group(func(r chi.Router) {

	})
}

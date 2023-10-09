package health

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type Handler struct{}

func NewHandler(r *chi.Mux) *Handler {
	h := &Handler{}
	h.SetupRoutes(r)
	return h
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	fmt.Println("Setting up health routes")

	r.Group(func(r chi.Router) {
		r.Get("/health", h.Health)
	})
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	write, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
	fmt.Printf("Wrote %v bytes\n", write)
}

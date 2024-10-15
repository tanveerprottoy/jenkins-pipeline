package resource

import (
	"net/http"

	"github.com/tanveerprottoy/jenkins-pipeline/service/pkg/response"
)

// Hanlder is responsible for extracting data
// from request body and building and seding response
type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	m, err := h.service.GetData(ctx)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	response.Respond(w, http.StatusOK, response.BuildData(m))
}

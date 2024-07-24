package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/domain/usecases"
	"github.com/go-chi/chi"
)

type DeletePostHandler struct {
	deletePostUseCase usecases.DeletePostUseCase
}

func NewDeletePostHandler(deletePostUseCase usecases.DeletePostUseCase) *DeletePostHandler {
	return &DeletePostHandler{deletePostUseCase: deletePostUseCase}
}

func (h *DeletePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	input := &usecases.DeletePostUseCaseInput{
		ID: chi.URLParam(r, "id"),
	}

	errUseCase := h.deletePostUseCase.Execute(input)
	if errUseCase != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errUseCase)
		return
	}

	w.WriteHeader(http.StatusOK)
}
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/usecases"
)

type CreatePostHandler struct {
	createPostUseCase usecases.CreatePostUseCase
	result            *errors.ErrorModel
}

func NewCreatePostHandler(createPostUseCase usecases.CreatePostUseCase) *CreatePostHandler {
	return &CreatePostHandler{createPostUseCase: createPostUseCase}
}

func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input usecases.CreatePostUseCaseInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errResponse := errors.NewErrorModel(nil, "Invalid request body")
		h.result = errResponse
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	output, errUseCase := h.createPostUseCase.Execute(&input)
	if errUseCase != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errUseCase)
		h.result = errUseCase
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *CreatePostHandler) GetError() *errors.ErrorModel {
	return h.result
}

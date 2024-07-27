package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
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

// Create Post godoc
// @Summary 		Create a new post
// @Description Create a new post
// @Tags 				posts
// @Accept 			json
// @Produce 		json
// @Param 			request	body dto.CreatePostInput true "Post Request"
// @Success 		201 {object} dto.CreatePostOutput
// @Failure 		400 {object} errors.ErrorModel
// @Router 			/post [post]
func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input dto.CreatePostInput

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

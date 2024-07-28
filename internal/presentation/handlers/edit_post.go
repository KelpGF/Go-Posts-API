package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/usecases"
	"github.com/go-chi/chi"
)

type EditPostHandler struct {
	editPostUseCase usecases.EditPostUseCase
	result          *errors.ErrorModel
}

func NewEditPostHandler(editPostUseCase usecases.EditPostUseCase) *EditPostHandler {
	return &EditPostHandler{editPostUseCase: editPostUseCase}
}

// Edit Post godoc
// @Summary 		Edit a new post
// @Description Edit a new post
// @Tags 				posts
// @Accept 			json
// @Produce 		json
// @Param 			id	path string	true	"Post ID" format(uuid)
// @Param 			request	body dto.EditPostInput true "Post Request"
// @Success 		200
// @Failure 		400 {object} errors.ErrorModel
// @Router 			/post/{id} [put]
func (h *EditPostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")

	log.Print(postId)

	var input dto.EditPostInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errResponse := errors.NewErrorModel(nil, "Invalid request body")
		h.result = errResponse
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	input.ID = postId

	errUseCase := h.editPostUseCase.Execute(&input)
	if errUseCase != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errUseCase)
		h.result = errUseCase
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *EditPostHandler) GetError() *errors.ErrorModel {
	return h.result
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/usecases"
	"github.com/go-chi/chi"
)

type DeletePostHandler struct {
	deletePostUseCase usecases.DeletePostUseCase
}

func NewDeletePostHandler(deletePostUseCase usecases.DeletePostUseCase) *DeletePostHandler {
	return &DeletePostHandler{deletePostUseCase: deletePostUseCase}
}

// Delete Post godoc
// @Summary 		Delete post
// @Description Delete post by id
// @Tags 				posts
// @Accept 			json
// @Produce 		json
// @Param 			id	path string	true	"Post ID" format(uuid)
// @Success 		200
// @Failure 		400 {object} errors.ErrorModel
// @Router 			/post/{id} [delete]
func (h *DeletePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	input := &dto.DeletePostInput{
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

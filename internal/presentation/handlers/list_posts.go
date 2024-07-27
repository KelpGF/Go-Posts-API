package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/usecases"
)

type ListPostsHandler struct {
	listPostsUseCase usecases.ListPostsUseCase
}

func NewListPostsHandler(listPostsUseCase usecases.ListPostsUseCase) *ListPostsHandler {
	return &ListPostsHandler{listPostsUseCase: listPostsUseCase}
}

// ListByPagination godoc
// @Summary 		List Posts by pagination
// @Description List Posts by pagination
// @Tags 				posts
// @Accept 			json
// @Produce 		json
// @Param 			page	query	int	false	"Page Number"
// @Param 			limit	query	int	false	"Limit of posts"
// @Param 			enumstring sort	query	string	false	"Sort by" Enums(asc, desc)
// @Param 			author_name	query	string	false	"Author Name"
// @Success 		200 {array} dto.ListPostsOutput
// @Router 			/post [get]
func (h *ListPostsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	authorName := r.URL.Query().Get("author_name")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	input := &dto.ListPostsInput{
		AuthorName:    authorName,
		PublishedSort: sort,
		Paginate: dto.Paginate{
			Page:  pageInt,
			Limit: limitInt,
		},
	}

	output := h.listPostsUseCase.Execute(input)

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

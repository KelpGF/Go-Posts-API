package repositories

import "github.com/KelpGF/Go-Posts-API/internal/domain/dto"

type ListPostsRepository interface {
	List(input *dto.ListPostsInput) *dto.ListPostsOutput
}

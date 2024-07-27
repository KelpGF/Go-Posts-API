package repositories

import "github.com/KelpGF/Go-Posts-API/internal/domain/dto"

type DeletePostRepository interface {
	Delete(input *dto.DeletePostInput) error
}

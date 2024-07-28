package factories

import (
	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/presentation/handlers"
	"gorm.io/gorm"
)

func EditPostHandler(db *gorm.DB) *handlers.EditPostHandler {
	postFactory := entities.NewPostFactory()
	findPostByIdRepository := repositories.NewFindPostByIdRepository(db, postFactory)
	editPostRepository := repositories.NewEditPostRepository(db)
	usecase := usecases.NewEditPostUseCase(findPostByIdRepository, editPostRepository)

	return handlers.NewEditPostHandler(usecase)
}

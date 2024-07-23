package factories

import (
	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/presentation/handlers"
	"gorm.io/gorm"
)

func CreatePostHandler(db *gorm.DB) *handlers.CreatePostHandler {
	postRepository := repositories.NewCreatePostRepository(db)
	postFactory := entities.NewPostFactory()
	usecase := usecases.NewCreatePostUseCase(postRepository, postFactory)

	return handlers.NewCreatePostHandler(usecase)
}

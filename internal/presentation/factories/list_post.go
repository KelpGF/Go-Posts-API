package factories

import (
	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/presentation/handlers"
	"gorm.io/gorm"
)

func ListPostsHandler(db *gorm.DB) *handlers.ListPostsHandler {
	postRepository := repositories.NewListPostsRepository(db)
	usecase := usecases.NewListPostsUseCase(postRepository)

	return handlers.NewListPostsHandler(usecase)
}

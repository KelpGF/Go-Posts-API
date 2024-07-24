package factories

import (
	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/presentation/handlers"
	"gorm.io/gorm"
)

func DeletePostHandler(db *gorm.DB) *handlers.DeletePostHandler {
	postRepository := repositories.NewDeletePostRepository(db)
	usecase := usecases.NewDeletePostUseCase(postRepository)

	return handlers.NewDeletePostHandler(usecase)
}

package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func createRouter(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	mapperPostsRoutes(router, db)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/docs/doc.json")))

	return router
}

func mapperPostsRoutes(router *chi.Mux, db *gorm.DB) {
	postRepository := repositories.NewCreatePostRepository(db)
	postFactory := entities.NewPostFactory()
	usecase := usecases.NewCreatePostUseCase(postRepository, postFactory)

	router.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		var input usecases.CreatePostUseCaseInput

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errResponse := errors.NewErrorModel(nil, "Invalid request body")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		output, errUseCase := usecase.Execute(&input)
		if errUseCase != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errUseCase)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(output)
	})

}

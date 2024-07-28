package presentation

import (
	"net/http"

	"github.com/KelpGF/Go-Posts-API/internal/presentation/factories"
	"github.com/KelpGF/Go-Posts-API/internal/presentation/handlers"
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
	router.Route("/post", func(r chi.Router) {
		r.Post("/", handlers.NewLogDecoratorHandler(factories.CreatePostHandler(db)).Handle)
		r.Get("/", factories.ListPostsHandler(db).Handle)
		r.Put("/{id}", handlers.NewLogDecoratorHandler(factories.EditPostHandler(db)).Handle)
		r.Delete("/{id}", factories.DeletePostHandler(db).Handle)
	})
}

package handlers

import (
	"log"
	"net/http"
)

type LogDecoratorHandler struct {
	handler HttpHandler
}

func NewLogDecoratorHandler(handler HttpHandler) *LogDecoratorHandler {
	return &LogDecoratorHandler{handler: handler}
}

func (h *LogDecoratorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	h.handler.Handle(w, r)

	result := h.handler.GetError()
	if result != nil {
		log.Printf("Erro Message: %s", result.Message)
		log.Printf("Erro list: %v", result.Errors)

		return
	}
}

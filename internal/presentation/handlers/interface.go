package handlers

import (
	"net/http"

	domainErrors "github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type HttpHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
	GetError() *domainErrors.ErrorModel
}

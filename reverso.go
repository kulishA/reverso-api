package reverso_api

import (
	"github.com/kulishA/reverso-api/domain"
	"github.com/kulishA/reverso-api/internal/api"
)

type TranslateInterface interface {
	Translate(request *domain.TranslationRequest) (*domain.TranslationResponse, error)
}

type ReversoApi struct {
	Translate TranslateInterface
}

func NewReversoApi() *ReversoApi {
	return &ReversoApi{Translate: api.NewTranslation()}
}

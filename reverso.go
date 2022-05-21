package reverso_api

import (
	"github.com/kulishA/reverso-api/internal/api"
)

type TranslateInterface interface {
	Translate(request *TranslationRequest) (*TranslationResponse, error)
}

type ReversoApi struct {
	Translation TranslateInterface
}

func NewReversoApi() *ReversoApi {

	return &ReversoApi{Translation: api.NewTranslation()}
}

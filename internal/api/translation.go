package api

import (
	"encoding/json"
	"fmt"
	"github.com/kulishA/reverso-api/domain"
)

type Translation struct {
}

func NewTranslation() *Translation {
	return &Translation{}
}

func (t *Translation) Translate(data *domain.TranslationRequest) (*domain.TranslationResponse, error) {
	res := domain.TranslationResponse{}

	d, err := postRequest(translationPath, data)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(d))
	err = json.Unmarshal(d, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

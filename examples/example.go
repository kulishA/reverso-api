package main

import (
	"fmt"
	reverso_api "github.com/kulishA/reverso-api"
	"github.com/kulishA/reverso-api/domain"
)

func main() {
	r := reverso_api.NewReversoApi()

	req := domain.TranslationRequest{
		Format: "text",
		From:   reverso_api.ENG,
		To:     reverso_api.RU,
		Input:  "Hello world!",
		Options: domain.Options{
			SentenceSplitter:  true,
			Origin:            "translation.web",
			ContextResults:    false,
			LanguageDetection: false,
		},
	}

	t, err := r.Translation.Translate(&req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(t)
}

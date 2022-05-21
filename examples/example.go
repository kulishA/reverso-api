package main

import (
	"fmt"
	reverso_api "github.com/kulishA/reverso-api"
)

func main() {
	r := reverso_api.NewReversoApi()

	req := reverso_api.TranslationRequest{
		Format: "text",
		From:   reverso_api.ENG,
		To:     reverso_api.RU,
		Input:  "Hello world!",
		Options: reverso_api.Options{
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

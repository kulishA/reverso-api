# Reverso Go API
![logotype](./assets/reversoapi-logo.png)

This API is not official! It allows you to manipulate with your text in different ways. Currently supported: context, translation


## Navigation
- [Installing](#installing)
- [Usage](#usage)
- [Examples](#examples)
  - [Translation] (#translation)

## Installing
```bash
go get -u https://github.com/kulishA/reverso-api
```

## Usage

```go
r := reverso_api.NewReversoApi()
req := reverso_api.TranslationRequest{...}
t, err := r.Translation.Translate(&req)
```

## Examples

### Transalation
```go
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
```
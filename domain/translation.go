package domain

type TranslationRequest struct {
	Format  string  `json:"format"`
	From    string  `json:"from"`
	To      string  `json:"to"`
	Input   string  `json:"input"`
	Options Options `json:"options"`
}

type Options struct {
	SentenceSplitter  bool   `json:"sentenceSplitter"`
	Origin            string `json:"origin"`
	ContextResults    bool   `json:"contextResults"`
	LanguageDetection bool   `json:"languageDetection"`
}

type TranslationResponse struct {
	Id                string            `json:"id"`
	From              string            `json:"from"`
	To                string            `json:"to"`
	Input             []string          `json:"input"`
	CorrectedText     string            `json:"correctedText"`
	Translation       []string          `json:"translation"`
	Engines           []string          `json:"engines"`
	LanguageDetection LanguageDetection `json:"language_detection"`
	ContextResults    ContextResult     `json:"contextResults"`
	Truncated         bool              `json:"truncated"`
	TimeTaken         int               `json:"timeTaken"`
}

type ContextResult struct {
	RudeWords             bool     `json:"rudeWords"`
	Colloquialisms        bool     `json:"colloquialisms"`
	RiskyWords            bool     `json:"riskyWords"`
	Results               []Result `json:"results"`
	TotalContextCallsMade int      `json:"totalContextCallsMade"`
	TimeTakenContext      int      `json:"timeTakenContext"`
}

type Result struct {
	Translation     string      `json:"translation"`
	SourceExamples  []string    `json:"sourceExamples"`
	TargetExamples  []string    `json:"targetExamples"`
	Rude            bool        `json:"rude"`
	Colloquial      bool        `json:"colloquial"`
	PartOfSpeech    interface{} `json:"partOfSpeech"`
	Frequency       int         `json:"frequency"`
	Vowels          interface{} `json:"vowels"`
	Transliteration interface{} `json:"transliteration"`
}

type LanguageDetection struct {
	DetectedLanguage                string `json:"detectedLanguage"`
	IsDirectionChanged              bool   `json:"isDirectionChanged"`
	OriginalDirection               string `json:"originalDirection"`
	OriginalDirectionContextMatches int    `json:"originalDirectionContextMatches"`
	ChangedDirectionContextMatches  int    `json:"changedDirectionContextMatches"`
	TimeTaken                       int    `json:"timeTaken"`
}

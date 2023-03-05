package types

// Messages is a map of strings
type Messages map[string]string

// Params is a map of strings
type Params map[string]string

// Language is a struct that contains the name and driver of a language
type Language struct {
	Name           string
	Driver         string
	TranslationDir string
}

package go_transify

import (
	"fmt"
	"github.com/hoitek-go/go-transify/language"
	"github.com/hoitek-go/go-transify/transify"
	"testing"
)

func TestTransify(t *testing.T) {
	// Create a new bundle
	bundle := transify.NewBundle(language.English)

	// Set the translation directory
	bundle.SetTranslationDir("./translations")

	// Load the messages
	err := bundle.LoadMessages()
	if err != nil {
		t.Error(err)
	}

	// Get the message
	message := bundle.T("asghar.reza")

	// Print the message
	fmt.Println(message)
}

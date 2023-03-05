package transify

import (
	"encoding/json"
	"fmt"
	"github.com/hoitek-go/go-transify/drivers"
	"github.com/hoitek-go/go-transify/types"
	"log"
	"os"
	"strings"
)

// Bundle is a struct that contains the language, translation directory and messages and ...
type bundle struct {
	Language types.Language
	Messages types.Messages
}

// CurrentBundle is the current bundle
var CurrentBundle *bundle

// NewBundle creates a new bundle
func NewBundle(language string) *bundle {
	CurrentBundle = &bundle{
		Language: types.Language{
			Name:   language,
			Driver: drivers.FileDriver,
		},
		Messages: types.Messages{},
	}
	return CurrentBundle
}

// SetDriver sets the driver
func (b *bundle) SetDriver(driver string) {
	b.Language.Driver = driver
}

// SetLanguage sets the language
func (b *bundle) SetLanguage(language string) {
	b.Language = types.Language{
		Name:   language,
		Driver: drivers.FileDriver,
	}
}

// GetLanguage gets the language
func (b *bundle) GetLanguage() types.Language {
	return b.Language
}

// SetTranslationDir sets the translation directory
func (b *bundle) SetTranslationDir(dir string) {
	fmt.Println("Setting translation dir to", dir)
	b.Language.TranslationDir = dir
}

// GetTranslationDir gets the translation directory
func (b *bundle) GetTranslationDir() string {
	return b.Language.TranslationDir
}

// LoadMessages loads the messages from the translation directory
func (b *bundle) LoadMessages() error {
	// Get driver
	driver, err := drivers.GetDriver(b.Language)
	if err != nil {
		return fmt.Errorf("error getting driver: %w", err)
	}

	// Load messages
	messages, err := driver.LoadMessages()
	if err != nil {
		return fmt.Errorf("error loading messages: %w", err)
	}

	// Set messages
	b.Messages = messages

	return nil
}

// T translates a key
func (b *bundle) T(key string, params ...types.Params) string {
	// Check if messages are loaded and save them if they are not
	message, ok := b.Messages[key]
	if !ok {
		// Create new empty map if messages are not loaded
		b.Messages[key] = ""

		// Save messages
		err := b.SaveMessages()
		if err != nil {
			log.Println("Error saving messages:", err)
		}

		// Return empty string
		return ""
	}

	// Replace params
	if len(params) > 0 {
		message = b.ReplaceParams(message, params[0])
	}

	// Return message
	return message
}

// SaveMessages saves the messages to the translation directory
func (b *bundle) SaveMessages() error {
	// Marshal messages
	messagesBytes, err := json.MarshalIndent(b.Messages, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling messages: %w", err)
	}

	// Open file and write to it
	file, err := os.OpenFile(b.Language.TranslationDir+"/"+b.Language.Name+".json", os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening translation file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(messagesBytes)
	if err != nil {
		return fmt.Errorf("error writing to translation file: %w", err)
	}

	return nil
}

// ReplaceParams replaces params in a message
func (b *bundle) ReplaceParams(message string, params types.Params) string {
	for key, value := range params {
		message = b.ReplaceParam(message, key, value)
	}
	return message
}

// ReplaceParam replaces a param in a message
func (b *bundle) ReplaceParam(message, key, value string) string {
	return strings.ReplaceAll(message, "{"+key+"}", value)
}

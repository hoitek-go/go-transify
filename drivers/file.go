package drivers

import (
	"encoding/json"
	"fmt"
	"github.com/hoitek-go/go-transify/types"
	"io"
	"log"
	"os"
)

type File struct {
	Language types.Language
}

func (f *File) LoadMessages() (types.Messages, error) {
	log.Println("Loading translations for", f.Language.Name, "from file...")

	//
	var (
		translationDir = f.Language.TranslationDir
		filename       = translationDir + "/" + f.Language.Name + ".json"
	)

	// Check if translation directory exists and create it if it doesn't
	if _, err := os.Stat(translationDir); os.IsNotExist(err) {
		_, err := os.Create(translationDir)
		if err != nil {
			return nil, fmt.Errorf("error creating translation dir: %w", err)
		}
	}

	// Check if file exists and create it if it doesn't
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("error creating translation file: %w", err)
		}
	}

	// Read messages from file
	return f.readMessagesFromFile(filename)
}

func (f *File) readMessagesFromFile(filename string) (types.Messages, error) {
	// Open file and read it
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening translation file: %w", err)
	}
	defer file.Close()

	// Read file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading translation file: %w", err)
	}

	// Unmarshal file
	messages := make(types.Messages)
	err = json.Unmarshal(fileBytes, &messages)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling translation file: %w", err)
	}

	return messages, nil
}

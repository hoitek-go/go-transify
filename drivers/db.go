package drivers

import "github.com/hoitek-go/go-transify/types"

type DB struct {
	Language types.Language
}

func (d *DB) LoadMessages() (types.Messages, error) {
	return nil, nil
}

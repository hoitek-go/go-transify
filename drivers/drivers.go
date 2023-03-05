package drivers

import (
	"errors"
	"github.com/hoitek-go/go-transify/ports"
	"github.com/hoitek-go/go-transify/types"
)

func GetDriver(language types.Language) (ports.Driver, error) {
	switch language.Driver {
	case FileDriver:
		return &File{
			Language: language,
		}, nil
	case DBDriver:
		return &DB{
			Language: language,
		}, nil
	default:
		return nil, errors.New("driver not found")
	}
}

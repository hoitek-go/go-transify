package ports

import "github.com/hoitek-go/go-transify/types"

type Driver interface {
	LoadMessages() (types.Messages, error)
}

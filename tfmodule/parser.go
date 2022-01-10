package tfmodule

import (
	"errors"
	"strings"
)

// Parser represents a terraform module parser.
type Parser interface {
	Parse() (*Module, error)
}

func NewParser(source string) (Parser, error) {
	if strings.HasPrefix(source, "./") || strings.HasPrefix(source, "../") {
		return newLocalParser(source), nil
	} else {
		return nil, errors.New("Invalid source")
	}
}

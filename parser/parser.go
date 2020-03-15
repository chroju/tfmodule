package parser

import (
	"github.com/chroju/tfmodule/module"
	hcl "github.com/hashicorp/hcl/v2"
)

type Parser struct {
	source string
}

func NewParser(source string) *Parser {
	return &Parser{
		source: source,
	}
}

// ParseTfModule parses terraform module and returns module structs
func (p *Parser) ParseTfModule(source string) (*module.Module, hcl.Diagnostics) {
	//
	return nil, nil
}

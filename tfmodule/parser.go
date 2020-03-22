package tfmodule

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// Parser represents a terraform module parser.
type Parser struct {
	source string
}

// NewParser return a new parser with given module source path.
func NewParser() *Parser {
	return &Parser{}
}

// ParseTfModule parses terraform module and returns module structs
func (p *Parser) ParseTfModule(source string) (*Module, error) {
	p.source = source
	var variables []Variable
	err := filepath.Walk(p.source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() || !strings.HasSuffix(info.Name(), ".tf") {
				return nil
			}

			src, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			file, diags := hclwrite.ParseConfig(src, path, hcl.InitialPos)
			if diags.HasErrors() {
				return diags
			}

			body := file.Body()
			for _, block := range body.Blocks() {
				switch block.Type() {
				case "variable":
					variables = append(variables, parseVariable(block))
				}
			}

			return nil
		})
	if err != nil {
		return nil, err
	}
	module := NewModule(p.source)
	module.Variables = &variables

	return module, nil
}

func parseVariable(block *hclwrite.Block) Variable {
	variable := Variable{
		Name:    block.Labels()[0],
		Default: hclwrite.TokensForValue(cty.StringVal("")),
	}
	body := block.Body()
	for k, v := range body.Attributes() {
		switch k {
		case "type":
			var typeTokens hclwrite.Tokens
			for _, t := range v.Expr().BuildTokens(nil) {
				if t.Type != hclsyntax.TokenNewline {
					typeTokens = append(typeTokens, t)
				}
			}
			variable.Type = typeTokens
		case "default":
			variable.Default = v.Expr().BuildTokens(nil)
		case "description":
			description := string(v.Expr().BuildTokens(nil).Bytes())
			variable.Description = description[2 : len(description)-1]
		}
	}
	return variable
}

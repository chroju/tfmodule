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
	var variables []Variable
	var outputs []Output
	var resources []Resource

	p.source = source
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
				case "output":
					outputs = append(outputs, parseOutput(block))
				case "resource":
					resources = append(resources, parseResource(block))
				}
			}

			return nil
		})
	if err != nil {
		return nil, err
	}
	module := NewModule(p.source)
	module.Variables = &variables
	module.Outputs = &outputs
	module.Resources = &resources

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

func parseOutput(block *hclwrite.Block) Output {
	output := Output{
		Name:        block.Labels()[0],
		Description: "",
	}
	body := block.Body()
	for k, v := range body.Attributes() {
		switch k {
		case "value":
			var typeTokens hclwrite.Tokens
			for _, t := range v.Expr().BuildTokens(nil) {
				if t.Type != hclsyntax.TokenNewline {
					typeTokens = append(typeTokens, t)
				}
			}
			output.Value = typeTokens
		case "description":
			description := string(v.Expr().BuildTokens(nil).Bytes())
			output.Description = description[2 : len(description)-1]
		}
	}
	return output
}

func parseResource(block *hclwrite.Block) Resource {
	resource := Resource{
		Type: block.Labels()[0],
		Name: block.Labels()[1],
	}
	return resource
}

package tfmodule

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
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
func (p *Parser) ParseTfModule(source string) (*Module, hcl.Diagnostics) {
	//
	var variables []*Variable
	err := filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if strings.HasSuffix(info.Name(), ".tf") {
				rel, err := filepath.Rel(source, path)
				if err != nil {
					panic(err)
				}

				src, _ := ioutil.ReadFile(path)
				exp, diags := hclwrite.ParseConfig(src, rel, hcl.InitialPos)
				if diags.HasErrors() {
					for _, d := range diags {
						println(d.Summary)
					}
				}
				body := exp.Body()
				for _, block := range body.Blocks() {
					if block.Type() == "variable" {
						variableBody := block.Body()
						newVariable := &Variable{Name: block.Labels()[0]}
						for k, v := range variableBody.Attributes() {
							switch k {
							case "type":
								newVariable.Type = v.Expr().BuildTokens(nil)
							case "default":
								newVariable.Default = v.Expr().BuildTokens(nil)
							case "description":
								newVariable.Description = string(v.Expr().BuildTokens(nil).Bytes())
							}
						}
						if newVariable.Default == nil {
							newVariable.Default = hclwrite.TokensForValue(cty.StringVal(""))
						}
						variables = append(variables, newVariable)
					}
				}
			}
			return nil
		})
	if err != nil {
		return nil, nil
	}
	module := &Module{
		Variables: variables,
	}
	return module, nil
}

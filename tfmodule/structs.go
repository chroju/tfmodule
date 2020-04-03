package tfmodule

import (
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// Resource is terraform module resource
type Resource struct {
	Type hclwrite.Tokens `hcl:"type,label"`
	Name hclwrite.Tokens `hcl:"name,label"`
}

// Variable is terraform module variable
type Variable struct {
	Name        string          `hcl:"name,label"`
	Type        hclwrite.Tokens `hcl:"type,attr"`
	Description string          `hcl:"description,attr"`
	Default     hclwrite.Tokens `hcl:"default,attr"`
}

// Output is terraform module output value info
type Output struct {
	Name        string          `hcl:"name,label"`
	Value       hclwrite.Tokens `hcl:"value,attr"`
	Description string          `hcl:"description,attr"`
}

// Module is a struct to express terraform module
type Module struct {
	Name      string      `hcl:"name,label"`
	Variables *[]Variable `hcl:"variable,block"`
	Outputs   *[]Output   `hcl:"output,block"`
	Resources *[]Resource `hcl:"resource,block"`
	Source    string      `hcl:"source,attr"`
}

// NewModule returns a new module
func NewModule(source string) *Module {
	separetedSourcePath := strings.Split(source, "/")
	name := separetedSourcePath[len(separetedSourcePath)-1]
	return &Module{
		Source: source,
		Name:   name,
	}
}

// String returns module HCL expression
func (m *Module) String() string {
	return m.PrintModuleTemplate(false)
}

func (m *Module) PrintModuleTemplate(isMinimum bool) string {
	f := hclwrite.NewEmptyFile()

	rootBody := f.Body()
	moduleBlock := rootBody.AppendNewBlock("module", []string{m.Name})
	moduleBody := moduleBlock.Body()

	moduleBody.SetAttributeValue("source", cty.StringVal(m.Source))
	moduleBody.AppendNewline()

	for _, v := range *m.Variables {
		if isMinimum && !reflect.DeepEqual(v.Default, hclwrite.TokensForValue(cty.StringVal(""))) {
			continue
		}
		moduleBody.AppendUnstructuredTokens(v.GenerateComment())
		moduleBody.SetAttributeRaw(v.Name, v.Default)
	}

	for _, v := range *m.Outputs {
		if isMinimum {
			continue
		}
		rootBody.AppendNewline()
		outputBlock := rootBody.AppendNewBlock("output", []string{m.Name + "_" + v.Name})
		outputBody := outputBlock.Body()
		tokens := hclwrite.Tokens{
			{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte("module"),
			},
			{
				Type:  hclsyntax.TokenDot,
				Bytes: []byte("."),
			},
			{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte(m.Name),
			},
			{
				Type:  hclsyntax.TokenDot,
				Bytes: []byte("."),
			},
			{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte(v.Name),
			},
		}
		outputBody.SetAttributeRaw("value", tokens)
	}

	return string(f.Bytes())
}

func (v *Variable) GenerateComment() hclwrite.Tokens {
	tokens := hclwrite.Tokens{
		{
			Type:  hclsyntax.TokenSlash,
			Bytes: []byte("//"),
		},
		{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(v.Description),
		},
		{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		},
		{
			Type:  hclsyntax.TokenSlash,
			Bytes: []byte("//"),
		},
		{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte("type:"),
		},
	}
	tokens = append(tokens, v.Type...)
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte("\n"),
	})
	return tokens
}

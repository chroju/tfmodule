package tfmodule

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func TestParseTfModule(t *testing.T) {
	module1 := &Module{
		Variables: []*Variable{
			{
				Name:        "no_default",
				Default:     hclwrite.TokensForValue(cty.StringVal("")),
				Type:        hclwrite.TokensForValue(cty.StringVal("string")),
				Description: "no default description",
			},
			{
				Name: "object_type",
				Default: hclwrite.TokensForValue(cty.ObjectVal(map[string]cty.Value{
					"name":  cty.StringVal("default"),
					"count": cty.NumberIntVal(1),
				})),
				Type: hclwrite.Tokens{
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("object"),
					},
					{
						Type:  hclsyntax.TokenOParen,
						Bytes: []byte("("),
					},
					{
						Type:  hclsyntax.TokenOBrace,
						Bytes: []byte("{"),
					},
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("name"),
					},
					{
						Type:  hclsyntax.TokenEqual,
						Bytes: []byte("="),
					},
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("string"),
					},
					{
						Type:  hclsyntax.TokenComma,
						Bytes: []byte(","),
					},
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("count"),
					},
					{
						Type:  hclsyntax.TokenEqual,
						Bytes: []byte("="),
					},
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("number"),
					},
					{
						Type:  hclsyntax.TokenComma,
						Bytes: []byte(","),
					},
					{
						Type:  hclsyntax.TokenCBrace,
						Bytes: []byte("}"),
					},
					{
						Type:  hclsyntax.TokenCParen,
						Bytes: []byte(")"),
					},
				},
				Description: "object type description",
			},
		},
	}

	var tests = []struct {
		source string
		module *Module
	}{
		{"./test/module1", module1},
	}

	p := NewParser("")
	for _, test := range tests {
		if module, _ := p.ParseTfModule(test.source); module != test.module {
			t.Errorf("source %s: %s\nExpected: %s", test.source, module, test.module)
		}
	}
}

package tfmodule

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func TestParseTfModule(t *testing.T) {
	module1 := &Module{
		Name:   "module1",
		Source: "./test/module1",
		Variables: &[]Variable{
			{
				Name:    "no_default",
				Default: hclwrite.TokensForValue(cty.StringVal("")),
				// Default: nil,
				Type: hclwrite.Tokens{
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("string"),
					},
				},
				Description: "no default description",
			},
			{
				Name: "object_type",
				Default: hclwrite.Tokens{
					{
						Type:  hclsyntax.TokenOBrace,
						Bytes: []byte("{"),
					},
					{
						Type:  hclsyntax.TokenNewline,
						Bytes: []byte("\n"),
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
						Type:  hclsyntax.TokenOQuote,
						Bytes: []byte("\""),
					},
					{
						Type:  hclsyntax.TokenQuotedLit,
						Bytes: []byte("default"),
					},
					{
						Type:  hclsyntax.TokenCQuote,
						Bytes: []byte("\""),
					},
					{
						Type:  hclsyntax.TokenComma,
						Bytes: []byte(","),
					},
					{
						Type:  hclsyntax.TokenNewline,
						Bytes: []byte("\n"),
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
						Type:  hclsyntax.TokenNumberLit,
						Bytes: []byte("1"),
					},
					{
						Type:  hclsyntax.TokenNewline,
						Bytes: []byte("\n"),
					},
					{
						Type:  hclsyntax.TokenCBrace,
						Bytes: []byte("}"),
					},
				},
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
		{"./test/not_exist", (*Module)(nil)},
	}

	p := NewParser("")
	for _, test := range tests {
		m, _ := p.ParseTfModule(test.source)
		if m == nil {
			if test.module == nil {
				continue
			} else {
				t.Errorf("source %s: %s\nExpected: %s", test.source, m, test.module)
			}
		}
		if m.String() != test.module.String() {
			t.Errorf("source %s: %s\nExpected: %s", test.source, m, test.module)
		}
	}
}

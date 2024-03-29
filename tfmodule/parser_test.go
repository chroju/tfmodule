package tfmodule

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func TestLocalParser(t *testing.T) {
	module1 := &Module{
		Source: "./test/module1",
		Resources: []*Resource{
			{
				Name: "instance",
				Type: "aws_instance",
			},
		},
		Outputs: []*Output{
			{
				Name: "test",
				Value: hclwrite.Tokens{
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("aws_instance"),
					},
					{
						Type:  hclsyntax.TokenDot,
						Bytes: []byte("."),
					},
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("test"),
					},
					{
						Type:  hclsyntax.TokenDot,
						Bytes: []byte("."),
					},
					{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("arn"),
					},
				},
				Description: "test instance",
			},
		},
		Variables: []*Variable{
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
		{"hashicorp/consul/aws", (*Module)(nil)},
	}

	for _, test := range tests {
		p := newLocalParser(test.source)
		m, _ := p.Parse()
		if m == nil {
			if test.module == nil {
				continue
			} else {
				t.Errorf("source %s: %v\nExpected: %v", test.source, m, test.module)
			}
		}
		if m.String() != test.module.String() {
			t.Errorf("source %s: %v\nExpected: %v", test.source, m, test.module)
		}
	}
}

package tfmodule

import (
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type templatePrinter struct {
	Module       *Module
	IsNoDefaults bool
	IsNoOutputs  bool
}

func newTemplatePrinter(m *Module, isNoDefaults, isNoOutputs bool) *templatePrinter {
	return &templatePrinter{
		Module:       m,
		IsNoDefaults: isNoDefaults,
		IsNoOutputs:  isNoOutputs,
	}
}

func (p *templatePrinter) Print() (string, error) {
	f := hclwrite.NewEmptyFile()

	if p.Module.Name == "" {
		separetedSourcePath := strings.Split(p.Module.Source, "/")
		p.Module.Name = separetedSourcePath[len(separetedSourcePath)-1]
	}

	rootBody := f.Body()
	moduleBlock := rootBody.AppendNewBlock("module", []string{p.Module.Name})
	moduleBody := moduleBlock.Body()

	moduleBody.SetAttributeValue("source", cty.StringVal(p.Module.Source))
	moduleBody.AppendNewline()

	for _, v := range p.Module.Variables {
		if p.IsNoDefaults && !reflect.DeepEqual(v.Default, hclwrite.TokensForValue(cty.StringVal(""))) {
			continue
		}
		moduleBody.AppendUnstructuredTokens(p.generateVariableComment(v))
		moduleBody.SetAttributeRaw(v.Name, v.Default)
	}

	for _, v := range p.Module.Outputs {
		if p.IsNoOutputs {
			continue
		}
		rootBody.AppendNewline()
		outputBlock := rootBody.AppendNewBlock("output", []string{p.Module.Name + "_" + v.Name})
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
				Bytes: []byte(p.Module.Name),
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

	return string(f.Bytes()), nil
}

func (p *templatePrinter) generateVariableComment(v *Variable) hclwrite.Tokens {
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

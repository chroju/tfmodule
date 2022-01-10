package tfmodule

import (
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type TemplatePrinter struct {
	Module       *Module
	IsNoDefaults bool
	IsNoOutputs  bool
}

func NewTemplatePrinter(m *Module, isNoDefaults, isNoOutputs bool) *TemplatePrinter {
	return &TemplatePrinter{
		Module:       m,
		IsNoDefaults: isNoDefaults,
		IsNoOutputs:  isNoOutputs,
	}
}

func (p *TemplatePrinter) Print() (string, error) {
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
		moduleBody.AppendUnstructuredTokens(v.generateComment())
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

package module

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Resource is terraform module resource
type Resource struct {
	Type hclwrite.Tokens `hcl:"type,label"`
	Name hclwrite.Tokens `hcl:"name,label"`
}

// Variable is terraform module variable
type Variable struct {
	Name        hclwrite.Tokens `hcl:"name,label"`
	Type        hclwrite.Tokens `hcl:"type,label"`
	Description hclwrite.Tokens `hcl:"description,attr"`
	Default     hclwrite.Tokens `hcl:"default,attr"`
}

// Output is terraform module output value info
type Output struct {
	Name        hclwrite.Tokens `hcl:"name,label"`
	Description hclwrite.Tokens `hcl:"description,attr"`
}

// Module is a struct to express terraform module
type Module struct {
	Variables []*Variable     `hcl:"variable,block"`
	Outputs   []*Output       `hcl:"output,block"`
	Resources []*Resource     `hcl:"resource,block"`
	Source    hclwrite.Tokens `hcl:"source,attr"`
	Remain    hcl.Body        `hcl:",remain"`
}

// String returns module HCL expression
func (m *Module) String() string {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()
	moduleBlock := rootBody.AppendNewBlock("module", nil)
	moduleBody := moduleBlock.Body()

	moduleBody.SetAttributeRaw("source", m.Source)
	moduleBody.AppendNewline()

	for _, v := range m.Variables {
		moduleBody.SetAttributeRaw(string(v.Name.Bytes()), v.Default)
	}

	return string(f.Bytes())
}

package tfmodule

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Resource is terraform module resource
type Resource struct {
	Type string `hcl:"type,label"`
	Name string `hcl:"name,label"`
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
	Variables []*Variable `hcl:"variable,block"`
	Outputs   []*Output   `hcl:"output,block"`
	Resources []*Resource `hcl:"resource,block"`
	Source    string      `hcl:"source,attr"`
}

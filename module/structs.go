package module

import "github.com/hashicorp/hcl/v2"

// Resource is terraform module resource
type Resource struct {
	Type string `hcl:"type,label"`
	Name string `hcl:"name,label"`
}

// Variable is terraform module variable
type Variable struct {
	Name        string  `hcl:"name,label"`
	Description *string `hcl:"description,attr"`
	Default     *string `hcl:"default,attr"`
}

// Output is terraform module output value info
type Output struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description,attr"`
}

// Module is a struct to express terraform module
type Module struct {
	Variables []*Variable `hcl:"variable,block"`
	Outputs   []*Output   `hcl:"output,block"`
	Resources []*Resource `hcl:"resource,block"`
	Remain    hcl.Body    `hcl:",remain"`
}

func (*m Module) String() string {
	return ""
}

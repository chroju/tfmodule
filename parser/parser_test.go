package parser

import (
	"testing"
	"github.com/chroju/tfmodule/module"
)

func TestParseTfModule(t *testing.T) () {
	module1 := *module.Module{
		Variables: []*module.Variable{
			{
				Name: "no_default",
				Default: nil,
				Type: "string",
				Description: "no default description",
			},
			{
				Name: "object_type",
				Default: `{
  name  = "default",
  count = 1
}`,
				Type: "object({name=string,count=number}",
				Description: "object type description",
			},
		}
	}

	var tests = []struct {
		source string
		module module.Module
	}{
		{"./module1", module1},
	}

	for _, test := range tests {
		if module, _ := ParseTfModule(test.source); module != t.module {
			t.Errorf("source %s: %s\nExpected: %s", test.source, module, test.module)
		}
	}
}

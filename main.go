package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
)

type Variable struct {
	Name        string `hcl:"variable,label"`
	Description *string `hcl:"description,attr"`
	Default     *string `hcl:"default,attr"`
}

type Config struct {
	Variables []Variable `hcl:"variable,block"`
	Remain    hcl.Body   `hcl:",remain"`
}

func main() {
	var conf Config
	var vars []Variable

	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	parser := hclparse.NewParser()

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if strings.HasSuffix(info.Name(), ".tf") {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					panic(err)
				}
				f, parseDiags := parser.ParseHCLFile(rel)
				if parseDiags.HasErrors() {
					panic(parseDiags.Error())
				}

				decodeDiags := gohcl.DecodeBody(f.Body, nil, &conf)
				if decodeDiags.HasErrors() {
					panic(decodeDiags.Error())
				}
				vars = append(vars, conf.Variables ...)
			}
			return nil
		})
	if err != nil {
		panic(nil)
	}

	fmt.Println(*vars[0].Description)
}

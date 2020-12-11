package main

import (
	"github.com/actgardner/gogen-avro/v7/generator"
	"github.com/actgardner/gogen-avro/v7/generator/flat"
	"github.com/actgardner/gogen-avro/v7/parser"
	"github.com/actgardner/gogen-avro/v7/resolver"
	"io/ioutil"
)

func GenerateAvro(files []string, targetDir string) error {
	pkg := generator.NewPackage("avro", "")
	namespace := parser.NewNamespace(false)
	gen := flat.NewFlatPackageGenerator(pkg, false)

	for _, fileName := range files {
		schema, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}

		_, err = namespace.TypeForSchema(schema)
		if err != nil {
			return err
		}
	}

	for _, def := range namespace.Roots {
		if err := resolver.ResolveDefinition(def, namespace.Definitions); err != nil {
			return err
		}
	}

	for _, def := range namespace.Roots {
		err := gen.Add(def)
		if err != nil {
			return err
		}
	}

	err := pkg.WriteFiles(targetDir)
	if err != nil {
		return err
	}
	return nil
}

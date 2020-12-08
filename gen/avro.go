package main

import (
	"fmt"
	"github.com/actgardner/gogen-avro/v7/generator"
	"github.com/actgardner/gogen-avro/v7/generator/flat"
	"github.com/actgardner/gogen-avro/v7/parser"
	"github.com/actgardner/gogen-avro/v7/resolver"
	"io/ioutil"
	"os"
)

func GenerateAvro(files []string, targetDir string) error {

	var err error
	pkg := generator.NewPackage("avro", "")
	namespace := parser.NewNamespace(false)
	gen := flat.NewFlatPackageGenerator(pkg, false)

	for _, fileName := range files {
		schema, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %q - %v\n", fileName, err)
			return err
		}

		_, err = namespace.TypeForSchema(schema)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding schema for file %q - %v\n", fileName, err)
			return err
		}
	}

	for _, def := range namespace.Roots {
		if err := resolver.ResolveDefinition(def, namespace.Definitions); err != nil {
			fmt.Fprintf(os.Stderr, "Error resolving definition for type %q - %v\n", def.Name(), err)
			return err
		}
	}

	for _, def := range namespace.Roots {
		err = gen.Add(def)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating code for schema - %v\n", err)
			return err
		}
	}

	err = pkg.WriteFiles(targetDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing source files to directory %q - %v\n", targetDir, err)
		return err

	}
	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	schemaFile, err := os.Open("../schema.avsc")
	if err != nil {
		log.Panic(err)
	}

	name, err := loadTopLevelEntityNameFromSchema(schemaFile)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("TopLevel: %s\n", name)
	generateFunction(name)

	generateGeneratorMap(name)
}

func loadTopLevelEntityNameFromSchema(schema *os.File) (string, error) {
	bytes, err := ioutil.ReadAll(schema)
	m := make(map[string]interface{})
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return "", err
	}
	fmt.Println(m)
	return m["name"].(string), nil
}

func generateFunction(entity string) {
	upperCasedEntity := strings.Title(entity)
	lowerCasedEntity := strings.ToLower(entity)

	f := NewFile(lowerCasedEntity)
	f.ImportName("avro", "avro")
	f.Func().
		Id(fmt.Sprintf("generate%s", upperCasedEntity)).Params(Id("amount").Int()).Index().Interface().
		Block(
			Id("sliced").Op(":=").Make(Index().Interface(), Id("amount")),
			For(
				Id("i").Op(":=").Range().Id("sliced")).
				Block(Id("sliced").Index(Id("i")).Op("=").
					Id("randomize").Params(Qual("avro", fmt.Sprintf("New%s", upperCasedEntity)).Call())),
			Return(Id("sliced")),
		)

	f.Func().Id("randomize").Params(Id(lowerCasedEntity).Interface()).Interface().
		Block(
			Return(Id(lowerCasedEntity)),
		)

	fmt.Printf("%#v", f)
}

func generateGeneratorMap(entity string) {
	upperCasedEntity := strings.Title(entity)
	lowerCasedEntity := strings.ToLower(entity)
	f := NewFile("generated")
	f.ImportName(fmt.Sprintf("drtest/generated/%s", lowerCasedEntity), lowerCasedEntity)
	f.ImportName("errors", "errors")

	f.Func().Id("Generate").Params(
		Id("structName").String(),
		Id("amount").Int()).
		Call(Index().Interface(), Id("error")).
		Block(
			Switch(Id("structName").
				Block(
					Case(Lit(lowerCasedEntity)).
						Block(
							Return(Qual(fmt.Sprintf("drtest/generated/%s", lowerCasedEntity), fmt.Sprintf("Generate%s", upperCasedEntity)).Call(Id("amount")), Nil()),
						),
					Default().Block(
						Return(Nil(), Qual("errors", "New").Call(Lit("struct not found"))))),
			),
		)

	fmt.Printf("%#v", f)
}

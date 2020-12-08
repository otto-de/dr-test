package main

import (
	"encoding/json"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	generatedDirName, err := filepath.Abs("../generated")
	if err != nil {
		log.Panic(err)
	}

	schemaFile, err := os.Open("../schema.avsc")
	if err != nil {
		log.Panic(err)
	}

	name, err := loadTopLevelEntityNameFromSchema(schemaFile)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("TopLevel: %s\n", name)

	err = cleanDirectory(generatedDirName)
	if err != nil {
		log.Panic(err)
	}

	lowerCasedEntity := strings.ToLower(name)

	code := generateGeneratorMap(name)
	err = writeFile("generator", code, generatedDirName)
	if err != nil {
		log.Panic(err)
	}

	code = generateFunction(name)
	err = writeFile(lowerCasedEntity, code, generatedDirName + "/" + lowerCasedEntity)
	if err != nil {
		log.Panic(err)
	}

	avroTargetDir := generatedDirName + "/" + lowerCasedEntity + "/avro"
	err = os.MkdirAll(avroTargetDir, 0755)
	if err != nil {
		log.Panic(err)
	}

	err = GenerateAvro([]string{"../schema.avsc"}, avroTargetDir)
	if err != nil {
		log.Panic(err)
	}
}

func writeFile(fileName string, content string, parentDir string) error {
	err := os.Mkdir(parentDir, 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(parentDir + "/" + fileName + ".go", []byte(content), 0644)
}

func cleanDirectory(dirToClean string) error {
	return os.RemoveAll(dirToClean)
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

func generateFunction(entity string) string {
	upperCasedEntity := strings.Title(entity)
	lowerCasedEntity := strings.ToLower(entity)
	packageName := fmt.Sprintf("drtest/generated/%s/avro", lowerCasedEntity)

	f := NewFile(lowerCasedEntity)
	f.ImportName(packageName, "avro")
	f.Func().
		Id(fmt.Sprintf("Generate%s", upperCasedEntity)).Params(Id("amount").Int()).Index().Interface().
		Block(
			Id("sliced").Op(":=").Make(Index().Interface(), Id("amount")),
			For(
				Id("i").Op(":=").Range().Id("sliced")).
				Block(Id("sliced").Index(Id("i")).Op("=").
					Id("randomize").Params(Qual(packageName, fmt.Sprintf("New%s", upperCasedEntity)).Call())),
			Return(Id("sliced")),
		)

	f.Func().Id("randomize").Params(Id(lowerCasedEntity).Interface()).Interface().
		Block(
			Return(Id(lowerCasedEntity)),
		)

	return fmt.Sprintf("%#v", f)
}

func generateGeneratorMap(entity string) string {
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
					Case(Lit(upperCasedEntity)).
						Block(
							Return(Qual(fmt.Sprintf("drtest/generated/%s", lowerCasedEntity), fmt.Sprintf("Generate%s", upperCasedEntity)).Call(Id("amount")), Nil()),
						),
					Default().Block(
						Return(Nil(), Qual("errors", "New").Call(Lit("struct not found"))))),
			),
		)

	return fmt.Sprintf("%#v", f)
}

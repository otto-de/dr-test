package main

import (
	"encoding/json"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var targetDir string

	app := &cli.App{
		Name:      "gen",
		Usage:     "Generate the code you need",
		UsageText: "gen [--target-dir dir] schema-files...",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "target-dir",
				Value:       "generated",
				Aliases:     []string{"t"},
				Destination: &targetDir,
			},
		},
		Action: func(c *cli.Context) error {
			return run(targetDir, c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(targetDir string, c *cli.Context) error {
	var schemaFiles []string

	fmt.Printf("Args: %s", c.Args())
	if c.NArg() == 0 {
		cli.ShowAppHelpAndExit(c, 1)
	}

	for i := range c.Args().Slice() {
		schemaFiles = append(schemaFiles, c.Args().Get(i))
	}

	targetDirName, err := filepath.Abs(targetDir)
	if err != nil {
		return err
	}

	err = cleanTargetDir(targetDirName)
	log.Printf("Generating code to directory %s", targetDirName)

	var records []string
	for i := range schemaFiles {
		schemaFile, err := os.Open(schemaFiles[i])
		if err != nil {
			return err
		}

		record, err := loadTopLevelEntityNameFromSchema(schemaFile)
		if err != nil {
			return err
		}

		records = append(records, record)

		fmt.Printf("Root record for schema %s -> %s\n", schemaFile.Name(), record)

		err = initSchemaDirectory(targetDirName, record)
		if err != nil {
			return err
		}

		lowerCasedEntity := strings.ToLower(record)

		code := generateFunction(record)
		err = writeFile(lowerCasedEntity+".go", code, targetDirName+"/"+lowerCasedEntity)
		if err != nil {
			return err
		}

	}

	avroTargetDir, err := initAvroDirectory(targetDirName)
	if err != nil {
		return err
	}

	err = GenerateAvro(schemaFiles, avroTargetDir)
	if err != nil {
		return err
	}

	code := generateGeneratorMap(records)
	err = writeFile("generator.go", code, targetDirName)
	if err != nil {
		return err
	}

	return nil
}

func writeFile(fileName string, content string, parentDir string) error {
	return ioutil.WriteFile(parentDir+"/"+fileName, []byte(content), 0644)
}

func cleanTargetDir(targetDir string) error {
	return os.RemoveAll(targetDir)
}

func initAvroDirectory(targetDir string) (string, error) {
	dir := targetDir + "/avro"
	log.Printf("Creating directory %s", dir)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func initSchemaDirectory(targetDir, entityName string) error {
	lowerCasedEntityName := strings.ToLower(entityName)
	dir := targetDir + "/" + lowerCasedEntityName
	log.Printf("Creating directory %s", dir)
	return os.MkdirAll(dir, 0755)
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
	return m["name"].(string), nil
}

func generateFunction(entity string) string {
	upperCasedEntity := strings.Title(entity)
	lowerCasedEntity := strings.ToLower(entity)
	packageName := "drtest/generated/avro"
	pkgPackageName := fmt.Sprintf("drtest/randomize/pkg")

	f := NewFile(lowerCasedEntity)
	f.ImportName(packageName, "avro")
	f.ImportName(pkgPackageName, "pkg")
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
			Return(Qual(pkgPackageName, "RandomizeWithDefaults").Call(Id(lowerCasedEntity))),
		)

	return fmt.Sprintf("%#v", f)
}

func generateGeneratorMap(records []string) string {
	f := NewFile("generated")
	f.ImportName("errors", "errors")

	var caseStatements []Code
	var recordNames []Code

	for i := range records {
		record := records[i]
		lowerCasedRecord := strings.ToLower(record)
		f.ImportName(fmt.Sprintf("drtest/generated/%s", lowerCasedRecord), lowerCasedRecord)

		caseStatementForRecord := generateCaseStatement(lowerCasedRecord)
		caseStatements = append(caseStatements, caseStatementForRecord)

		recordNames = append(recordNames, Lit(lowerCasedRecord))
	}

	caseStatements = append(caseStatements, Default().Block(
		Return(Nil(), Qual("errors", "New").Call(Lit("record not found")))))

	f.Func().Id("Generate").Params(
		Id("recordName").String(),
		Id("amount").Int()).
		Call(Index().Interface(), Id("error")).
		Block(
			Switch(Id("recordName").
				Block(caseStatements...),
			),
		)

	f.Func().Id("GetRecordNames").Params().Index().String().Block(
		Return(Index().String().Values(recordNames...)),
	)

	return fmt.Sprintf("%#v", f)
}

func generateCaseStatement(record string) *Statement {
	upperCasedRecord := strings.Title(record)
	return Case(Lit(record)).
		Block(
			Return(Qual(fmt.Sprintf("drtest/generated/%s", record), fmt.Sprintf("Generate%s", upperCasedRecord)).Call(Id("amount")), Nil()),
		)
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	dynamodb = "dynamodb"
	es       = "elasticsearch"
	http     = "http"
	lambda   = "lambda"
	none     = "none"
)

const (
	query        = "query"
	mutation     = "mutation"
	subscription = "subscription"
)

type AppsyncYaml struct {
	Setting *Setting `yaml:"setting"`
}

type Setting struct {
	DataSources []*DataSource `yaml:"dataSources"`
}

type DataSource struct {
	Type string `yaml:"type"`
	Name string `yaml:"Name"`
}

type Fields struct {
	Fields []*Field `yaml:"fields"`
}

type Field struct {
	Type       string `yaml:"type"`
	Field      string `yaml:"field"`
	Request    string `yaml:"request"`
	Response   string `yaml:"response"`
	DataSource string `yaml:"dataSource"`
}

func parseDirectory(path string) []string {
	var directories []string
	validField := regexp.MustCompile(`.+/(mutation|query|subscription).+/(req|res).vtl`)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !validField.MatchString(path) {
			return nil
		}

		directories = append(directories, path)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return directories
}

func getFields(paths []string) *Fields {
	var fields []*Field
	for _, path := range paths {
		var req, res string
		fieldDetail := strings.Split(path, "/")
		isFielsExist := false

		if fieldDetail[4] == "req.vtl" {
			req = fieldDetail[1] + "/" + fieldDetail[2] + "/" + fieldDetail[3] + "/" + fieldDetail[4]
		}

		if fieldDetail[4] == "res.vtl" {
			res = fieldDetail[1] + "/" + fieldDetail[2] + "/" + fieldDetail[3] + "/" + fieldDetail[4]
			beforeItem := len(fields) - 1

			if len(fields) != 0 &&
				fields[beforeItem].DataSource == fieldDetail[1] &&
				fields[beforeItem].Type == strings.Title(fieldDetail[2]) &&
				fields[beforeItem].Field == fieldDetail[3] &&
				fields[beforeItem].Request != "" {

				fields[beforeItem].Response = res
				isFielsExist = true
			}
		}

		if !isFielsExist {
			fields = append(fields, &Field{
				DataSource: fieldDetail[1],
				Type:       strings.Title(fieldDetail[2]),
				Field:      fieldDetail[3],
				Request:    req,
				Response:   res,
			})
		}
	}

	// insert default request and response
	for _, field := range fields {
		if field.Request == "" {
			field.Request = "req.vtl"
		}

		if field.Response == "" {
			field.Response = "res.vtl"
		}
	}

	return &Fields{Fields: fields}
}

func writeFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	argDir := flag.String("dir", "", "a string of directory to parse to generate mapping template")
	argOut := flag.String("out", "", "a string of file to write generate mapping template file")
	flag.Parse()

	directories := parseDirectory(*argDir)

	fields := getFields(directories)
	results, err := yaml.Marshal(fields)
	if err != nil {
		panic(err)
	}
	for _, field := range fields.Fields {
		fmt.Println(field)
	}

	writeFile(*argOut, results)
}

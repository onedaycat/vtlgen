package main

import (
	"flag"
	"io/ioutil"

	"github.com/onedaycat/vtlgen"
	yaml "gopkg.in/yaml.v2"
)

const (
	mappingTemplateFilename = "mappingTemplates.yml"
)

func main() {
	argDir := flag.String("dir", "", "select directory that have datasource_generate.yml and folder mappingTemplates.")
	flag.Parse()

	mappingTemplates := vtlgen.GenerateMappingTemplatesAndFunctions(*argDir)
	datasource := vtlgen.GenerateDatasources(*argDir)

	results, err := yaml.Marshal(mappingTemplates)
	if err != nil {
		panic(err)
	}

	// testdata/mappingTemplates.yml
	err = ioutil.WriteFile(*argDir+mappingTemplateFilename, results, 0644)
	if err != nil {
		panic(err)
	}
}

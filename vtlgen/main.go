package main

import (
	"flag"
	"io/ioutil"

	"github.com/onedaycat/vtlgen"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	argDir := flag.String("dir", "", "a string of directory to parse to generate mapping template")
	argOut := flag.String("out", "", "a string of file to write generate mapping template file")
	flag.Parse()

	mappingTemplates := vtlgen.GenerateMappingTemplates(*argDir)

	results, err := yaml.Marshal(mappingTemplates)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*argOut, results, 0644)
	if err != nil {
		panic(err)
	}
}

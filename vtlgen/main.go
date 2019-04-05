package main

import (
	"flag"
	"io/ioutil"

	v "github.com/onedaycat/vtlgen"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	argDir := flag.String("dir", "", "select directory that have datasource_generate.yml and folder mappingTemplates.")
	flag.Parse()

	if *argDir == v.EmptyString {
		*argDir = v.DotString
	}

	mt := v.GenerateMappingTemplates(*argDir)

	mtYML, err := yaml.Marshal(mt)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*argDir+v.PathDelim+v.ResolverFilename, mtYML, 0644)
	if err != nil {
		panic(err)
	}

	d := v.GenerateDatasources(*argDir)

	dYML, err := yaml.Marshal(d)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*argDir+v.PathDelim+v.DatasourceFilename, dYML, 0644)
	if err != nil {
		panic(err)
	}
}

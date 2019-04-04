package vtlgen

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	datasource        = 0
	graphqlType       = 1
	field             = 2
	requestOrResponse = 3
)

const (
	emptyString    = ""
	pathDelim      = "/"
	reqFilename    = "req.vtl"
	resFilename    = "res.vtl"
	beforeFilename = "before.vtl"
	afterFilename  = "after.vtl"
	pathResolver   = "/resolver"
	configFilename = "config.yml"
)

// meaning: (datasource)/(graphqlType)/(field)/(requestOrResponse)
var isValidFilename = regexp.MustCompile("(.+)/config.yml")

func GenerateMappingTemplates(parseDirectory string) *MappingTemplates {
	var templates []*Template
	parseDirectory = path.Clean(parseDirectory)

	err := filepath.Walk(parseDirectory+pathResolver, func(path string, info os.FileInfo, err error) error {
		if !isValidFilename.MatchString(path) {
			return nil
		}

		config, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		template := Template{}
		err = yaml.Unmarshal(config, &template)
		if err != nil {
			panic(err)
		}

		// normal or pipeline resolver
		path = strings.Replace(path, parseDirectory+pathDelim, emptyString, 1)

		if template.DataSource != emptyString {
			req := strings.Replace(path, configFilename, reqFilename, 1)
			res := strings.Replace(path, configFilename, resFilename, 1)
			template.Request = req
			template.Response = res
		} else {
			req := strings.Replace(path, configFilename, beforeFilename, 1)
			res := strings.Replace(path, configFilename, afterFilename, 1)
			template.Kind = "PIPELINE"
			template.Request = req
			template.Response = res
		}

		templates = append(templates, &template)

		return nil
	})
	if err != nil {
		panic(err)
	}

	if len(templates) == 0 {
		panic("not match directory structure to generate mappingtemplates")
	}

	return &MappingTemplates{Templates: templates}
}

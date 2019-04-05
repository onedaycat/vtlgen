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
	pathResolver   = "/resolver"
	pathFunctions  = "/function"
	reqFilename    = "req.vtl"
	resFilename    = "res.vtl"
	beforeFilename = "before.vtl"
	afterFilename  = "after.vtl"
	configFilename = "config.yml"
	pipeline       = "PIPELINE"
)

// meaning: (datasource)/(graphqlType)/(field)/(requestOrResponse)
var isValidFilename = regexp.MustCompile("(.+)/config.yml")

func GenerateMappingTemplatesAndFunctions(parseDirectory string) *MappingTemplates {
	var err error
	var templates []*Template
	var functions []*Function
	parseDirectory = path.Clean(parseDirectory)

	// resolver
	err = filepath.Walk(parseDirectory+pathResolver, func(path string, info os.FileInfo, err error) error {
		if !isValidFilename.MatchString(path) {
			return nil
		}

		config, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		template := &Template{}
		err = yaml.Unmarshal(config, template)
		if err != nil {
			panic(err)
		}

		// normal or pipeline resolver
		path = strings.Replace(path, parseDirectory+pathDelim, emptyString, 1)

		if template.Datasource != emptyString {
			req := strings.Replace(path, configFilename, reqFilename, 1)
			res := strings.Replace(path, configFilename, resFilename, 1)
			template.Request = req
			template.Response = res
		} else {
			req := strings.Replace(path, configFilename, beforeFilename, 1)
			res := strings.Replace(path, configFilename, afterFilename, 1)
			template.Kind = pipeline
			template.Request = req
			template.Response = res
		}

		templates = append(templates, template)

		return nil
	})
	if err != nil {
		panic(err)
	}

	// function
	err = filepath.Walk(parseDirectory+pathFunctions, func(path string, info os.FileInfo, err error) error {
		if !isValidFilename.MatchString(path) {
			return nil
		}

		config, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		fn := Function{}
		err = yaml.Unmarshal(config, &fn)
		if err != nil {
			panic(err)
		}

		// normal or pipeline resolver
		path = strings.Replace(path, parseDirectory+pathDelim, emptyString, 1)
		req := strings.Replace(path, configFilename, reqFilename, 1)
		res := strings.Replace(path, configFilename, resFilename, 1)
		fn.Request = req
		fn.Response = res

		functions = append(functions, &fn)

		return nil
	})
	if err != nil {
		panic(err)
	}

	if len(templates) == 0 {
		panic("not match directory structure to generate mappingtemplates")
	}

	return &MappingTemplates{
		Templates: templates,
		Functions: functions,
	}
}

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

var isValidFilename = regexp.MustCompile("(.+)/config.yml")

func GenerateMappingTemplates(parseDirectory string) *MappingTemplatesGenerated {
	var err error
	var templates []*Template
	var functions []*Function
	parseDirectory = path.Clean(parseDirectory)

	// resolver
	err = filepath.Walk(parseDirectory+PathMappingTemplates+PathResolver, func(fpath string, info os.FileInfo, err error) error {
		if !isValidFilename.MatchString(fpath) {
			return nil
		}

		config, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(err)
		}

		template := &Template{}
		err = yaml.Unmarshal(config, template)
		if err != nil {
			panic(err)
		}

		if parseDirectory == DotString {
			fpath = strings.Replace(fpath, "mapping-templates/", EmptyString, 1)
		} else {
			fpath = strings.Replace(fpath, parseDirectory+PathMappingTemplates+PathDelim, EmptyString, 1)
		}

		// normal or pipeline resolver
		if template.Datasource != EmptyString {
			req := strings.Replace(fpath, ConfigFilename, ReqFilename, 1)
			res := strings.Replace(fpath, ConfigFilename, ResFilename, 1)
			template.Request = req
			template.Response = res
		} else {
			req := strings.Replace(fpath, ConfigFilename, BeforeFilename, 1)
			res := strings.Replace(fpath, ConfigFilename, AfterFilename, 1)
			template.Kind = PipelineString
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
	err = filepath.Walk(parseDirectory+PathMappingTemplates+PathFunctions, func(fpath string, info os.FileInfo, err error) error {
		if !isValidFilename.MatchString(fpath) {
			return nil
		}

		config, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(err)
		}

		fn := Function{}
		err = yaml.Unmarshal(config, &fn)
		if err != nil {
			panic(err)
		}

		if parseDirectory == DotString {
			fpath = strings.Replace(fpath, "mapping-templates/", EmptyString, 1)
		} else {
			fpath = strings.Replace(fpath, parseDirectory+PathMappingTemplates+PathDelim, EmptyString, 1)
		}
		req := strings.Replace(fpath, ConfigFilename, ReqFilename, 1)
		res := strings.Replace(fpath, ConfigFilename, ResFilename, 1)
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

	return &MappingTemplatesGenerated{
		Templates: templates,
		Functions: functions,
	}
}

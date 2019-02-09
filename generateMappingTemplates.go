package vtlgen

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	DataSource        = 0
	Type              = 1
	Field             = 2
	RequestOrResponse = 3
)

func GenerateMappingTemplates(parseDirectory string) *MappingTemplates {
	var mappingTemplates []*Template

	parseDirectory = path.Clean(parseDirectory)
	cutParseDirectory := regexp.MustCompile(parseDirectory + "/")
	// meaning: (datasource)/(type)/(field)/(req|res)
	isValidFilename := regexp.MustCompile("(.+)/(.+)/(.+)/(req|res).vtl")

	err := filepath.Walk(parseDirectory, func(path string, info os.FileInfo, err error) error {
		var req, res string

		isMappingTemplatesExist := false

		if !isValidFilename.MatchString(path) {
			return nil
		}

		path = cutParseDirectory.ReplaceAllString(path, "")
		templateDetail := strings.Split(path, "/")

		if templateDetail[RequestOrResponse] == "req.vtl" {
			req = templateDetail[DataSource] + "/" + templateDetail[Type] + "/" + templateDetail[Field] + "/" + templateDetail[RequestOrResponse]
		}

		if templateDetail[RequestOrResponse] == "res.vtl" {
			res = templateDetail[DataSource] + "/" + templateDetail[Type] + "/" + templateDetail[Field] + "/" + templateDetail[RequestOrResponse]
			beforeItem := len(mappingTemplates) - 1

			if len(mappingTemplates) != 0 &&
				mappingTemplates[beforeItem].DataSource == templateDetail[DataSource] &&
				mappingTemplates[beforeItem].Type == strings.Title(templateDetail[Type]) &&
				mappingTemplates[beforeItem].Field == templateDetail[Field] &&
				mappingTemplates[beforeItem].Request != "" {

				mappingTemplates[beforeItem].Response = res
				isMappingTemplatesExist = true
			}
		}

		if !isMappingTemplatesExist {
			mappingTemplates = append(mappingTemplates, &Template{
				DataSource: templateDetail[DataSource],
				Type:       strings.Title(templateDetail[Type]),
				Field:      templateDetail[Field],
				Request:    req,
				Response:   res,
			})
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	if len(mappingTemplates) == 0 {
		panic("not match directory structure to generate mappingtemplates")
	}

	// insert default request and response
	for _, template := range mappingTemplates {
		if template.Request == "" {
			template.Request = "req.vtl"
		}

		if template.Response == "" {
			template.Response = "res.vtl"
		}
	}

	return &MappingTemplates{MappingTemplates: mappingTemplates}
}

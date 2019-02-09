package vtlgen

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	datasource        = 0
	graphqlType       = 1
	field             = 2
	requestOrResponse = 3
)

// meaning: (datasource)/(type)/(field)/(req|res)
var IsValidFilename = regexp.MustCompile("(.+)/(.+)/(.+)/(req|res).vtl")

func GenerateMappingTemplates(parseDirectory string) *MappingTemplates {
	var mappingTemplates []*Template

	parseDirectory = path.Clean(parseDirectory)

	err := filepath.Walk(parseDirectory, func(path string, info os.FileInfo, err error) error {
		var req, res string

		isMappingTemplatesExist := false

		if !IsValidFilename.MatchString(path) {
			return nil
		}

		path = strings.Replace(path, parseDirectory+"/", "", 1)
		templateDetail := strings.Split(path, "/")

		if templateDetail[requestOrResponse] == "req.vtl" {
			req = templateDetail[datasource] + "/" + templateDetail[graphqlType] + "/" + templateDetail[field] + "/" + templateDetail[requestOrResponse]
		}

		if templateDetail[requestOrResponse] == "res.vtl" {
			res = templateDetail[datasource] + "/" + templateDetail[graphqlType] + "/" + templateDetail[field] + "/" + templateDetail[requestOrResponse]
			beforeItem := len(mappingTemplates) - 1

			if len(mappingTemplates) != 0 &&
				mappingTemplates[beforeItem].DataSource == templateDetail[datasource] &&
				mappingTemplates[beforeItem].GraphqlType == strings.Title(templateDetail[graphqlType]) &&
				mappingTemplates[beforeItem].Field == templateDetail[field] &&
				mappingTemplates[beforeItem].Request != "" {

				mappingTemplates[beforeItem].Response = res
				isMappingTemplatesExist = true
			}
		}

		if !isMappingTemplatesExist {
			mappingTemplates = append(mappingTemplates, &Template{
				DataSource:  templateDetail[datasource],
				GraphqlType: strings.Title(templateDetail[graphqlType]),
				Field:       templateDetail[field],
				Request:     req,
				Response:    res,
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

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

const (
	emptyString = ""
	pathDelim   = "/"
	reqFilename = "req.vtl"
	resFilename = "res.vtl"
)

// meaning: (datasource)/(graphqlType)/(field)/(requestOrResponse)
var isValidFilename = regexp.MustCompile("(.+)/(.+)/(.+)/(req|res).vtl")

func GenerateMappingTemplates(parseDirectory string) *MappingTemplates {
	var templates []*Template
	parseDirectory = path.Clean(parseDirectory)

	err := filepath.Walk(parseDirectory, func(path string, info os.FileInfo, err error) error {
		var req, res string
		isTemplatesExist := false

		if !isValidFilename.MatchString(path) {
			return nil
		}

		path = strings.Replace(path, parseDirectory+pathDelim, emptyString, 1)
		templateDetail := strings.Split(path, pathDelim)

		if templateDetail[requestOrResponse] == reqFilename {
			req = templateDetail[datasource] + pathDelim + templateDetail[graphqlType] + pathDelim + templateDetail[field] + pathDelim + templateDetail[requestOrResponse]
		}

		if templateDetail[requestOrResponse] == resFilename {
			res = templateDetail[datasource] + pathDelim + templateDetail[graphqlType] + pathDelim + templateDetail[field] + pathDelim + templateDetail[requestOrResponse]
			beforeItem := len(templates) - 1

			if len(templates) != 0 &&
				templates[beforeItem].DataSource == templateDetail[datasource] &&
				templates[beforeItem].GraphqlType == strings.Title(templateDetail[graphqlType]) &&
				templates[beforeItem].Field == templateDetail[field] &&
				templates[beforeItem].Request != emptyString {

				templates[beforeItem].Response = res
				isTemplatesExist = true
			}
		}

		if !isTemplatesExist {
			templates = append(templates, &Template{
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

	if len(templates) == 0 {
		panic("not match directory structure to generate mappingtemplates")
	}

	// insert default request and response
	for _, template := range templates {
		if template.Request == emptyString {
			template.Request = reqFilename
		}

		if template.Response == emptyString {
			template.Response = resFilename
		}
	}

	return &MappingTemplates{Templates: templates}
}

package vtlgen

import "strings"

func GenerateMappingTemplates(paths []string) *Fields {
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

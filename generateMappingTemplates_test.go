package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMappingTemplates(t *testing.T) {
	expect := &MappingTemplates{
		MappingTemplates: []*Template{
			&Template{
				Type:       "Mutation",
				Field:      "createProduct",
				Request:    "haloLambda/mutation/createProduct/req.vtl",
				Response:   "haloLambda/mutation/createProduct/res.vtl",
				DataSource: "haloLambda",
			},
			&Template{
				Type:       "Query",
				Field:      "product",
				Request:    "haloLambda/query/product/req.vtl",
				Response:   "res.vtl",
				DataSource: "haloLambda",
			},
			&Template{
				Type:       "Subscription",
				Field:      "product",
				Request:    "req.vtl",
				Response:   "haloLambda/subscription/product/res.vtl",
				DataSource: "haloLambda",
			},
		},
	}

	mappingTemplates := GenerateMappingTemplates("./mapping-templates")

	require.Equal(t, expect, mappingTemplates)
}

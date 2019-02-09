package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessGenerateMappingTemplates(t *testing.T) {
	expect := &MappingTemplates{
		MappingTemplates: []*Template{
			&Template{
				GraphqlType: "Mutation",
				Field:       "createProduct",
				Request:     "haloLambda/mutation/createProduct/req.vtl",
				Response:    "haloLambda/mutation/createProduct/res.vtl",
				DataSource:  "haloLambda",
			},
			&Template{
				GraphqlType: "Namespace",
				Field:       "languages",
				Request:     "haloLambda/namespace/languages/req.vtl",
				Response:    "haloLambda/namespace/languages/res.vtl",
				DataSource:  "haloLambda",
			},
			&Template{
				GraphqlType: "Query",
				Field:       "product",
				Request:     "haloLambda/query/product/req.vtl",
				Response:    "res.vtl",
				DataSource:  "haloLambda",
			},
			&Template{
				GraphqlType: "Subscription",
				Field:       "product",
				Request:     "req.vtl",
				Response:    "haloLambda/subscription/product/res.vtl",
				DataSource:  "haloLambda",
			},
		},
	}

	mappingTemplates := GenerateMappingTemplates("testdata/mapping-templates")

	require.Equal(t, expect, mappingTemplates)
}

func TestNotMatchGenerateMappingTemplates(t *testing.T) {
	require.PanicsWithValue(t, "not match directory structure to generate mappingtemplates", func() { GenerateMappingTemplates("./somethingWrong") })
}

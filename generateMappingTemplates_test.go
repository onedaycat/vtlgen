// +build integration

package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessGenerateMappingTemplates(t *testing.T) {
	expect := &MappingTemplates{
		Templates: []*Template{
			&Template{
				GraphqlType: "Mutation",
				Field:       "createProduct",
				Request:     "resolver/mutation.createProduct/req.vtl",
				Response:    "resolver/mutation.createProduct/res.vtl",
				DataSource:  "productLambda",
			},
			&Template{
				GraphqlType: "Mutation",
				Field:       "languages",
				Request:     "resolver/mutation.languages/req.vtl",
				Response:    "resolver/mutation.languages/res.vtl",
				DataSource:  "haloLambda",
			},
			&Template{
				GraphqlType: "Namespace",
				Field:       "languages",
				Request:     "resolver/namespace.languages/before.vtl",
				Response:    "resolver/namespace.languages/after.vtl",
				Kind:        "PIPELINE",
				Functions: []string{
					"oneFunction",
					"twoFunction",
				},
			},
			&Template{
				GraphqlType: "Query",
				Field:       "product",
				Request:     "resolver/query.product/req.vtl",
				Response:    "resolver/query.product/res.vtl",
				DataSource:  "productLambda",
			},
			&Template{
				GraphqlType: "Subscription",
				Field:       "product",
				Request:     "resolver/subscription.product/req.vtl",
				Response:    "resolver/subscription.product/res.vtl",
				DataSource:  "productLambda",
			},
		},
	}

	mappingTemplates := GenerateMappingTemplates("testdata/mapping-templates2")
	require.Equal(t, expect, mappingTemplates)
}

func TestNotMatchGenerateMappingTemplates(t *testing.T) {
	require.PanicsWithValue(t, "not match directory structure to generate mappingtemplates", func() { GenerateMappingTemplates("./somethingWrong") })
}

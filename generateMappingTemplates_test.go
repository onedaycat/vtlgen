// +build integration

package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessDatasourceGenerateMappingTemplates(t *testing.T) {
	expect := &MappingTemplatesGenerated{
		Templates: []*Template{
			{
				GraphqlType: "Mutation",
				Field:       "createProduct",
				Request:     "resolver/mutation.createProduct/req.vtl",
				Response:    "resolver/mutation.createProduct/res.vtl",
				Datasource:  "productLambda",
			},
			{
				GraphqlType: "Mutation",
				Field:       "languages",
				Request:     "resolver/mutation.languages/req.vtl",
				Response:    "resolver/mutation.languages/res.vtl",
				Datasource:  "haloLambda",
			},
			{
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
			{
				GraphqlType: "Query",
				Field:       "product",
				Request:     "resolver/query.product/req.vtl",
				Response:    "resolver/query.product/res.vtl",
				Datasource:  "productLambda",
			},
			{
				GraphqlType: "Subscription",
				Field:       "product",
				Request:     "resolver/subscription.product/req.vtl",
				Response:    "resolver/subscription.product/res.vtl",
				Datasource:  "productLambda",
			},
		},
		Functions: []*Function{
			{
				Name:       "oneFunction",
				Datasource: "productLambda",
				Request:    "function/oneFunction/req.vtl",
				Response:   "function/oneFunction/res.vtl",
			},
			{
				Name:       "twoFunction",
				Datasource: "haloLambda",
				Request:    "function/twoFunction/req.vtl",
				Response:   "function/twoFunction/res.vtl",
			},
		},
	}

	mappingTemplates := GenerateMappingTemplates("testdata")
	require.Equal(t, expect, mappingTemplates)
}

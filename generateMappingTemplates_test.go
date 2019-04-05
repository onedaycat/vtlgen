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
				Request:     "mapping-templates/resolver/mutation.createProduct/req.vtl",
				Response:    "mapping-templates/resolver/mutation.createProduct/res.vtl",
				Datasource:  "productLambda",
			},
			{
				GraphqlType: "Mutation",
				Field:       "languages",
				Request:     "mapping-templates/resolver/mutation.languages/req.vtl",
				Response:    "mapping-templates/resolver/mutation.languages/res.vtl",
				Datasource:  "haloLambda",
			},
			{
				GraphqlType: "Namespace",
				Field:       "languages",
				Request:     "mapping-templates/resolver/namespace.languages/before.vtl",
				Response:    "mapping-templates/resolver/namespace.languages/after.vtl",
				Kind:        "PIPELINE",
				Functions: []string{
					"oneFunction",
					"twoFunction",
				},
			},
			{
				GraphqlType: "Query",
				Field:       "product",
				Request:     "mapping-templates/resolver/query.product/req.vtl",
				Response:    "mapping-templates/resolver/query.product/res.vtl",
				Datasource:  "productLambda",
			},
			{
				GraphqlType: "Subscription",
				Field:       "product",
				Request:     "mapping-templates/resolver/subscription.product/req.vtl",
				Response:    "mapping-templates/resolver/subscription.product/res.vtl",
				Datasource:  "productLambda",
			},
		},
		Functions: []*Function{
			{
				Name:       "oneFunction",
				Datasource: "productLambda",
				Request:    "mapping-templates/function/oneFunction/req.vtl",
				Response:   "mapping-templates/function/oneFunction/res.vtl",
			},
			{
				Name:       "twoFunction",
				Datasource: "haloLambda",
				Request:    "mapping-templates/function/twoFunction/req.vtl",
				Response:   "mapping-templates/function/twoFunction/res.vtl",
			},
		},
	}

	mappingTemplates := GenerateMappingTemplates("testdata")
	require.Equal(t, expect, mappingTemplates)
}

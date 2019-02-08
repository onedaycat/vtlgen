package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFields(t *testing.T) {
	paths := []string{
		"mapping-templates/haloLambda/mutation/createProduct/req.vtl",
		"mapping-templates/haloLambda/mutation/createProduct/res.vtl",
		"mapping-templates/haloLambda/query/product/req.vtl",
		"mapping-templates/haloLambda/subscription/product/res.vtl",
	}

	expect := &Fields{
		Fields: []*Field{
			&Field{
				Type:       "Mutation",
				Field:      "createProduct",
				Request:    "haloLambda/mutation/createProduct/req.vtl",
				Response:   "haloLambda/mutation/createProduct/res.vtl",
				DataSource: "haloLambda",
			},
			&Field{
				Type:       "Query",
				Field:      "product",
				Request:    "haloLambda/query/product/req.vtl",
				Response:   "res.vtl",
				DataSource: "haloLambda",
			},
			&Field{
				Type:       "Subscription",
				Field:      "product",
				Request:    "req.vtl",
				Response:   "haloLambda/subscription/product/res.vtl",
				DataSource: "haloLambda",
			},
		},
	}

	fiels := getFields(paths)

	require.Equal(t, expect, fiels)
}

// +build integration

package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessGenerateDatasources(t *testing.T) {
	expect := &Datasources{
		Datasources: []*Datasource{
			{
				Type: "NONE",
				Name: "local",
			},
			{
				Type: "AWS_LAMBDA",
				Name: "accountQuery",
				Config: &Config{
					LambdaFunctionArn: "arn:aws:lambda:${env:AWS_REGION}:${self:provider.accountId}:function:sel-account-qry-${self:provider.stage}-qry:$LATEST",
					ServiceRoleArn:    "arn:aws:iam::805785015849:role/${self:service}-${self:provider.stage}",
				},
			},
			{
				Type: "AWS_LAMBDA",
				Name: "accountMutation",
				Config: &Config{
					LambdaFunctionArn: "arn:aws:lambda:${env:AWS_REGION}:${self:provider.accountId}:function:sel-account-cmd-${self:provider.stage}-cmd:$LATEST",
					ServiceRoleArn:    "arn:aws:iam::805785015849:role/${self:service}-${self:provider.stage}",
				},
			},
			{
				Type: "AWS_LAMBDA",
				Name: "storeMutation",
				Config: &Config{
					LambdaFunctionArn: "arn:aws:lambda:${env:AWS_REGION}:${self:provider.accountId}:function:sel-store-cmd-${self:provider.stage}-cmd:XXXX",
					ServiceRoleArn:    "XXXX",
				},
			},
		},
	}

	datasources := GenerateDatasources("testdata/datasource_generate.yml")
	require.Equal(t, expect, datasources)
}

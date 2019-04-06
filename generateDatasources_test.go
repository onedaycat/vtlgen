package vtlgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessGenerateDatasources(t *testing.T) {
	expect := &DatasourcesGenerated{
		Datasources: []*Datasource{
			{
				Type: "NONE",
				Name: "local",
			},
			{
				Type: "AWS_LAMBDA",
				Name: "accountQuery",
				Config: &Config{
					LambdaFunctionArn: "arn:aws:lambda:${self:provider.region}:FAKE12345678:function:sel-account-qry-${self:provider.stage}-qry:$LATEST",
					ServiceRoleArn:    "arn:aws:iam::FAKE12345678:role/${self:service}-${self:provider.stage}",
				},
			},
			{
				Type: "AWS_LAMBDA",
				Name: "accountMutation",
				Config: &Config{
					LambdaFunctionArn: "arn:aws:lambda:${self:provider.region}:FAKE12345678:function:sel-account-cmd-${self:provider.stage}-cmd:$LATEST",
					ServiceRoleArn:    "arn:aws:iam::FAKE12345678:role/${self:service}-${self:provider.stage}",
				},
			},
			{
				Type: "AWS_LAMBDA",
				Name: "storeMutation",
				Config: &Config{
					LambdaFunctionArn: "arn:aws:lambda:${self:provider.region}:FAKE12345678:function:sel-store-cmd-${self:provider.stage}-cmd:VVVV",
					ServiceRoleArn:    "arn:aws:iam::FAKE12345678:role/XXXX",
				},
			},
		},
	}

	datasources := GenerateDatasources("testdata/datasource.yml")
	require.Equal(t, expect, datasources)
}

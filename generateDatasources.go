package vtlgen

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func GenerateDatasources(parseDirectory string) *DatasourcesGenerated {
	var err error
	config, err := ioutil.ReadFile(parseDirectory)
	if err != nil {
		panic(err)
	}

	dg := &DatasourceConfig{}
	err = yaml.Unmarshal(config, dg)
	if err != nil {
		panic(err)
	}

	var dss []*Datasource

	// none datasource
	dss = append(dss, &Datasource{
		Type: NoneString,
		Name: dg.NoneDatasource,
	})

	// lambda datasource
	for _, ld := range dg.LambdaDatasources {
		var version, roleArn string
		ds := &Datasource{}
		ds.Config = &Config{}

		if ld.ServiceRoleArn == EmptyString {
			roleArn = strings.Replace(dg.ServiceRoleArn, "${accountId}", dg.AccountID, 1)
		} else {
			roleArn = ld.ServiceRoleArn
		}

		if ld.Version == EmptyString {
			version = LastestString
		} else {
			version = ld.Version
		}

		ds.Type = AwsLambdaString
		ds.Name = ld.Name
		ds.Config.ServiceRoleArn = roleArn
		ds.Config.LambdaFunctionArn = "arn:aws:lambda:${env:AWS_REGION}:${self:provider.accountId}:function:" + ld.Service + "-${self:provider.stage}-" + ld.Handler + ColonString + version

		dss = append(dss, ds)
	}

	return &DatasourcesGenerated{
		Datasources: dss,
	}
}

package vtlgen

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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
		var version string
		ds := &Datasource{}
		ds.Config = &Config{}

		if ld.Version == EmptyString {
			version = LastestString
		} else {
			version = ld.Version
		}

		ds.Type = AwsLambdaString
		ds.Name = ld.Name

		if ld.ServiceRole != EmptyString {
			ds.Config.ServiceRoleArn = fmt.Sprintf("arn:aws:iam::%s:role/%s", dg.AccountID, ld.ServiceRole)
		} else {
			ds.Config.ServiceRoleArn = fmt.Sprintf("arn:aws:iam::%s:role/%s", dg.AccountID, dg.ServiceRole)
		}
		ds.Config.LambdaFunctionArn = fmt.Sprintf("arn:aws:lambda:${self:provider.region}:%s:function:%s-${self:provider.stage}:%s", dg.AccountID, ld.Service, version)

		dss = append(dss, ds)
	}

	return &DatasourcesGenerated{
		Datasources: dss,
	}
}

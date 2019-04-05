package vtlgen

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// const (
// 	datasource        = 0
// 	graphqlType       = 1
// 	field             = 2
// 	requestOrResponse = 3
// )

// const (
// 	emptyString    = ""
// 	pathDelim      = "/"
// 	pathResolver   = "/resolver"
// 	pathFunctions  = "/function"
// 	reqFilename    = "req.vtl"
// 	resFilename    = "res.vtl"
// 	beforeFilename = "before.vtl"
// 	afterFilename  = "after.vtl"
// 	configFilename = "config.yml"
// 	pipeline       = "PIPELINE"
// )

func GenerateDatasources(parseDirectory string) *Datasources {
	var err error
	config, err := ioutil.ReadFile(parseDirectory)
	if err != nil {
		panic(err)
	}

	dg := &DatasourceGenerate{}
	err = yaml.Unmarshal(config, dg)
	if err != nil {
		panic(err)
	}

	var dss []*Datasource

	// none datasource
	dss = append(dss, &Datasource{
		Type: "NONE",
		Name: dg.NoneDatasource,
	})

	// lambda datasource
	for _, ld := range dg.LambdaDatasources {
		var version, roleArn string
		ds := &Datasource{}
		ds.Config = &Config{}

		if ld.ServiceRoleArn == "" {
			roleArn = strings.Replace(dg.ServiceRoleArn, "${accountId}", dg.AccountID, 1)
		} else {
			roleArn = ld.ServiceRoleArn
		}

		if ld.Version == "" {
			version = "$LATEST"
		} else {
			version = ld.Version
		}

		ds.Type = "AWS_LAMBDA"
		ds.Name = ld.Name
		ds.Config.ServiceRoleArn = roleArn
		ds.Config.LambdaFunctionArn = "arn:aws:lambda:${env:AWS_REGION}:${self:provider.accountId}:function:" + ld.Service + "-${self:provider.stage}-" + ld.Handler + ":" + version

		dss = append(dss, ds)
	}

	return &Datasources{
		Datasources: dss,
	}
}

package vtlgen

type Datasources struct {
	Datasources []*Datasource `yaml:"Datasources"`
}

type Datasource struct {
	Type   string  `yaml:"type"`
	Name   string  `yaml:"name"`
	Config *Config `yaml:"config"`
}

type Config struct {
	LambdaFunctionArn string `yaml:"lambdaFunctionArn"`
	ServiceRoleArn    string `yaml:"ServiceRoleArn"`
}

type DatasourceGenerate struct {
	AccountID         string              `yaml:"accountId"`
	ServiceRoleArn    string              `yaml:"serviceRoleArn"`
	NoneDatasource    string              `yaml:"noneDatasource"`
	LambdaDatasources []*LambdaDatasource `yaml:"lambdaDatasources"`
}

type LambdaDatasource struct {
	Name           string `yaml:"name"`
	Service        string `yaml:"service"`
	Version        string `yaml:"version"`
	Handler        string `yaml:"handler"`
	ServiceRoleArn string `yaml:"serviceRoleArn"`
}

package vtlgen

type DatasourcesGenerated struct {
	Datasources []*Datasource `yaml:"DatasourcesGenerated"`
}

type Datasource struct {
	Type   string  `yaml:"type"`
	Name   string  `yaml:"name"`
	Config *Config `yaml:"config,omitempty"`
}

type Config struct {
	LambdaFunctionArn string `yaml:"lambdaFunctionArn"`
	ServiceRoleArn    string `yaml:"ServiceRoleArn"`
}

type DatasourceConfig struct {
	AccountID         string                    `yaml:"accountId"`
	ServiceRoleArn    string                    `yaml:"serviceRoleArn"`
	NoneDatasource    string                    `yaml:"noneDatasource"`
	LambdaDatasources []*LambdaDatasourceConfig `yaml:"lambdaDatasources"`
}

type LambdaDatasourceConfig struct {
	Name           string `yaml:"name"`
	Service        string `yaml:"service"`
	Version        string `yaml:"version"`
	Handler        string `yaml:"handler"`
	ServiceRoleArn string `yaml:"serviceRoleArn"`
}

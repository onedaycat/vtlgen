package vtlgen

type DatasourcesGenerated struct {
	Datasources []*Datasource `yaml:"datasources"`
}

type Datasource struct {
	Type   string  `yaml:"type"`
	Name   string  `yaml:"name"`
	Config *Config `yaml:"config,omitempty"`
}

type Config struct {
	LambdaFunctionArn string `yaml:"lambdaFunctionArn"`
	ServiceRoleArn    string `yaml:"serviceRoleArn"`
}

type DatasourceConfig struct {
	AccountID         string                    `yaml:"accountId"`
	ServiceRole       string                    `yaml:"serviceRole"`
	NoneDatasource    string                    `yaml:"noneDatasource"`
	LambdaDatasources []*LambdaDatasourceConfig `yaml:"lambdaDatasources"`
}

type LambdaDatasourceConfig struct {
	Name        string `yaml:"name"`
	Service     string `yaml:"service"`
	Version     string `yaml:"version"`
	ServiceRole string `yaml:"serviceRole"`
}

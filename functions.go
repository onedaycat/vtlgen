package vtlgen

type Function struct {
	Name       string `yaml:"name"`
	Request    string `yaml:"request"`
	Response   string `yaml:"response"`
	Datasource string `yaml:"dataSource"`
}

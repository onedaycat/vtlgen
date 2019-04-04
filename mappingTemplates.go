package vtlgen

type MappingTemplates struct {
	Templates []*Template `yaml:"fields"`
	Functions []*Function `yaml:"functions"`
}

type Template struct {
	GraphqlType string   `yaml:"type"`
	Field       string   `yaml:"field"`
	Request     string   `yaml:"request"`
	Response    string   `yaml:"response"`
	DataSource  string   `yaml:"dataSource,omitempty"`
	Kind        string   `yaml:"kind,omitempty"`
	Functions   []string `yaml:"functions,omitempty"`
}

type Function struct {
	Name       string `yaml:"name"`
	Request    string `yaml:"request"`
	Response   string `yaml:"response"`
	DataSource string `yaml:"dataSource"`
}

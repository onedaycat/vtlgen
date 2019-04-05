package vtlgen

type MappingTemplates struct {
	Templates []*Template `yaml:"fields,omitempty"`
	Functions []*Function `yaml:"functions,omitempty"`
}

type Template struct {
	GraphqlType string   `yaml:"type"`
	Field       string   `yaml:"field"`
	Request     string   `yaml:"request"`
	Response    string   `yaml:"response"`
	Datasource  string   `yaml:"dataSource,omitempty"`
	Kind        string   `yaml:"kind,omitempty"`
	Functions   []string `yaml:"functions,omitempty"`
}

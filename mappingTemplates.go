package vtlgen

type MappingTemplates struct {
	MappingTemplates []*Template `yaml:"fields"`
}

type Template struct {
	GraphqlType string `yaml:"type"`
	Field       string `yaml:"field"`
	Request     string `yaml:"request"`
	Response    string `yaml:"response"`
	DataSource  string `yaml:"dataSource"`
}

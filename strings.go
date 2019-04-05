package vtlgen

// string
const (
	EmptyString     = ""
	DotString       = "."
	PathDelim       = "/"
	PipelineString  = "PIPELINE"
	NoneString      = "NONE"
	LastestString   = "$LATEST"
	AwsLambdaString = "AWS_LAMBDA"
	ColonString     = ":"
)

// file
const (
	ReqFilename                = "req.vtl"
	ResFilename                = "res.vtl"
	BeforeFilename             = "before.vtl"
	AfterFilename              = "after.vtl"
	ConfigFilename             = "config.yml"
	ResolverFilename           = "resolver.yml"
	DatasourceFilename         = "datasource.yml"
	DatasourceGenerateFilename = "datasource_generate.yml"
)

// path
const (
	PathResolver         = "/resolver"
	PathFunctions        = "/function"
	PathMappingTemplates = "/mapping-templates"
)

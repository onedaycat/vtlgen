package cmd

import (
	"io/ioutil"

	v "github.com/onedaycat/vtlgen"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var temPath string
var temOut string

func init() {
	TemplateCmd.Flags().StringVarP(&temPath, "path", "p", ".", "Specific path that has mapping-templates folder")
	TemplateCmd.Flags().StringVarP(&temOut, "out", "o", "mapping-templates.generated.yml", "Specific output file Default: mapping-templates.generated.yml")
}

var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate resolvers and functions",
	Long:  `Generate resolvers and functions`,
	Run: func(cmd *cobra.Command, args []string) {
		mt := v.GenerateMappingTemplates(temPath)

		mtYML, err := yaml.Marshal(mt)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(temOut, mtYML, 0644)
		if err != nil {
			panic(err)
		}
	},
}

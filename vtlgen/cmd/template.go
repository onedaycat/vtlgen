package cmd

import (
	"io/ioutil"

	v "github.com/onedaycat/vtlgen"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var path string

func init() {
	TemplateCmd.Flags().StringVarP(&path, "path", "p", "", "Specific path that has mapping-templates folder")
}

var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate resolver.yml",
	Long:  `Generate resolver.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		if path == v.EmptyString {
			path = v.DotString
		}

		mt := v.GenerateMappingTemplates(path)

		mtYML, err := yaml.Marshal(mt)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(path+v.PathDelim+v.ResolverFilename, mtYML, 0644)
		if err != nil {
			panic(err)
		}
	},
}

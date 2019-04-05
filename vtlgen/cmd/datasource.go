package cmd

import (
	"io/ioutil"
	"strings"

	v "github.com/onedaycat/vtlgen"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var file string

func init() {
	DatasourceCmd.Flags().StringVarP(&file, "file", "f", "", "Specific file datasource_generate.yml")
}

var DatasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "Generate datasource.yml",
	Long:  `Generate datasource.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		if file == v.EmptyString {
			file = v.DotString
		}

		d := v.GenerateDatasources(file)

		dYML, err := yaml.Marshal(d)
		if err != nil {
			panic(err)
		}

		fd := strings.Replace(file, "_generate", "", 1)
		err = ioutil.WriteFile(fd, dYML, 0644)
		if err != nil {
			panic(err)
		}

	},
}

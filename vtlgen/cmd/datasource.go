package cmd

import (
	"io/ioutil"
	"strings"

	v "github.com/onedaycat/vtlgen"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var dataFile string
var dataOut string

func init() {
	DatasourceCmd.Flags().StringVarP(&dataFile, "file", "f", "datasource.yml", "Specific file default: datasource.yml")
	DatasourceCmd.Flags().StringVarP(&dataOut, "out", "o", "", "Specific out file default: <file>.generated.yml")
}

var DatasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "Generate datasource",
	Long:  `Generate datasource`,
	Run: func(cmd *cobra.Command, args []string) {
		d := v.GenerateDatasources(dataFile)

		dYML, err := yaml.Marshal(d)
		if err != nil {
			panic(err)
		}

		outPath := dataOut
		if outPath == v.EmptyString {
			outPath = strings.Replace(dataFile, ".yml", ".generated.yml", 1)
		}

		err = ioutil.WriteFile(outPath, dYML, 0644)
		if err != nil {
			panic(err)
		}

	},
}

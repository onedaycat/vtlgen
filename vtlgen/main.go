package main

import (
	"github.com/onedaycat/vtlgen/vtlgen/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "vtlgen",
	}
	rootCmd.AddCommand(cmd.DatasourceCmd)
	rootCmd.AddCommand(cmd.TemplateCmd)
	rootCmd.Execute()
}

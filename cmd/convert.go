package cmd

import (
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts oapi3 to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Converting to markdown...")
		service.ConvertToMarkdown()
	},
}


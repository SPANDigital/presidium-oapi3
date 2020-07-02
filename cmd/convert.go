package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service"
	"github.com/spf13/cobra"
)

var file string
var outputDir string

func init() {
	rootCmd.AddCommand(convertCmd)
	rootCmd.PersistentFlags().StringVar(&file, "file", "", "OpenAPI 3 spec file to generated documentation")
	rootCmd.PersistentFlags().StringVar(&outputDir, "outputDir", ".", "The output directory")
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts OAPI3 to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Converting to markdown...")
		markdownService, err := service.NewMarkdownService()
		if err != nil {
			log.Fatal(err)
		}
		err = markdownService.ConvertToMarkdown(file, outputDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

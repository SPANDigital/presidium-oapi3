package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service"
	"github.com/spf13/cobra"
)

var file string
var outputDir string
var referenceURL string

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&file, "file", "f", "", "OpenAPI 3 spec file")
	_ = convertCmd.MarkFlagRequired("file")
	convertCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "The output directory")
	_ = convertCmd.MarkFlagRequired("outputDir")
	convertCmd.Flags().StringVarP(&referenceURL, "referenceURL", "r", "", "The reference URL")
	_ = convertCmd.MarkFlagRequired("referenceURL")
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts an OpenAPI 3 spec to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Converting to markdown...")
		markdownService, err := service.NewMarkdownService(referenceURL)
		if err != nil {
			log.Fatal(err)
		}
		err = markdownService.ConvertToMarkdown(file, outputDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

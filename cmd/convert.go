package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/markdown"
	"github.com/spf13/cobra"
)

var file string
var outputDir string
var referenceURL string
var apiName string
var titleFormat string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts an OpenAPI 3 spec to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Converting to markdown...")
		markdownService, err := markdown.NewMarkdownService(referenceURL, apiName)
		if err != nil {
			log.Fatal(err)
		}
		if titleFormat == "MethodURL" {
			err = markdownService.ConvertToMarkdown(file, outputDir, true)
		} else {
			err = markdownService.ConvertToMarkdown(file, outputDir, false)
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Flags
	convertCmd.Flags().StringVarP(&file, "file", "f", "", "OpenAPI 3 spec file")
	convertCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "The output directory")
	convertCmd.Flags().StringVarP(&referenceURL, "referenceURL", "r", "", "The reference URL")
	convertCmd.Flags().StringVarP(&apiName, "apiName", "n", "", "The name under which the generated docs will be grouped")

	// Required flags
	_ = convertCmd.MarkFlagRequired("file")
	_ = convertCmd.MarkFlagRequired("outputDir")
	_ = convertCmd.MarkFlagRequired("referenceURL")
}

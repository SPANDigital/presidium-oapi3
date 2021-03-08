package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service"
	"github.com/spf13/cobra"
)

var file string
var outputDir string
var referenceURL string
var apiName string
var titleFormat string

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&file, "file", "f", "", "OpenAPI 3 spec file")
	_ = convertCmd.MarkFlagRequired("file")
	convertCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "The output directory")
	_ = convertCmd.MarkFlagRequired("outputDir")
	convertCmd.Flags().StringVarP(&referenceURL, "referenceURL", "r", "", "The reference URL")
	_ = convertCmd.MarkFlagRequired("referenceURL")
	convertCmd.Flags().StringVarP(&apiName, "apiName", "n", "", "The name under which the generated docs will be grouped")
	convertCmd.Flags().StringVarP(&titleFormat, "titleFormat", "t", "", "The template format used to create the title for each operation. \nValid options are: \n\t- operationId: (Default) Uses the value of the operationId field.\n\t- MethodURL: Uses a combination of the Method property and the URL.")
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts an OpenAPI 3 spec to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Converting to markdown...")
		markdownService, err := service.NewMarkdownService(referenceURL, apiName)
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

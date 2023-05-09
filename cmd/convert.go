package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/markdown"
	"github.com/spf13/cobra"
)

var file string

var cfg markdown.Config

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts an OpenAPI 3 spec to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Converting to markdown...")
		markdownService, err := markdown.NewMarkdownService(cfg)
		if err != nil {
			log.Fatal(err)
		}

		err = markdownService.ConvertToMarkdown(file)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Flags
	convertCmd.Flags().StringVarP(&file, "file", "f", "", "OpenAPI 3 spec file")
	convertCmd.Flags().StringVarP(&cfg.OutputDir, "outputDir", "o", "", "The output directory")
	convertCmd.Flags().StringVarP(&cfg.ReferenceURL, "referenceURL", "r", "reference", "The reference URL")
	convertCmd.Flags().StringVarP(&cfg.ApiName, "apiName", "n", "", "The name under which the generated docs will be grouped")
	convertCmd.Flags().StringVarP(&cfg.TitleFormat, "titleFormat", "t", "", "The template format used to create the title for each operation. \nValid options are: \n\t- operationId: (Default) Uses the value of the operationId field.\n\t- MethodURL: Uses a combination of the Method property and the URL.")
	convertCmd.Flags().BoolVarP(&cfg.SortFilePath, "sortFilePath", "s", false, "Sort by filepath by prefixing a weight to the filename. Default is to use front matter weight")
	convertCmd.Flags().BoolVarP(&cfg.InlineProperties, "inlineProperties", "i", false, "Inline properties in the request and response schemas")
	convertCmd.Flags().BoolVarP(&cfg.AllowExternalRefs, "allowExternalRefs", "e", false, "Allow external references in the OpenAPI spec. ")

	// Required flags
	_ = convertCmd.MarkFlagRequired("file")
	_ = convertCmd.MarkFlagRequired("outputDir")
}

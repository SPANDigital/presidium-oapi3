package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/infrastructure/log"
	"github.com/SPANDigital/presidium-oapi3/pkg/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
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
		markdownService.ConvertToMarkdown("https://adss.qa.spandigital.net/api/v2/api-docs/new-admin", "admin")
		markdownService.ConvertToMarkdown("https://adss.qa.spandigital.net/api/v2/api-docs/auth", "auth")
		markdownService.ConvertToMarkdown("https://adss.qa.spandigital.net/api/v2/api-docs/graph", "graph")
		markdownService.ConvertToMarkdown("https://adss.qa.spandigital.net/api/v2/api-docs/ingestion", "ingestion")
		markdownService.ConvertToMarkdown("https://adss.qa.spandigital.net/api/v2/api-docs/logging", "logging")
		markdownService.ConvertToMarkdown("https://adss.qa.spandigital.net/api/v2/api-docs/developer", "developer")
	},
}

package cmd

import (
	"github.com/SPANDigital/presidium-oapi3/pkg/log"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "presidium-oapi3",
	Short: "A Presidium tool that converts OAPI3 spec's to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

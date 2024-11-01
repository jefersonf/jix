package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jix",
	Short: "Jira Issue eXtractor",
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

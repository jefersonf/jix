package cmd

import (
	"log"

	"github.com/jefersonf/jix/jira"
	"github.com/spf13/cobra"
)

var (
	projectKey string

	outputPath   string
	outputFormat string

	verbose bool
)

var jixCmd = &cobra.Command{
	Use:   "jix",
	Short: "Jira Issue eXtractor",
	Run: func(cmd *cobra.Command, args []string) {
		saveToFile(extractIssues())
	},
}

func init() {
	jixCmd.Flags().StringVarP(&projectKey, "project-key", "p", "", "Jira project key")
	jixCmd.Flags().StringVarP(&outputFormat, "format", "f", "jsonl", "Output format (only JSONL and CSV are available)")
	jixCmd.Flags().StringVarP(&outputPath, "output-path", "o", "./data", "Path to the output file")
	jixCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Set verbose mode")
}

func extractIssues() []jira.Issue {
	verboseLog("JIX started\n")
	issues, err := jira.FetchIssues(projectKey)
	if err != nil {
		log.Fatalln(err)
		return []jira.Issue{}
	}
	return issues
}

func verboseLog(format string, v ...any) {
	if verbose {
		log.Printf(format, v...)
	}
}

func Exec() {
	if err := jixCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

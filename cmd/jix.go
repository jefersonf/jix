package cmd

import (
	"fmt"
	"log"

	"github.com/jeferson/jix/jira"
	"github.com/spf13/cobra"
)

var (
	projectKey   string
	outputFormat string
	outputPath   string
	verbose      bool
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
	jixCmd.Flags().StringVarP(&outputPath, "output", "o", "./issues", "Path to the output file")
	jixCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Set verbose mode")
}

func extractIssues() []jira.Issue {
	verboseLog("start JIX\n")
	issues, err := jira.FetchIssues(projectKey)
	if err != nil {
		log.Fatalln(err)
		return []jira.Issue{}
	}
	return issues
}

func saveToFile(issues []jira.Issue) {
	for _, issue := range issues {
		fmt.Println(issue)
	}
	verboseLog("saving %s issues into %s folder\n", projectKey, outputPath)
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

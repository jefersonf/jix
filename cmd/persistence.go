package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jefersonf/jix/jira"
)

func saveToFile(issues []jira.Issue) {
	// TODO decide output format and then save to file
	verboseLog("saving %s issues into %s folder\n", projectKey, outputPath)
	if outputFormat == "jsonl" {
		saveToJSONLFile(issues)
	} else {
		saveToFileCSV(issues)
	}
}

func saveToFileCSV(issues []jira.Issue) {
	panic("unimplemented")
}

func saveToJSONLFile(issues []jira.Issue) {

	jsonFilePath := fmt.Sprintf("%s/%s.jsonl", outputFormat, projectKey)

	file, err := os.Create(jsonFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, issue := range issues {
		writeIssueToFile(file, &issue)
	}

	verboseLog("JSONL data has been written to %s\n", jsonFilePath)
}

func writeIssueToFile(file *os.File, issue *jira.Issue) {
	jsonData, err := json.Marshal(*issue)
	if err != nil {
		log.Println("Error marshalling issue to JSON:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}

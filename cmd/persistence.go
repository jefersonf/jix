package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jefersonf/jix/jira"
)

func saveToFile(issues []jira.Issue) {
	verboseLog("%v items extracted\n", len(issues))
	verboseLog("saving %s issues into %s folder\n", projectKey, outputPath)

	if outputFormat == "jsonl" {
		saveToJSONLFile(issues)
	} else {
		saveToCSVFile(issues)
	}
}

func saveToCSVFile(issues []jira.Issue) {
	panic("unimplemented")
}

func saveToJSONLFile(issues []jira.Issue) {

	jsonFilePath := strings.ToLower(fmt.Sprintf("%s/%s.jsonl", outputPath, projectKey))

	file, err := createOutputDir(jsonFilePath)
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

	jsonData = append(jsonData, byte('\n'))
	_, err = file.Write(jsonData)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}

func createOutputDir(outputPath string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0770); err != nil {
		return nil, err
	}
	return os.Create(outputPath)
}

package cmd

import (
	"encoding/csv"
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

	_, err := createOutputDir(outputPath)
	if err != nil {
		log.Println("Error creating output dir:", err)
		return
	}

	if outputFormat == "jsonl" {
		saveToJSONLFile(issues)
	} else if outputFormat == "csv" {
		saveToCSVFile(issues)
	} else {
		log.Fatalln("unsupported output format")
	}
}

func saveToCSVFile(issues []jira.Issue) {
	csvFilePath := strings.ToLower(fmt.Sprintf("%s/%s.csv", outputPath, projectKey))

	file, err := os.Create(csvFilePath)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Key", "Summary", "StatusDescription", "StatusName"})

	for _, issue := range issues {
		writer.Write([]string{
			issue.ID,
			issue.Key,
			issue.Fields.Summary,
			issue.Fields.Status.Description,
			issue.Fields.Status.Name,
		})
	}

	verboseLog("CSV data has been written to %s\n", csvFilePath)
}

func saveToJSONLFile(issues []jira.Issue) {

	jsonFilePath := strings.ToLower(fmt.Sprintf("%s/%s.jsonl", outputPath, projectKey))

	file, err := os.Create(jsonFilePath)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, issue := range issues {
		jsonData, err := json.Marshal(issue)
		if err != nil {
			log.Println("Error marshalling issue to JSON:", err)
		} else {
			jsonData = append(jsonData, byte('\n'))
			writeIssueToFile(file, jsonData)
		}
	}

	verboseLog("JSONL data has been written to %s\n", jsonFilePath)
}

func writeIssueToFile(file *os.File, data []byte) {
	_, err := file.Write(data)
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

func createOutputDir(outputPath string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0770); err != nil {
		return nil, err
	}
	return nil, nil
}

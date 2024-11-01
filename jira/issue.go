package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Status  struct {
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}

func FetchIssues(projectKey string) ([]Issue, error) {

	jira, err := getCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to load jira credentials: %w", err)
	}

	jqlQuery := fmt.Sprintf("project = %s", projectKey)
	url := fmt.Sprintf("%s/rest/api/2/search", jira.APIHost)

	params := map[string]interface{}{
		"jql":        jqlQuery,
		"fields":     []string{"summary", "status"},
		"maxResults": 1000,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal jira params: %w", err)
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(jira.UserEmail, jira.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch issues: %d", resp.StatusCode)
	}

	var result struct {
		Issues []Issue `json:"issues"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Issues, nil
}

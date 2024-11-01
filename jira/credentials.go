package jira

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type jiraCredentials struct {
	UserEmail string
	APIHost   string
	APIKey    string
}

func getCredentials() (*jiraCredentials, error) {

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading .env file")
	}

	var userEmail, apiHost, apiKey string

	if apiHost = os.Getenv("JIRA_API_HOST"); len(apiHost) == 0 {
		return nil, errors.New("api host is missing")
	}

	if userEmail = os.Getenv("JIRA_USER_EMAIL"); len(userEmail) == 0 {
		return nil, errors.New("user email is missing")
	}
	if apiKey = os.Getenv("JIRA_API_KEY"); len(apiKey) == 0 {
		return nil, errors.New("api key is missing")
	}

	return &jiraCredentials{
		UserEmail: userEmail,
		APIHost:   apiHost,
		APIKey:    apiKey,
	}, nil
}

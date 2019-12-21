package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// get a single issue
func getIssueURL(owner, repo, number string) string {
	return fmt.Sprintf("https://api.github.com/%s/%s/issues/%s", owner, repo, number)
}

// get a group issues
func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("http://api.github.com/repos/%s/%s/issues", owner, repo)
}

func setAuthorization(req *http.Request) error {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set")
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	return nil
}

// geta issue
func GetIssue(owner, repo, number string) (*Issue, error) {
	req, err := http.NewRequest("GET", getIssueURL(owner, repo, number), nil)
	if err != nil {
		//return nil, err
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//return nil, err
		panic(err)
		//return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		panic(err)
	}
	return &issue, nil
}

func CreateIssue(owner, repo string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(fields)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", getIssuesURL(owner, repo), buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Content-Type", "application/json")
	err = setAuthorization(req)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("post issue failed: %s", resp.Status)
	}
	panic(err)
}

func patchIssue(owner, repo, number string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", getIssueURL(owner, repo, number), buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Content-Type", "application/json")
	err = setAuthorization(req)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("patch issue failed: %s", resp.Status)
	}

	return nil
}

// UpdateIssue
func UpdateIssue(owner, repo, number string, fields map[string]string) error {
	return patchIssue(owner, repo, number, fields)
}

// ReopenIssue
func ReopenIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "open",
	}
	return patchIssue(owner, repo, number, fields)
}

// CloseIssue
func CloseIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "closed",
	}
	return patchIssue(owner, repo, number, fields)
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	Comments  int
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!-

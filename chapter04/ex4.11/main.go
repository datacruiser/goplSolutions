/*
Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
invoking their preferred text editor when substantial text input is required.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/jijiwhywhy/goplSolutions/chapter04/ex4.11/editor"
	"github.com/jijiwhywhy/goplSolutions/chapter04/ex4.11/github"
)

const usage = `usage:
  create OWNER REPO
  get    OWNER REPO ISSUE_NUMBER
  edit   OWNER REPO ISSUE_NUMBER
  close  OWNER REPO ISSUE_NUMBER
  reopen OWNER REPO ISSUE_NUMBER
`

var templ = template.Must(template.New("issue").Funcs(template.FuncMap{"formatTime": formatTime}).Parse(`
Number:   {{.Number}}
URL:      {{.HTMLURL}}
User:     {{.User.Login}}
Title:    {{.Title | printf "%.64s"}}
State:    {{.State}}
Comments: {{.Comments}}
Created:  {{.CreatedAt | formatTime}}
Updated:  {{.UpdatedAt | formatTime}}

{{if ne (len .Body) 0}}{{.Body}}{{else}}(no body){{end}}
`))

func formatTime(t time.Time) string {
	return t.Format("2019-12-21 15:07:00")
}

func create(owner, repo string) {
	fields := map[string]string{
		"title": "",
		"body":  "",
	}

	if err := editor.Edit(fields); err != nil {
		log.Fatal(err)
	}

	if err := github.CreateIssue(owner, repo, fields); err != nil {
		log.Fatal(err)
	}
}

func get(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	templ.Execute(os.Stdout, issue)
}

func edit(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]string{
		"title": issue.Title,
		"body":  issue.Body,
	}

	if err = editor.Edit(fields); err != nil {
		log.Fatal(err)
	}

	if err = github.UpdateIssue(owner, repo, number, fields); err != nil {
		log.Fatal(err)
	}
}

func close(owner, repo, number string) {

	if err := github.CloseIssue(owner, repo, number); err != nil {
		log.Fatal(err)
	}
}

func reopen(owner, repo, number string) {

	if err := github.ReopenIssue(owner, repo, number); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 4 {
		command, owner, repo := os.Args[1], os.Args[2], os.Args[3]
		switch command {
		case "create":
			create(owner, repo)
		default:
			fmt.Fprintf(os.Stderr, usage)
			os.Exit(1)
		}
	} else if len(os.Args) == 5 {
		command, owner, repo, number := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
		switch command {
		case "get":
			get(owner, repo, number)
		case "edit":
			edit(owner, repo, number)
		case "close":
			close(owner, repo, number)
		case "reopen":
			reopen(owner, repo, number)
		default:
			fmt.Fprintf(os.Stderr, usage)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
}

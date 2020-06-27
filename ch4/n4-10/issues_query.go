package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// IssuesURL url
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult result
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	sorts := make([][]*Issue, 3)

	for _, item := range result.Items {
		days := int(time.Since(item.CreatedAt).Hours() / 24)
		if days < 31 {
			sorts[0] = append(sorts[0], item)
		} else if days <= 365 {
			sorts[1] = append(sorts[1], item)
		} else if days > 365 {
			sorts[2] = append(sorts[2], item)
		}
	}
	fmt.Printf("less than one month:\n")
	for _, item := range sorts[0] {

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("less than one year:\n")
	for _, item := range sorts[1] {

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("more than one year:\n")
	for _, item := range sorts[2] {

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

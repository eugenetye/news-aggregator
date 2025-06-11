package rss

import (
	"net/http"
	"strings"
	"time"
)

import "github.com/PuerkitoBio/goquery"

func ScrapeFullContent(url string) (string, error) {
	// Create HTTP client with timeout and proper headers
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Return empty string (not error) for non-200 responses so fallback works
	if resp.StatusCode != 200 {
		return "", nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// TechCrunch uses WordPress blocks - target wp-block-paragraph specifically
	var contentParts []string
	
	// First try to get all wp-block-paragraph elements
	doc.Find(".wp-block-paragraph").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			contentParts = append(contentParts, text)
		}
	})
	
	// Join all content parts
	content := strings.Join(contentParts, "\n\n")

	// Clean up the final content
	content = strings.TrimSpace(content)

	return content, nil
}
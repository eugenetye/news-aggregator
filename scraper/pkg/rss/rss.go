package rss

import (
	"html"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"scraper/pkg/model"
)

func Fetch(url string, source string) ([]model.Article, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	var articles []model.Article
	for _, item := range feed.Items {
		link := item.Link
		fullContent, err := ScrapeFullContent(link)
		
		if err != nil || fullContent == "" {
			fullContent = item.Content
			if fullContent == "" {
				fullContent = item.Description
			}
		} 

		fullContent = html.UnescapeString(strings.TrimSpace(fullContent))
		if fullContent == "" {
			continue // Skip articles without content
		}

		timestamp := time.Now()
		if item.PublishedParsed != nil {
			timestamp = *item.PublishedParsed
		}

		articles = append(articles, model.Article{
			ID:        item.GUID,
			Title:     item.Title,
			URL:       item.Link,
			Source:    source,
			Timestamp: timestamp,
			Content:   fullContent,
		})

		if len(articles) == 2 {
			break // Stop once we have 2 valid articles
		}
	}
	return articles, nil
}

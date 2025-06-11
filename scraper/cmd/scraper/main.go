package main

import (
	"fmt"
	"log"

	"scraper/pkg/reddit"
	"scraper/pkg/rss"
	"scraper/pkg/model"
)

func main() {
	var allArticles []string

	// Fetch from Reddit
	redditSubs := []string{"stocks", "golf", "lifeprotips", "selfimprovement"}
	for _, sub := range redditSubs {
		articles, err := reddit.Fetch(sub)
		if err != nil {
			log.Println("Reddit fetch error:", err)
			continue
		}

		for _, a := range articles {
			allArticles = append(allArticles, formatArticle(a))
		}
	}

	// Fetch from RSS feeds
	rssFeeds := map[string]string{
		"TechCrunch": "https://techcrunch.com/feed/",
	}
	for name, url := range rssFeeds {
		articles, err := rss.Fetch(url, name)
		if err != nil {
			log.Println("RSS fetch error for", name, ":", err)
			continue
		}
		for _, a := range articles {
			allArticles = append(allArticles, formatArticle(a))
		}
	}

	for  _, line := range allArticles {
		fmt.Println(line)
		fmt.Println("--------------------------------------------------")
	}
}

func formatArticle(a model.Article) string {
	return fmt.Sprintf("%s [%s]\n%s\n\n%s", a.Title, a.Source, a.URL, a.Content)
}

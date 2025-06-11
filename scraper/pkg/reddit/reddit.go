package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"scraper/pkg/model"
	"strings"
	"errors"
	"os"
	"log"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getAccessToken() (string, error) {
	clientID := os.Getenv("REDDIT_CLIENT_ID")
	clientSecret := os.Getenv("REDDIT_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return "", errors.New("missing Reddit API credentials")
	}
	
	req, _ := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader("grant_type=client_credentials"))
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("User-Agent", "news-aggregator/0.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err!= nil {
		return "", err
	}

	return result.AccessToken, nil
}

type redditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				ID        string    `json:"id"`
				Title     string    `json:"title"`
				URL       string    `json:"url"`
				Created   float64   `json:"created_utc"`
				Permalink string    `json:"permalink"`
				SelfText  string    `json:"selftext"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func Fetch(subreddit string) ([]model.Article, error) {
	token, err := getAccessToken()
	if err!= nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	url := fmt.Sprintf("https://oauth.reddit.com/r/%s/top.json?limit=30&t=day", subreddit)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "news-aggregator/0.1")

	res, err := httpClient.Do(req)
	if err!= nil {
		return nil, fmt.Errorf("failed to fetch data from Reddit: %w", err)
	}
	defer res.Body.Close()

	// Check if we got HTML instead of JSON (rate limited)
    if strings.Contains(res.Header.Get("Content-Type"), "text/html") {
        return nil, fmt.Errorf("reddit returned HTML - likely rate limited or blocked")
    }

	var rr redditResponse
	err = json.NewDecoder(res.Body).Decode(&rr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var articles []model.Article
	for _, child := range rr.Data.Children {
		post := child.Data
		if post.SelfText == "" {
			continue // Skip posts without content
		}
		articles = append(articles, model.Article{
			ID:		post.ID,
			Title:    post.Title,
			URL:      post.URL,
			Source:   fmt.Sprintf("reddit.com/r/%s", subreddit),
			Timestamp: time.Unix(int64(post.Created), 0),
			Content: post.SelfText,
		})

		if len (articles) >= 2 {
			break // Limit to 2 articles per subreddit
		}
	}
	log.Printf(" Reddit: fetched %d articles from r/%s", len(articles), subreddit)

	return articles, nil
}
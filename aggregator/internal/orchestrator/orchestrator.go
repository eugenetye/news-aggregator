package orchestrator

import (
	"aggregator/internal/llm"
	"aggregator/internal/storage"
	"log"
	"scraper/pkg/reddit"
	"scraper/pkg/rss"
	"sync"
	"github.com/redis/go-redis/v9"
	"time"
	"os"
)

func RunAggregation(outputPath string) error {
	var articles []storage.AggregatedArticle
	var mu sync.Mutex
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 2) // Limit concurrency to 2

	handleArticle := func(a storage.AggregatedArticle) {
		defer wg.Done()
		defer func() { <-semaphore}()

		summary, err := llm.Summarize(a.Content)
		if err != nil {
			log.Println("LLM summary error for article", a.Title, ":", err)
		}
		a.Summary = summary

		mu.Lock()
		articles = append(articles, a)
		mu.Unlock()
	}

	// Fetch from Reddit
	redditSubs := []string{"stocks", "lifeprotips", "selfimprovement"}
	for _, sub := range redditSubs {
		redditArticles, err := reddit.Fetch(sub)
		if err != nil {
			log.Println("Reddit fetch error:", err)
			continue
		}

		for _, a := range redditArticles {
			wg.Add(1)
			semaphore <- struct{}{} // Acquire semaphore
			go handleArticle(storage.FromModel(a))
		}
	}

	// Fetch from RSS feeds
	rssFeeds := map[string]string{
		"TechCrunch": "https://techcrunch.com/feed/",
	}
	for name, url := range rssFeeds {
		result, err := rss.Fetch(url, name)
		if err != nil {
			log.Println("RSS fetch error for", name, ":", err)
			continue
		}
		for _, a := range result {
			wg.Add(1)
			semaphore <- struct{}{}
			go handleArticle(storage.FromModel(a))
		}
	}

	wg.Wait() // Wait for all goroutines to finish
	log.Printf("Fetched and summarized %d articles from Reddit and RSS feeds", len(articles))

	// Fix Redis connection
    redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr == "" {
        redisAddr = "localhost:6379" // fallback for local dev
    }

    rdb := redis.NewClient(&redis.Options{
        Addr: redisAddr,
        DB:   0,
    })


	today := time.Now().Format("2006-01-02")
	err := storage.SaveToRedis(rdb, today, articles)
	if err != nil {
		log.Fatalf("Failed to save articles to Redis: %v", err)
	}
	return nil
}
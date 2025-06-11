package storage

import (
	"encoding/json"
	"os"
	"time"

	"scraper/pkg/model"
)

type AggregatedArticle struct {
	ID 	  		string    `json:"id"`
	Title 		string	`json:"title"`
	URL   		string    `json:"url"`
	Source 		string    `json:"source"`
	Timestamp 	time.Time `json:"timestamp"`
	Content 	string    `json:"content"`
	Summary		string	`json:"summary"`
}

func FromModel(m model.Article) AggregatedArticle {
	return AggregatedArticle{
		ID:        m.ID,
		Title:     m.Title,
		URL:       m.URL,
		Source:    m.Source,
		Timestamp: m.Timestamp,
		Content:   m.Content,
	}
}

func SaveToJSON(path string, articles []AggregatedArticle) error {
	data, err := json.MarshalIndent(articles, "", "  ")
	if err!= nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
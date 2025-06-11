package model

import "time"

type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	URL		 	string    `json:"url"`
	Source 		string    `json:"source"`
	Timestamp   time.Time `json:"timestamp"`
	Content   	string    `json:"content,omitempty"`
}
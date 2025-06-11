package handler

import (
	"api-gateway/internal/redis"
	"encoding/json"
	"net/http"
	"time"
	"math/rand"

	"github.com/gin-gonic/gin"
)

type AlexaItem struct {
	Uid	  			string `json:"uid"`
	UpdateDate		string `json:"updateDate"`
	TitleText		string `json:"titleText"`
	MainText		string `json:"mainText"`
	RedirectionUrl 	string `json:"redirectionUrl"`
}

func GetAlexaBriefing(c *gin.Context) {
	today := time.Now().Format("2006-01-02")
	key := "feed:" + today

	raw, err := redis.Rdb.Get(redis.Ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feed not found for today"})
		return
	}

	var articles []map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &articles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse feed data"})
		return
	}

	var items []AlexaItem
	for _, a := range articles {

		items = append(items, AlexaItem{
			Uid:			a["id"].(string),
			UpdateDate:		time.Now().UTC().Format("2006-01-02T07:00:00.000Z"),	
			TitleText:		a["title"].(string),
			MainText: 		a["summary"].(string),
			RedirectionUrl: a["url"].(string),
		})
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int){
		items[i], items[j] = items[j], items[i]
	})

	c.JSON(http.StatusOK, items)
}
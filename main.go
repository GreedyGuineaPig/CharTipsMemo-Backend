package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tweet struct {
	ID   string `json:"id"`
	Url  string `json:"url"`
	Char string `json:"char"`
}

var tweets = []tweet{
	{ID: "1", Url: "url1", Char: "SO"},
	{ID: "2", Url: "url2", Char: "KY"},
	{ID: "3", Url: "url3", Char: "MA"},
}

func getAll(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tweets)
}

func addTweet(context *gin.Context) {
	var newTweet tweet

	if err := context.BindJSON(&newTweet); err != nil {
		return
	}

	newTweet.ID = strconv.Itoa(len(tweets) + 1)

	tweets = append(tweets, newTweet)

	context.IndentedJSON(http.StatusCreated, newTweet)
}

func getChartweet(context *gin.Context) {
	char := context.Param("char")
	charTweets, err := getTweetByChar(char)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "character not found"})
	}
	context.IndentedJSON(http.StatusOK, charTweets)
}

func patchTweetById(context *gin.Context) {
	id := context.Param("id")
	url := context.Param("url")
	thisTweet, err := getTweetById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	thisTweet.Url = url
	context.IndentedJSON(http.StatusOK, thisTweet)
}

func deleteTweeet(context *gin.Context) {
	id := context.Param("id")
	index, err := getIndexBy(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	tweets = append(tweets[:index], tweets[index+1:]...)
	context.IndentedJSON(http.StatusOK, tweets)
}

func getIndexBy(id string) (int, error) {
	for i, t := range tweets {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("tweet with the id not found")
}

func getTweetById(id string) (*tweet, error) {
	for i, t := range tweets {
		if t.ID == id {
			return &tweets[i], nil
		}
	}
	return nil, errors.New("tweet with the id not found")
}

func getTweetByChar(char string) ([]tweet, error) {
	var charTweets []tweet
	for _, t := range tweets {
		if t.Char == char {
			charTweets = append(charTweets, t)
		}
	}
	if len(charTweets) == 0 {
		return charTweets, errors.New("char not found")
	}
	return charTweets, nil
}

func main() {
	router := gin.Default()
	router.GET("/", getAll)
	router.POST("/", addTweet)
	router.PATCH("/:id/:url", patchTweetById)
	router.DELETE("/:id", deleteTweeet)
	router.GET("/:char", getChartweet)
	router.Run("localhost:8080")
}

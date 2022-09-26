package controller

import (
	"net/http"

	"gomysql-api/model"

	"github.com/gin-gonic/gin"
)

func GetAll(context *gin.Context) {
	posts := model.FindAll()
	context.IndentedJSON(http.StatusOK, posts)
}

func AddPost(context *gin.Context) {
	var newPost model.Post

	if err := context.BindJSON(&newPost); err != nil {
		return
	}
	// don't accept arbitrary Ids
	newPost.Create()

	context.IndentedJSON(http.StatusCreated, newPost)
}

func GetCharpost(context *gin.Context) {
	char := context.Param("char")
	posts := model.FindAllByChar(char)
	context.IndentedJSON(http.StatusOK, posts)
}

func PatchPost(context *gin.Context) {
	var fixedPost model.Post

	if err := context.BindJSON(&fixedPost); err != nil {
		return
	}
	var oldPost = model.Post{ID: fixedPost.ID}
	oldPost.Find()

	oldPost.Version = fixedPost.Version
	oldPost.Char = fixedPost.Char
	oldPost.IsAntiChar = fixedPost.IsAntiChar
	oldPost.Body = fixedPost.Body

	oldPost.Updates()
	context.IndentedJSON(http.StatusOK, oldPost)
}

func DeletePost(context *gin.Context) {
	id := context.Param("id")
	var post model.Post
	post.FirstById(id)
	post.Delete()
	context.IndentedJSON(http.StatusOK, post)
}

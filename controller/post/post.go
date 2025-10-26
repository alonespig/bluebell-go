package post

import (
	"bluebell/froms"
	"bluebell/model"
	"bluebell/pkg/code"
	"bluebell/pkg/response"
	"bluebell/pkg/snowflake"
	"bluebell/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var form froms.CreatePostForm
	if err := c.ShouldBindJSON(&form); err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, nil)
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)

	post := &model.Post{
		PostID:      snowflake.GenID(),
		AuthorID:    userIDInt,
		CommunityID: form.CommunityID,
		Title:       form.Title,
		Content:     form.Content,
	}
	if err := service.CreatePost(post); err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func GetPostDetail(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
		return
	}
	postDetail, err := service.GetPostDetail(postIDInt)
	if err != nil {
		response.JSON(c, http.StatusNotFound, code.NotFound, err.Error())
		return
	}
	response.Success(c, postDetail)
}

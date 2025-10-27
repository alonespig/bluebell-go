package post

import (
	"bluebell/froms"
	"bluebell/global"
	"bluebell/model"
	"bluebell/pkg/code"
	"bluebell/pkg/response"
	"bluebell/pkg/snowflake"
	"bluebell/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary 创建帖子
// @Description 创建帖子
// @Accept json
// @Produce json
// @Param form body froms.CreatePostForm true "创建帖子表单"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post [post]
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

// @Summary 获取帖子详情
// @Description 获取帖子详情
// @Accept json
// @Produce json
// @Param id path int true "帖子ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post/{id} [get]
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

// @Summary 获取帖子列表
// @Description 获取帖子列表
// @Accept json
// @Produce json
// @Param page query int true "页码"
// @Param pageSize query int true "每页条数"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post [get]
func GetPostList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
		return
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
		return
	}
	postList, err := service.GetPostListByPage(pageInt, pageSizeInt)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	response.Success(c, postList)
}

// @Summary 投票
// @Description 投票 for post
// @Accept json
// @Produce json
// @Param form body froms.VotePostForm true "投票表单"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/vote [post]
func VoteForPost(c *gin.Context) {
	var form froms.VotePostForm
	if err := c.ShouldBindJSON(&form); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			// 翻译成中文
			messages := make([]string, 0)
			for _, e := range errs {
				messages = append(messages, e.Translate(global.Trans))
			}
			response.JSON(c, http.StatusBadRequest, code.InvalidParams, messages)
			return
		}
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, nil)
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)
	if err := service.VoteForPost(userIDInt, form); err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 获取用户投票
// @Description 获取用户投票
// @Accept json
// @Produce json
// @Param id path int true "帖子ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post/{id}/vote [get]
func GetUserVote(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)
	userVote, err := service.GetUserVote(userIDInt, postIDInt)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	rep := gin.H{
		"code": 1000,
		"msg":  "ok",
		"data": gin.H{
			"user_vote": userVote,
		},
	}
	c.JSON(http.StatusOK, rep)
}

// @Summary 获取用户帖子列表
// @Description 获取用户帖子列表
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post/user/{id} [get]
func GetPostListByUserID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)
	postList, err := service.GetPostListByUserID(userIDInt)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	rsp := gin.H{
		"list":  postList,
		"total": len(postList),
	}
	response.Success(c, rsp)
}

// @Summary 更新帖子
// @Description 更新帖子
// @Accept json
// @Produce json
// @Param id path int true "帖子ID"
// @Param form body froms.UpdatePostForm true "更新帖子表单"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post/{id} [put]
func UpdatePost(c *gin.Context) {
	var form froms.UpdatePostForm
	if err := c.ShouldBindJSON(&form); err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, nil)
		return
	}
	postID := c.Param("id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)
	if err := service.UpdatePost(userIDInt, postIDInt, form.Title, form.Content); err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 删除帖子
// @Description 删除帖子
// @Accept json
// @Produce json
// @Param id path int true "帖子ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post/{id} [delete]
// @Security ApiKeyAuth
func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)
	if err := service.DeletePost(userIDInt, postIDInt); err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 获取帖子列表状态
// @Description 获取帖子列表状态
// @Accept json
// @Produce json
// @Param form body froms.GetPostListStatusForm true "帖子列表状态表单"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/post/status [get]
func GetPostListByStatus(c *gin.Context) {
	var form froms.GetPostListStatusForm
	if err := c.ShouldBindJSON(&form); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			// 翻译成中文
			messages := make([]string, 0)
			for _, e := range errs {
				messages = append(messages, e.Translate(global.Trans))
			}
			response.JSON(c, http.StatusBadRequest, code.InvalidParams, messages)
			return
		}
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, nil)
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
		return
	}
	userIDInt := userID.(int)
	postList, err := service.GetPostListByStatus(form.PostIDs, userIDInt)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
		return
	}
	response.Success(c, postList)
}

package router

import (
	"bluebell/controller/community"
	"bluebell/controller/post"
	"bluebell/controller/user"
	"bluebell/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.Title = "Demo API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	router.Use(middleware.CORS())
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		v1.POST("signup", user.Register)
		v1.POST("login", user.Login)
		v1.GET("community", community.GetCommunityList)
		v1.GET("community/:id", community.GetCommunityDetail)

		v1.POST("post", middleware.JWTAuth(), post.CreatePost)
		v1.GET("post/:id", post.GetPostDetail)
		v1.GET("posts2", post.GetPostList)

		v1.GET("/users/me/posts", middleware.JWTAuth(), post.GetPostListByUserID)
		v1.POST("vote", middleware.JWTAuth(), post.VoteForPost)
		v1.GET("post/:id/vote", middleware.JWTAuth(), post.GetUserVote)
		v1.PUT("/post/:id", middleware.JWTAuth(), post.UpdatePost)
		v1.DELETE("/post/:id", middleware.JWTAuth(), post.DeletePost)
	}

	return router
}

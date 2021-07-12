package router

import (
	app "github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/infrastructure/persistence"
	"github.com/yino/nlp-controller/interfaces/corp"
	"github.com/yino/nlp-controller/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCoreRouter(c *gin.Engine, repo *persistence.Repositories) {
	userApp := app.NewUserApp(repo.User)
	UserInterFace := corp.NewUsersInterface(userApp)
	v1 := c.Group("v1")
	{
		core := v1.Group("core")
		{
			// 登录注册
			core.POST("/login", UserInterFace.HandlerUserLogin)
			core.POST("/register", UserInterFace.HandlerUserRegister)

			core.Use(middleware.CorpAuthTokenMiddleware(userApp))
			{
				// user
				core.GET("/user/info", UserInterFace.HandlerUserInfo)
				core.GET("/user/edit", UserInterFace.HandlerUserEdit)

				// question
				//core.GET("/question/index", interfaces.HandlerQuestionPage)
				//core.POST("/question/add", interfaces.HandlerQuestionAdd)
				//core.POST("/question/edit", interfaces.HandlerQuestionEdit)
				//core.GET("/question/delete", interfaces.HandlerQuestionDelete)
				//core.GET("/question/train", interfaces.HandlerQuestionTrain)
			}
		}
	}

}

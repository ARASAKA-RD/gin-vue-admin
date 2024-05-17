package example

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WordRouter struct{}

func (e *WordRouter) InitWordRouter(Router *gin.RouterGroup) {
	wordRouter := Router.Group("word").Use(middleware.OperationRecord())
	wordRouterWithoutRecord := Router.Group("word")
	exaWordApi := v1.ApiGroupApp.ExampleApiGroup.WordApi
	{
		wordRouter.POST("word", exaWordApi.createWord)
	}
	{
		wordRouterWithoutRecord.GET("word", exaWordApi.getOneWord)
	}
}
package routes

import (
	"ginweb/controller/index"
	"github.com/gin-gonic/gin"
)

func InitRouter(route *gin.Engine) {
	//业务接口
	index.InitRouter(route)

}

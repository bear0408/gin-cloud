package index

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//index 模块下 无需中间件过滤的
	indexNoAction := r.Group("/index/")
	{
		test := TestController{}
		indexNoAction.GET("test/test", test.Test) //.POST()
	}

}

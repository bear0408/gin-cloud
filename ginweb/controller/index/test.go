package index

import (
	"ginweb/common"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

type TestController struct {
}

func (this *TestController) Test(c *gin.Context) {

	str := "84?page=pages/squareVideoDetailNew"
	//str := "9¶ms=eyJ1cmwiOiJwYWdlcy93b3Jrcy9zdG9yeVdpbGxEZXRhaWxzL2NvbWljIiwicGFyYW1zIjp7InN0b3J5X2lkIjoyLCJjaGFwdGVySWQiOjExfX0="

	reg1 := regexp.MustCompile(`^\d*(\.)?\d*`)
	match := reg1.FindString(str)
	id, _ := strconv.Atoi(match)

	common.Success(c, id)
	return

}

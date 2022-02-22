package v1

import (
	"ginIMApi/constants/e"
	"ginIMApi/packages/utils"
	"github.com/gin-gonic/gin"
)

func UserEmoticon(c *gin.Context) {
	appG := utils.Gin{C: c}
	//返回数据
	appG.Response(e.SUCCESS, nil)
}

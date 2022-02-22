package utils

import (
	"ginIMApi/packages/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage get page offset parameters
func GetPageOffset(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

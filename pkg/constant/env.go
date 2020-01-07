package constant

import "github.com/gin-gonic/gin"

// 是否为生产环境
var ProductionMode = gin.Mode() == gin.ReleaseMode

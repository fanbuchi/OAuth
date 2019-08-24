package router

import (
	"github.com/gin-gonic/gin"
	"oauth_api/handler"
)

func Init(r *gin.Engine) {
	oauth := new(handler.Oauth)
	r.LoadHTMLGlob("html/*")
	v1 := r.Group("/v1")
	{
		v1.GET("/authorize", oauth.Authorize)
	}
}

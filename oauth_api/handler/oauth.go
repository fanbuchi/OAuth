package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Oauth struct {
}

func (o *Oauth) Authorize(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

package statusReady

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ready(c * gin.Context){
	c.String(http.StatusOK, "Ready")
}

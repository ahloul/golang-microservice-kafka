package middleware

import (
	"github.com/ahloul/loyalty-reports/infrastructure/helpers"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtGuard(c *gin.Context) {
	_, err := helpers.VerifyToken(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": []map[string]interface{}{
				{
					"status": http.StatusUnauthorized,
					"title":  "unauthenticated",
					"detail": gotrans.T("app.unauthenticated"),
				},
			},
		})
		return
	} else {
		c.Next()
	}
}

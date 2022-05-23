package http

import (
	"github.com/ahloul/loyalty-reports/purchase/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"net/http"
)

type PurchaseHandler struct {
	repo repository.PurchaseRepository
}

func (p PurchaseHandler) Index(e *gin.Engine) {
	data := p.repo.Find()
	e.GET("/purchases",
		func(c *gin.Context) {
			c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
			if err := jsonapi.MarshalPayload(c.Writer, data); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"errors": []map[string]interface{}{
						{
							"status": http.StatusInternalServerError,
							"title":  "incorrect_response_body",
							"detail": err.Error(),
						},
					},
				})
			}
		})

	return
}

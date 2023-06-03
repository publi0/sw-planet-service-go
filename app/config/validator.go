package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"net/http"
)

func HeaderValidator(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.GetHeader("correlation-id")

		if correlationId == "" {
			log.Ctx(ctx).Info().Msg("request missing correlation-id, returning bad request")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "missing [correlation-id] header"})
		}

		c.Next()
	}
}

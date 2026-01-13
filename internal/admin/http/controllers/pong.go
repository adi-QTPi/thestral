package controllers

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Pong(c *gin.Context) {
	s.response.Success(c, "PONG", nil)
}

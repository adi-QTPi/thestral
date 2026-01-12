package controllers

import (
	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/gin-gonic/gin"
)

func (s *Service) CreateProxy(c *gin.Context) {
	var req dto.RouteInput

	if err := c.ShouldBind(&req); err != nil {
		s.response.BadRequest(c, "Bad Input", err)
		return
	}

	if err := s.store.Create(req); err != nil {
		s.response.ServerError(c, err)
		return
	}

	s.response.Success(c, "New Proxy Created", nil)
}

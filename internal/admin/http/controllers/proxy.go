package controllers

import (
	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/gin-gonic/gin"
)

func (s *Service) CreateProxy(c *gin.Context) {
	var req dto.CreateRouteInput

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

func (s *Service) DeleteProxy(c *gin.Context) {
	var req dto.DeleteRouteInput

	if err := c.ShouldBindJSON(&req); err != nil {
		s.response.BadRequest(c, "Bad Input", err)
		return
	}

	if err := s.store.Delete(req); err != nil {
		s.response.ServerError(c, err)
		return
	}

	s.response.Success(c, "Proxy Deleted", nil)
}

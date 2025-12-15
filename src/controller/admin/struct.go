package controller

import "github.com/adi-QTPi/thestral/src/model"

type AdminController struct {
	Engine *model.Engine
}

func NewAdminController(e *model.Engine) *AdminController {
	return &AdminController{
		Engine: e,
	}
}

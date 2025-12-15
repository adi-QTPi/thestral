package controller

import "github.com/adi-QTPi/thestral/src/model"

type ProxyController struct {
	Engine *model.Engine
}

func NewProxyController(e *model.Engine) *ProxyController {
	return &ProxyController{
		Engine: e,
	}
}

package main

import (
	"github.com/adi-QTPi/thestral/src/api/admin"
	"github.com/adi-QTPi/thestral/src/api/proxy"
	"github.com/adi-QTPi/thestral/src/model"
)

func main() {
	e := model.NewEngine()
	adminRouter := admin.Router(e)
	proxyRouter := proxy.Router(e)

	go admin.Serve(adminRouter)
	proxy.Serve(proxyRouter)
}

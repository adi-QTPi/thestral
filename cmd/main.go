package main

import (
	"github.com/adi-QTPi/thestral/src/api/admin"
	"github.com/adi-QTPi/thestral/src/api/proxy"
)

func main() {
	proxyRouter := proxy.Router()
	adminRouter := admin.Router(proxyRouter)

	go admin.Serve(adminRouter)
	proxy.Serve(proxyRouter)
}

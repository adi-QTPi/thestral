package main

import (
	"log"

	"github.com/adi-QTPi/thestral/src/api/admin"
	"github.com/adi-QTPi/thestral/src/api/proxy"
	"github.com/adi-QTPi/thestral/src/config"
	"github.com/adi-QTPi/thestral/src/model"
)

func main() {
	cfg := config.LoadConfig()

	e := model.NewEngine(cfg)

	if err := e.LoadRedis(); err != nil {
		log.Fatalf("redis loading error : %v", err)
	}

	adminRouter := admin.Router(e)
	proxyRouter := proxy.Router(e)

	go admin.Serve(adminRouter, cfg)
	proxy.Serve(proxyRouter, cfg)
}

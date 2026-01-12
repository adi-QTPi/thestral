package main

import (
	"fmt"

	admin "github.com/adi-QTPi/thestral/internal/admin/http"
	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/listener"
	"github.com/adi-QTPi/thestral/internal/proxy"
	public "github.com/adi-QTPi/thestral/internal/proxy/http"
	"github.com/adi-QTPi/thestral/internal/store"
)

func main() {
	initServices()
}

func initServices() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error initialising config : %v", err)
		return
	}

	p := proxy.NewService()

	s, err := store.NewService(cfg)
	if err != nil {
		fmt.Printf("Error initialising db store 2: %v", err)
		return
	}

	l := listener.NewService(cfg, p, s)

	l.Load()
	l.Run()

	go admin.InitServer(cfg, s)
	public.InitServer(cfg, p)
}

package main

import (
	"github.com/shura1014/cfg"
	"github.com/shura1014/cfg/g"
)

func main() {
	config, err := cfg.LoadConfig("./cfg", "app.yaml")
	if err != nil {
		g.Error(err)
	}
	ip := config.GetString("app.server.ip")
	port := config.GetInt("app.server.port")
	enable := config.GetBool("app.server.cors.enable")
	timeout := config.GetTime("app.server.timeout")
	languages := config.GetArray("app.server.languages")

	g.Info(ip)
	g.Info(port)
	g.Info(enable)
	g.Info(timeout)
	g.Info(languages)
	all, _ := config.GetAll()
	g.Info(all)
}

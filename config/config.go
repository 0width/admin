package config

import "git.xios.club/xios/gc"

type WebConfig struct {
	Port int `value:"${web.port:=8080}"`
}

func init() {
	gc.RegisterBean(new(WebConfig))
}

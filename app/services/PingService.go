package services

import "iris.dom/app/config/values/environment"

type PingService interface {
	Pong() *map[string]string
}

func NewPingService(env environment.EnvironmentValue) PingService {
	return &pinger{
		env: env,
	}
}

type pinger struct {
	env environment.EnvironmentValue
}

func (p *pinger) Pong() *map[string]string {
	m := make(map[string]string)
	m["message"] = p.env.String()
	return &m
}

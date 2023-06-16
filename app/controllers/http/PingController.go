package http

import "iris.dom/app/services"

type PingController struct {
	Service services.PingService
}

type PingResponse struct {
	Pong string `json:"pong"`
}

func (p *PingController) Get() (PingResponse, error) {
	m := p.Service.Pong()

	return PingResponse{
		Pong: (*m)["message"],
	}, nil
}

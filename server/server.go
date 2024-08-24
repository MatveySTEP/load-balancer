package server

import "load_balancer/client"

type SpartimilluServer struct {
	client client.Client
}

func NewSpartimilluServer(client client.Client) *SpartimilluServer {
	return &SpartimilluServer{client: client}
}

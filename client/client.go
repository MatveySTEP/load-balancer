package client

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client interface {
	ForwardRequest(req *gin.Context) *http.Response
	HealthCheck()
}

type Clients struct {
	conf           ConfClient
	counter        int
	healthyServers map[string]bool
	mu             sync.Mutex
}

func NewClients(conf ConfClient) *Clients {
	return &Clients{conf: conf, healthyServers: make(map[string]bool)}
}

func (s *Clients) ForwardRequest(c *gin.Context) *http.Response {
	for {
		if len(s.healthyServers) == 0 {
			s.HealthCheck()
		}

		s.mu.Lock()
		index := s.counter % len(s.conf.addresses)
		address := s.conf.addresses[index]
		s.counter++

		if s.healthyServers[address] == true {
			s.mu.Unlock()
			switch c.Request.Method {
			case http.MethodGet:
				return sendGetRequestToAnotherServer(address + c.Request.RequestURI)
			}
		}
		s.mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Clients) HealthCheck() {
	for _, address := range s.conf.addresses {
		resp, err := http.Get(address)

		s.mu.Lock()
		if err == nil && resp.StatusCode == http.StatusOK {
			s.healthyServers[address] = true
		} else {
			s.healthyServers[address] = false
		}
		s.mu.Unlock()
	}
}

func sendGetRequestToAnotherServer(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Can't read the response resp from the GET request")
	}
	return resp
}

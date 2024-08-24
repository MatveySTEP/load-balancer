package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"load_balancer/client"
	"net/http"
)

type Server struct {
	client client.Client
}

func NewServer(client client.Client) *Server {
	return &Server{client: client}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, c *gin.Context) {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		fmt.Fprint(c.Writer, fmt.Sprintf("Received request from %s\n", c.Request.RemoteAddr))
		fmt.Fprint(c.Writer, fmt.Sprintf("%s / %s\n", c.Request.Method, c.Request.Proto))
		fmt.Fprint(c.Writer, fmt.Sprintf("Host: %s\n", c.Request.Host))
		fmt.Fprint(c.Writer, fmt.Sprintf("User-Agent: %s\n", c.Request.Header.Get("User-Agent")))
		fmt.Fprint(c.Writer, fmt.Sprintf("Accept: %+v\n\n", c.Request.Header.Get("Accept")))
		fmt.Fprint(c.Writer, fmt.Sprintf("Replied with a hello message\n"))
		fmt.Fprintf(c.Writer, "Hello From Backend Server")

	})

	resp := s.client.ForwardRequest(c)

	fmt.Printf("Response from server: %s %s\n\n", resp.Proto, resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading the response body", http.StatusInternalServerError)
	}

	stringBody := string(body)
	fmt.Fprint(w, stringBody)
	fmt.Println(stringBody)

}

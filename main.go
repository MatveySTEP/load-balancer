package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Fprint(c.Writer, fmt.Sprintf("Received request from %s\n", c.Request.RemoteAddr))
		fmt.Fprint(c.Writer, fmt.Sprintf("%s / %s\n", c.Request.Method, c.Request.Proto))
		fmt.Fprint(c.Writer, fmt.Sprintf("Host: %s\n", c.Request.Host))
		fmt.Fprint(c.Writer, fmt.Sprintf("User-Agent: %s\n", c.Request.Header.Get("User-Agent")))
		fmt.Fprint(c.Writer, fmt.Sprintf("Accept: %+v\n\n", c.Request.Header.Get("Accept")))
		fmt.Fprint(c.Writer, fmt.Sprintf("Replied with a hello message\n"))
		fmt.Fprintf(c.Writer, "Hello From Backend Server")

	})
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("Error listening and serve: ", err)
	}

}

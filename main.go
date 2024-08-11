package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Printf("Received request from %s\n", c.Request.RemoteAddr)
		fmt.Printf("%s / %s\n", c.Request.Method, c.Request.Proto)
		fmt.Printf("Host: %s\n", c.Request.Host)
		fmt.Printf("User-Agent: %s\n", c.Request.Header.Get("User-Agent"))
		fmt.Printf("Accept: %+v\n\n", c.Request.Header.Get("Accept"))
		fmt.Printf("Replied with a hello message\n")
		fmt.Fprintf(c.Writer, "Hello From Backend Server")

	})
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("Error listening and serve: ", err)
	}

}

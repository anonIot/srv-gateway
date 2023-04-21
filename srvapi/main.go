package main

import (
	"fmt"

	"github.com/ambelovsky/gosf"
)

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage(request.Message.Text)
}

func init() {
	// Listen on an endpoint
	gosf.Listen("echo", echo)
}

func main() {
	// Start the server using a basic configuration
	fmt.Println("Socket start port : 9999")
	gosf.Startup(map[string]interface{}{"port": 9999})
}

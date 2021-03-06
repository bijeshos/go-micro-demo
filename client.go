package main

import (
	"context"
	"fmt"

	pg "github.com/bijeshos/go.microservices.demo/proto/greeter"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := pg.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &pg.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)
}

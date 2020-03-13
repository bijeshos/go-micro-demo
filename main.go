package main

import (
	"context"
	"fmt"

	pc "github.com/bijeshos/go.microservices.demo/proto/contact"
	pg "github.com/bijeshos/go.microservices.demo/proto/greeter"
	"github.com/micro/go-micro/v2"
)

/*
Example usage of top level service initialisation
*/

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pg.Request, rsp *pg.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Setup and the client
func runClient(service micro.Service) {
	// Create new greeter client
	greeter := pg.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &pg.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Msg)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	initAddress() //added as a placeholder

	// Register handler
	pg.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

func initAddress() {
	contact := pc.Contact{}
	contact.FirstName = "John"
	contact.LastName = "Doe"
	contact.Email = "joh.doe@example.com"
	contact.Uid = 1001
	fmt.Println(contact)
}

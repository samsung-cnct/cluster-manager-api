package main

import (
	"fmt"
	"log"

	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client starting")

	address := "localhost"
	port := 9050
	serverAddr := fmt.Sprintf("%s:%d", address, port)
	name := "Bob"

	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewClusterClient(conn)

	r, err := c.HelloWorld(context.Background(), &pb.HelloWorldMsg{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}

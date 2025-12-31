package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// configure the endpoint where it will receive request
	// What the server will do when a request comes on that endpoint
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server!")
	})

	// Handlers are functions that handle the bussiness logic for an endpoint
	// Handler will be executed when that endpoint receives a request from a client

	const port string = ":8080" // localhost - using port 3000 on our computer
	// Accepting external requests on port 3000 on our computer

	// listen and serve the requests
	// Server listens to requests on a specific port
	// Requests are entering through port 3000 and responses are also getting sent out from the port 3000
	// Listen - accepting requests
	// Serve means sending responses
	// Server receives requests and sends responses
	fmt.Println("Server Listening on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}

}

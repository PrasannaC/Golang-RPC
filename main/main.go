package main

import (
	"Go_RPC/RPC_Client"
	"Go_RPC/RPC_Server"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Main function to run the application
func main() {
	fmt.Println("Starting server...")

	// Create the object which is intended to be exposed by the RPC server
	stringOps := RPC_Server.StringOps{}

	// Register the created object
	rpc.Register(&stringOps)

	// Set the server to handle HTTP request formats.
	rpc.HandleHTTP()

	// Announce on local network and listener object
	l, e := net.Listen("tcp", ":6550")
	if e != nil {
		log.Fatal("error while listening: ", e)
	}

	// Server the client call in a new goroutine
	go http.Serve(l, nil)

	/*
	Calling client here for example.
	Should be called from another host for practical usage.
	 */
	RPC_Client.Client("localhost", "6550")
}

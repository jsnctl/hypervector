package main

import "github.com/jsnctl/hypervector/pkg/server"

func main() {
	server := server.NewServer(nil)
	server.RunServer()
}

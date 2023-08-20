package main

import "example.com/test/server"

func main() {
	server.NewServer(":3000").Init()
}

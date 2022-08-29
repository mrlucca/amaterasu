package main

import "amaterasu/src/tcp"

func main() {
	server := tcp.NewServer()
	server.Run()
}

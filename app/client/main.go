package main

import "os"

func main() {
	client := ClientServer{
		Port: os.Getenv("CLIENT_PORT"),
	}
	client.Run()
}

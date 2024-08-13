package main

import "os"

func main() {
	Run(os.Getenv("MATHS_API_PORT"))
}

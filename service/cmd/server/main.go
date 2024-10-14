package main

import "github.com/tanveerprottoy/jenkins-pipeline/service/internal/server"

// entry point
func main() {
	a := server.NewApp()
	a.Run()
}

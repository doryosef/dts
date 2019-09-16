package main

import (

	"./command"
	"./listener"
	"./logger"
)

func main() {
	log := logger.GetInstance()
	log.Println("starting app")
	listener.StartTelnet()
	listener.StartRestServer()
	command.Start()
}
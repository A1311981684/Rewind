package main

import (
	"errors"
	"log"
	"rewind/models"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	models.LoadConfigures()

	switch models.Configure.Role {
	case models.ROLE_SERVER:
		models.StartServer(models.LocalIP, models.Configure.LocalPort)
	case models.ROLE_CLIENT:
		models.ConnectServer(models.Configure.TargetIP, models.Configure.TargetPort)
	default:
		panic(errors.New("invalid parameter: configure.Role - " + models.Configure.Role))
	}
}

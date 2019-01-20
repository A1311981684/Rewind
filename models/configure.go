package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configurePath = "./conf/conf.json"
const (
	ROLE_SERVER = "server"
	ROLE_CLIENT = "client"
)
type AppConfig struct {
	Role          string
	LocalPort     int
	ClientTimeOut int
	RetryCount    int
	TargetIP      string
	TargetPort    int
}
var Configure AppConfig

func LoadConfigures() {
	jsonBytes, err := ioutil.ReadFile(configurePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonBytes, &Configure)
	if err != nil {
		panic(err)
	}
	log.Printf("\n--->This configure has been loaded:\n\tAppRole:%s\n\tLocalPort:%d\n\tClientTimeOut:%d\n\t" +
		"RetryCount:%d\n\tTargetIP:%s\n\tTargetPort:%d", Configure.Role, Configure.LocalPort, Configure.ClientTimeOut,
		Configure.RetryCount, Configure.TargetIP, Configure.TargetPort)
}

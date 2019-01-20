package models

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var conn net.Conn
var readChan = make(chan float64)
var stopChan = make(chan bool)

func ConnectServer(ip string, port int){
	tcpAddr, err := net.ResolveTCPAddr("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("server is not available now")
		return
	}
	log.Println("connection to server has been established")
	go handleConnClient()
	go deviceMock()
	<- make(chan int)
}

func handleConnClient(){

	go readConn(conn, readChan, stopChan)

	var run = true
	for ; run; {
		select {
		case readStr := <-readChan:
			log.Println("received from server:", readStr)
		case stop := <-stopChan:
			if stop {
				err := conn.Close()
				if err != nil {
					log.Println(err)
				}
				run = false
				break
			}
		default:
			continue
		}

	}
	log.Println("handle finished")
}

func sendData(data string){
	_, err := conn.Write([]byte(data))
	if err != nil {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
		log.Println(err)
	}
}

func deviceMock(){
	for {
		randomData := strconv.FormatFloat(rand.Float64()*1.5, 'f', 6, 64)
		sendData(randomData)
		time.Sleep(time.Millisecond * 50)
	}
}
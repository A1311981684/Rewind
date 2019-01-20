package models

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

const DISCONNECT_FLAG = "vYSyONySB3BhS0jb1Kz25GoEwKIKBVn61DU3CflLLvukGR5W3NBIN0HLDEv6HoPC"

var keepListen bool
var LocalIP = "127.0.0.1"

func StartServer(ip string, port int) {
	keepListen = true
	address := ip + ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	for ; keepListen; {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Println("client in...:", conn.RemoteAddr())
		go handleConnServer(conn)
	}
}

func handleConnServer(conn net.Conn) {

	readChan := make(chan float64)
	stopChan := make(chan bool)

	go readConn(conn, readChan, stopChan)

	var run = true
	for ; run; {
		select {
		case readStr := <-readChan:
			log.Println(readStr)
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

func readConn(conn net.Conn, rChan chan float64, sChan chan bool) {
	readIn := make([]byte, 1024)
	for {
		readSize, err := conn.Read(readIn)
		if err != nil {
			log.Println(err.Error())
			break
		}

		rData := string(readIn[:readSize])

		if rData == DISCONNECT_FLAG {
			sChan <- true
			break
		}
		rChan <- verifiedData(rData)
	}
}

func verifiedData(data string) float64 {
	f, err := strconv.ParseFloat(data, 64)
	if err == nil {
		return f
	}
	return -3.3
}

func writeConn(conn net.Conn, data string) {
	_, err := conn.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

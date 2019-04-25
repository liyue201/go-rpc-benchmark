package main

import (
	"log"
	"net"
	"net/rpc"
)

type EchoServer bool

func (s *EchoServer) Echo(req *string, res *string) error {
	return nil
}

func main() {
	echo := new(EchoServer)
	if err := rpc.Register(echo); err != nil {
		log.Println(err)
		return
	}

	var listener *net.TCPListener
	if tcpAddr, err := net.ResolveTCPAddr("tcp", ":8000"); err != nil {
		log.Println(err)
		return
	} else {
		if listener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
			log.Println(err)
			return
		}
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
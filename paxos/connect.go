package paxos

import (
	"net"
	"github.com/containous/traefik/log"
)

func Conn(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}


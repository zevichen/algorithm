package paxos

import (
	"testing"
	"net"
	"time"
	"log"
)

func TestServer2(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	PrintPac(err)

	defer CloseListen(listener)

	to := make(chan []byte)

	go ServerReadMsg(to)

	for {
		conn, e := listener.Accept()
		if PrintError(e) {
			continue
		}
		go Handle(conn, to)
	}

}
func ServerReadMsg(bytes chan []byte) {
	for {
		select {
		case b := <-bytes:
			log.Println("serverreadmsg")
			log.Println(string(b))
		}
	}
}
func Handle(conn net.Conn, to chan []byte) {
	log.Println("handle")
	temp := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if PrintError(err) {
			break
		}
		if n <= 0 {
			continue
		}
		bytes := buf[:n]

		if len(temp) > 0 {
			bytes = append(temp, bytes...)
		}
		log.Println("unpack")
		temp = Unpack(bytes, to)
	}
}

func TestClient2(t *testing.T) {
	conn, e := net.Dial("tcp", "127.0.0.1:8080")
	PrintPac(e)

	defer CloseConn(conn)

	to := make(chan []byte)

	go ClientWriteMsg(to)

	for {
		select {
		case b := <-to:
			log.Println("pack:", string(b))
			conn.Write(Pack(b))
		}
	}
}
func ClientWriteMsg(bytes chan []byte) {
	for {
		log.Println("clientwritemsg")
		bytes <- []byte(RdmString(3))
		time.Sleep(time.Second)
	}
}

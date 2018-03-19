package paxos

import (
	"testing"
	"net"
	"time"
	"log"
	"io/ioutil"
	"fmt"
)

func TestAcceptor_Run(t *testing.T) {

	addrs := []string{
		"127.0.0.1:8090",
		"127.0.0.1:8091",
		"127.0.0.1:8092",
		"127.0.0.1:8093",
		"127.0.0.1:8094",
	}
	log.Println("starting...")
	for i := range addrs {
		go NewAcceptor().Run(addrs[i])
	}

	time.Sleep(time.Hour)
}

func TestNewAcceptor(t *testing.T) {
	addr := "127.0.0.1:8090"

	conn, e := net.Dial("tcp", addr)
	PrintPac(e)
	conn.Write(Pack(ToBytes(&Msg{
		Id:    10,
		Value: "",
		Addr:  addr,
	})))

	bytes, _ := ioutil.ReadAll(conn)
	to := make(chan []byte)
	Unpack(bytes, to)
	fmt.Println(string(<-to))
	CloseConn(conn)

	time.Sleep(time.Second * 2)

	conn, e = net.Dial("tcp", addr)
	PrintPac(e)
	conn.Write(Pack(ToBytes(&Msg{
		Id:    10,
		Value: "helloValue",
		Addr:  addr,
	})))

	bytes, _ = ioutil.ReadAll(conn)
	Unpack(bytes, to)
	fmt.Println(string(<-to))
	CloseConn(conn)

}

package paxos

import (
	"testing"
	"net"
	"github.com/containous/traefik/log"
	"time"
	"io/ioutil"
	"strconv"
	"fmt"
)

func TestServer(t *testing.T) {
	addrs := []string{
		"127.0.0.1:8090",
		"127.0.0.1:8091",
		"127.0.0.1:8092",
		"127.0.0.1:8093",
		"127.0.0.1:8094",
	}

	for i := range addrs {

		go func(ad []string, i int) {
			lsn, e := net.Listen("tcp", ad[i])
			PrintError(e)
			for {
				conn, i2 := lsn.Accept()
				PrintPac(i2)

				go func(conn net.Conn) {
					to := make(chan []byte)
					go readMsg(to)
					for {
						bytes, i3 := ioutil.ReadAll(conn)
						if len(bytes) <= 0 {
							time.Sleep(time.Second)
							continue
						}
						PrintError(i3)

						Unpack(bytes, to)
					}
				}(conn)
			}
		}(addrs, i)
	}

	time.Sleep(time.Hour)
}

func readMsg(to chan []byte) {
	for r := range to {
		log.Println("message:" + string(r))
	}
}

func TestClient(t *testing.T) {
	addrs := []string{
		"127.0.0.1:8090",
		"127.0.0.1:8091",
		"127.0.0.1:8092",
		"127.0.0.1:8093",
		"127.0.0.1:8094",
	}
	for i := range addrs {
		go func(ad []string, i int) {
			for {
				func() {
					c, e := net.Dial("tcp", addrs[i])
					if e != nil {
						fmt.Print(e.Error())
						return
					}

					defer CloseConn(c)
					c.Write(Pack([]byte("【" + strconv.Itoa(i) + "】")))
					time.Sleep(time.Microsecond * 200)
				}()
			}
		}(addrs, i)
	}
	time.Sleep(time.Hour)
	//conn, e := net.Dial("tcp", addrs[0])
	//PrintPac(e)
	//conn.Write([]byte("hello"))
	//defer conn.Close()
}

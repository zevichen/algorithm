package learner

import (
	"net"
	"io/ioutil"
	"log"
	"github.com/zevichen/algorithm/paxos"
	"time"
)

func Learner() {
	to := make(chan []byte)

	go func(t chan []byte) {
		for {
			select {
			case b := <-t:
				log.Println("【learner.msg】", string(b))
			case <-time.After(time.Second * 10):
				return
			}
		}
	}(to)

	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	paxos.PrintPac(err)
	defer paxos.CloseListen(listener)

	temp := make([]byte, 0)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		bytes, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(temp) != 0 {
			bytes = append(temp, bytes...)
		}
		temp = paxos.Unpack(bytes, to)
	}
}

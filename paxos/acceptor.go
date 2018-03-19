package paxos

import (
	"net"
	"sync"
	"log"
)

type Acceptor struct {
	Id      int64
	Value   string
	mutex   sync.Mutex
	Learner net.Conn
	Addr    string
}

func NewAcceptor() *Acceptor {
	return &Acceptor{}
}

func (acceptor *Acceptor) Run(addr string) {
	acceptor.Addr = addr
	log.Printf("acceptor run:%s", addr)

	ln, e := net.Listen("tcp", addr)
	PrintPac(e)
	defer ln.Close()

	learn, e := net.Dial("tcp", "127.0.0.1:8888")
	PrintPac(e)
	acceptor.Learner = learn

	for {
		conn, err := ln.Accept()
		if PrintError(err) {
			continue
		}

		go acceptor.HandleConn(conn)

	}
}

func (acceptor *Acceptor) HandleConn(conn net.Conn) {
	log.Println("acceptor.Handle")

	temp := make([]byte, 0)
	to := make(chan []byte)
	buf := make([]byte, 1024)
	sig := make(chan bool)

	go acceptor.handleRequest(conn, to, sig)

	for {
		n, err := conn.Read(buf)
		if PrintError(err) {
			CloseConn(conn)
			sig <- true
			break
		}
		if n <= 0 {
			continue
		}

		bytes := buf[:n]
		log.Println("bytes:", string(bytes))
		if len(temp) > 0 {
			bytes = append(temp, bytes...)
		}
		temp = Unpack(bytes, to)
	}
}

func (acceptor *Acceptor) handleRequest(conn net.Conn, to chan []byte, sig chan bool) {

	for {
		select {
		case msg := <-to:
			log.Println("msg:", string(msg))
			push := &Msg{}
			ToStruct(msg, push)

			acceptor.mutex.Lock()
			if acceptor.Id <= push.Id {
				if IsBlank(push.Value) {
					acceptor.Value = push.Value
				}
				acceptor.Id = push.Id

				go acceptor.noticeLearner()
			}
			acceptor.mutex.Unlock()

			conn.Write(Pack(ToBytes(&Msg{
				Id:    acceptor.Id,
				Value: acceptor.Value,
				Addr:  acceptor.Addr,
			})))
		case s := <-sig:
			if s {
				CloseConn(conn)
				return
			}
		}

	}
}
func (acpt *Acceptor) noticeLearner() {
	acpt.Learner.Write(Pack(ToBytes(&Msg{
		Id:    acpt.Id,
		Value: acpt.Value,
		Addr:  acpt.Addr,
	})))
}

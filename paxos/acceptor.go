package paxos

import (
	"net"
	"github.com/containous/traefik/log"
	"io/ioutil"
	"time"
	"sync"
)

type Acceptor struct {
	Id         int64
	Value      string
	State      int
	PreTimeout time.Duration
	mutex      sync.Mutex
	Proposers  map[int64]net.Conn
}

func NewAcceptor() *Acceptor {
	return &Acceptor{
		State:      PRE_ACCEPTOR,
		PreTimeout: 5 << 13,
	}
}

func (acceptor *Acceptor) Run(addr string) {
	ln, e := net.Listen("tcp", addr)
	if e != nil {
		panic(e)
	}

	acceptorTicker(acceptor)

	for {
		conn, i := ln.Accept()
		if i != nil {
			continue
		}
		go acceptor.Handle(conn)
	}
}

func acceptorTicker(acceptor *Acceptor) {
	t := time.NewTicker(acceptor.PreTimeout)
	for {
		select {
		case <-t.C:
			acceptor.mutex.Lock()
			acceptor.State = ACCEPTOR
			acceptor.mutex.Unlock()
		}
	}
}

func (acceptor *Acceptor) Handle(conn net.Conn) {
	log.Println(ioutil.ReadAll(conn))
	//acceptor.Proposers[xx]=conn

}

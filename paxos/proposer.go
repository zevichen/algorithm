package paxos

import (
	"time"
	"net"
	"sync"
	"io/ioutil"
	"testing"
	"log"
	"math/rand"
)

type TakeFn func(conn net.Conn, to chan []byte)

type Proposer struct {
	Id     int64
	Value  string
	Health int
	Total  int
	Addrs  []string
	State  map[string]bool
	Port   string
	mutex  sync.Mutex
	Count  int
	log    *testing.T
}

func NewProposer(addrs []string, port string) *Proposer {

	log.Println(addrs, port)
	l := len(addrs)
	return &Proposer{
		Id:    int64(rand.Intn(5)),
		Addrs: addrs,
		Port:  port,
		Total: l,
		mutex: sync.Mutex{},
		State: make(map[string]bool, l),
	}
}

func (p *Proposer) Run() {

	go receiveMsg(p)

	broadcast(p)
}

func health(p *Proposer, addr string, state bool) {
	log.Println("health addr:", addr, "state:", state)
	p.mutex.Lock()

	if !state {
		go retry(addr, p)
	}
	p.mutex.Lock()
	p.State[addr] = state
	p.mutex.Unlock()

}

//func tick(addr string, p *Proposer) {
//	for k, v := range p.State {
//		if !v {
//			go retry(addr, p)
//		}
//	}
//}

func retry(addr string, p *Proposer) {

	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			if unicast(addr, p) {
				health(p, addr, true)
			} else {
				health(p, addr, false)
			}
		case <-time.After(time.Second * 10):
			return
		}
	}
}

func receiveMsg(p *Proposer) {

	to := make(chan []byte)

	dealRsp(to, p)

	log.Println("receiveMsg.unpack.loop")

	reconnect(p, to, func(conn net.Conn, to chan []byte) {
		defer CloseConn(conn)

		log.Println("receiveMsg.unpack.loop.single")
		temp := make([]byte, 0)
		for {
			log.Println("receiveMsg.unpack.loop.single.for")
			bytes, err := ioutil.ReadAll(conn)
			if PrintError(err) || len(bytes) <= 0 {
				time.Sleep(time.Second)
				continue
			}

			log.Println("receiveMsg.unpack.loop.single.for.unpack")
			if len(temp) != 0 {
				bytes = append(temp, bytes...)
			}
			temp = Unpack(bytes, to)
		}
	})

}
func dealRsp(t chan []byte, prop *Proposer) {
	go func(t chan []byte, p *Proposer) {
		log.Println("receiveMsg.go.start")

		for {
			select {
			case bytes := <-t:
				log.Println("receiveMsg readBytes", string(bytes))
				msg := &Msg{}
				if !IsBlank(msg.Value) {
					log.Println("ending...msg.Value", msg.Value)
					return
				}
				ToStruct(bytes, msg)
				handleRsp(p, msg)
			}
		}
	}(t, prop)
}

func handleRsp(proposer *Proposer, msg *Msg) {

	log.Println("handleRsp", proposer.Id, proposer.Value, &msg)

	proposer.mutex.Lock()
	if msg.Id > proposer.Id {
		proposer.Id = msg.Id
		if !IsBlank(msg.Value) {
			proposer.Value = msg.Value
		}
	} else {
		proposer.Count++
	}
	if proposer.Count >= proposer.Total/2+1 {
		proposer.Value = RdmString(5)
		log.Printf("【handleRsp.Proposal】 id:%d value:%s\n", proposer.Id, proposer.Value)
	}
	proposer.mutex.Unlock()
	broadcast(proposer)

	log.Println("handleRsp.end", msg.Id, msg.Value, msg.Addr, proposer)
}

func unicast(addr string, p *Proposer) bool {
	conn, e := Chat(addr)
	if PrintError(e) {
		return false
	}
	defer CloseConn(conn)

	conn.Write(Pack(ToBytes(&Msg{
		Id:    p.Id,
		Value: p.Value,
	})))

	return true

}

func defaultTakeFn(conn net.Conn, to chan []byte) {
	defer CloseConn(conn)

	conn.Write(Pack(<-to))
}

func broadcast(p *Proposer) {
	log.Println("proposor.broadcast health:", p.Health)

	to := make(chan []byte)
	reconnect(p, to, defaultTakeFn)

}

func reconnect(p *Proposer, to chan []byte, fn TakeFn) {
	for i := range p.Addrs {
		addr := p.Addrs[i]
		conn, e := net.Dial("tcp", addr)
		if PrintError(e) {
			health(p, addr, false)
			go retry(addr, p)
			return
		}
		health(p, addr, true)
		log.Println("proposor.broadcast.single pId:", p.Id, "p.Value:", p.Value)

		to <- Pack(ToBytes(&Msg{
			Id:    p.Id,
			Value: p.Value,
		}))
		go fn(conn, to)
	}
}

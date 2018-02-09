package paxos

import (
	"time"
	"net"
)

type Proposer struct {
	Id        int64
	Value     string
	State     int
	Acceptors map[string]net.Conn
	Health    int
	Total     int
	Addrs     []string
	Port      string
}

func NewProposer(addrs []string, port string) *Proposer {

	return &Proposer{
		Id:    time.Now().UnixNano(),
		Value: GetRandomString(10),
		State: PRE_PROPOSAL,
		Addrs: addrs,
		Port:  ":" + port,
	}
}

func (p *Proposer) Run() {

	acceptor := NewAcceptor()
	acceptor.Run(p.Port)

	acceptors := make(map[string]net.Conn)

	//检测
	health := 0
	for i := range p.Addrs {
		conn, err := Conn(p.Addrs[i])
		acceptors[p.Addrs[i]] = conn
		if err == nil {
			health++
		}
	}
	p.Health = health
	p.Total = len(p.Addrs)
	p.Acceptors = acceptors

	//预提交
	preProposal(p)

	//定时检测返回值

	//一段时间后修改状态,进入下一步

}

func preProposal(p *Proposer) {

	for k, v := range p.Acceptors {

		v.Write(ToBytes(&Push{
			p.Id,
			p.Value,
			k,
		}))

	}
}

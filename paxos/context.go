package paxos

const (
	_            = iota
	PRE_PROPOSAL
	PROPOSAL
	PRE_ACCEPTOR
	ACCEPTOR
)

type Push struct {
	Id    int64
	Value string
	Addr  string
}

type Pull struct {
	Id    int64
	Value string
}


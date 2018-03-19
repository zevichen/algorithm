package paxos

const (
	_            = iota
	PRE_PROPOSAL
	PROPOSAL
)

type Msg struct {
	Id    int64
	Value string
	Addr  string
}

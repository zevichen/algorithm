package paxos

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestPaxos(t *testing.T) {
	addrs := []string{
		"127.0.0.1:8080",
		"127.0.0.1:8081",
		"127.0.0.1:8082",
		"127.0.0.1:8083",
		"127.0.0.1:8084",
	}
	proposer := NewProposer(addrs, "8080")
	proposer.PreProposal()

}

type Object struct {
	Value string `json:"value"`
}

func TestOther(t *testing.T) {
	obj := &Object{
		Value: "aaa",
	}

	bytes := ToBytes(obj)
	obj2 := &Object{}
	ToStruct(bytes, obj2)
	marshal, _ := json.Marshal(obj2)
	fmt.Println(string(marshal))

}

func TestEndian(t *testing.T) {
	SystemEndian()
}

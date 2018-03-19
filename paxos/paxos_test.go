package paxos

import (
	"testing"
	"encoding/json"
	"fmt"
	"time"
	"log"
)

func TestProposer_Run(t *testing.T) {
	addrs := []string{
		"127.0.0.1:8090",
		"127.0.0.1:8091",
		"127.0.0.1:8092",
		"127.0.0.1:8093",
		"127.0.0.1:8094",
	}
	log.Println("starting...")
	for i := range addrs {
		go NewProposer(addrs, addrs[i]).Run()
	}

	time.Sleep(time.Hour)
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

func SubWork() {
	log.Println("fewi")
}

func TestMap(t *testing.T) {
	//m := make(map[string]bool)
	//m["1"] = true
	//m["2"] = false
	//
	//fmt.Println(len(m))
	//delete(m, "1")
	//fmt.Println(len(m))
	for i := 0; i < 5; i++ {
		time.Sleep(time.Microsecond)
		go SubWork()
	}

	time.Sleep(time.Second * 3)
}

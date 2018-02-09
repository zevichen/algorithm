package paxos

import (
	"math/rand"
	"time"
	"encoding/binary"
	"unsafe"
	"encoding/json"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ToBytes(data interface{}) []byte {
	//return *(*[]byte)(unsafe.Pointer(&data))

	marshal, e := json.Marshal(data)
	if e != nil {
		panic(e)
	}
	return []byte(marshal)
}

func ToStruct(data []byte, arg interface{}) {
	err := json.Unmarshal(data, arg)
	if err != nil {
		panic(err)
	}
}

const IntSize = int(unsafe.Sizeof(0))

func SystemEndian() binary.ByteOrder {
	var i = 0x1
	bs := (*[IntSize]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		return binary.BigEndian
	} else {
		return binary.LittleEndian
	}
}

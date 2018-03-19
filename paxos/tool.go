package paxos

import (
	"math/rand"
	"encoding/binary"
	"unsafe"
	"encoding/json"
	"strings"
	"reflect"
	"log"
	"bytes"
	"net"
)

func RdmString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(rand.Int63()))
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
		PrintError(err)
	}
}

func IsBlank(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}

const IntSize = int(unsafe.Sizeof(0))

func SystemEndian() binary.ByteOrder {
	var i = 0x1
	bs := (*[IntSize]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		return binary.LittleEndian
	} else {
		return binary.BigEndian
	}
}
func PrintInfo(in interface{}, info ...string) {
	switch reflect.TypeOf(in).Kind().String() {
	case "struct":
		json, err := json.Marshal(in)
		PrintError(err)
		log.Println(info, json)
	case "string":
		log.Println(info, in)
	default:
		log.Println(info, in)
	}
}
func PrintError(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
func PrintPac(err error) {
	if err != nil {
		panic(err)
	}
}

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, SystemEndian(), x)
	return bytesBuffer.Bytes()
}
func BytesToInt(b []byte) int {
	var x int32
	binary.Read(bytes.NewBuffer(b), SystemEndian(), &x)
	return int(x)
}
func CloseConn(c net.Conn) {
	if c != nil {
		if e := c.Close(); e != nil {
			log.Fatal("close conn error:", e.Error())
		}
	}
}
func CloseListen(lsn net.Listener) {
	if lsn != nil {
		if e := lsn.Close(); e != nil {
			log.Fatal("close conn error:", e.Error())
		}
	}
}

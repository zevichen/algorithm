package paxos

import (
	"testing"
	"log"
	"time"
	"strconv"
	"fmt"
	"unsafe"
)

func TestUnpack(t *testing.T) {
	toChan := make(chan []byte)

	//buf := make([]byte, 30)
	//log.Println("buf.len:", len(buf))
	//buf = strconv.AppendInt(buf, 1, 32)
	//tempBytes := Unpack(buf, toChan)
	//log.Println("tempBytes.len:", len(tempBytes), "tempBytes:", tempBytes)

	go func(t chan []byte) {

		for {
			select {
			case b := <-t:
				log.Println("msg:", string(b))
			}
		}
	}(toChan)

	buf := append(Pack([]byte("aaa")), Pack([]byte("bbb"))...)
	buf = append(buf, Pack([]byte("ccc"))...)

	tempBytes := Unpack(buf, toChan)
	log.Println("temp.len:", len(tempBytes), "temp:", tempBytes)

	time.Sleep(time.Second)
}

func TestAppendInt(t *testing.T) {
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

	fmt.Println("--------------")

	b := []byte{0, 0, 0, 53}
	v := *(*uint32)(unsafe.Pointer(&b[0]))
	fmt.Printf("0x%X\n", v)
	fmt.Printf("%d\n", v)
	fmt.Printf("%v\n", *(*[4]byte)(unsafe.Pointer(&v)))

	fmt.Println("-----------------------")
	b = []byte{0, 0, 0, 5}
	fmt.Println(b)
	fmt.Println(IntToBytes(5))
	fmt.Println(BytesToInt(b))
	fmt.Println("--------------------")

}

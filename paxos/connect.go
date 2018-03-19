package paxos

import (
	"net"
	"log"
)

func Chat(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func WriteMsg(conn net.Conn, msg Msg) {
	conn.Write(ToBytes(msg))
}

const (
	HeadLen = 4
)

func Pack(b []byte) []byte {
	buf := make([]byte, 0)
	l := IntToBytes(len(b))
	buf = append(buf, l...)
	buf = append(buf, b...)
	return buf
}

//解决tcp打包发送的问题（粘包）
func Unpack(from []byte, to chan []byte) []byte {
	length := len(from)
	if length <= HeadLen {
		return from
	}

	var i = 0
	for ; i < length; i++ {
		if HeadLen+i >= length {
			break
		}

		head := from[i:HeadLen+i]
		msgSize := HeadLen + BytesToInt(head)
		if length < msgSize+i {
			break
		}
		to <- from[i+HeadLen:i+msgSize]
		i += msgSize - 1
	}

	if i == length {
		return make([]byte, 0)
	}

	return from[i:]
}

package learner

import (
	"testing"
	"net"
	"time"
	"github.com/zevichen/algorithm/paxos"
)

func TestLearner(t *testing.T) {
	Learner()
}

//粘包问题。
//解决方法：
//1.定义简单协议[length,data]协议头来读取固定长度数据，未截取的缓存起来
//2.发送一次断开客户端链接，后续发送重新链接
func TestPut(t *testing.T) {
	conn, e := net.Dial("tcp", "127.0.0.1:8888")
	paxos.PrintPac(e)
	conn.Write(paxos.Pack([]byte("aaa")))
	conn.Close()
	time.Sleep(time.Microsecond)
	conn, e = net.Dial("tcp", "127.0.0.1:8888")
	paxos.PrintPac(e)
	conn.Write(paxos.Pack([]byte("bbb")))
	conn.Close()
}

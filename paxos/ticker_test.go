package paxos

import (
	"testing"
	"time"
	"fmt"
)

func TestTicker(t *testing.T) {

	task1 := SetTask(&GoodsTask{}, time.Second)
	task2 := SetTask(&GoodsTask{}, time.Second)
	task3 := SetTask(&GoodsTask{}, time.Second)
	task4 := SetTask(&GoodsTask{}, time.Second)
	task5 := SetTask(&GoodsTask{}, time.Second)

	task1.Name = "task1"
	task2.Name = "task2"
	task3.Name = "task3"
	task4.Name = "task4"
	task5.Name = "task5"
	group := SetGroup(task1, task2, task3, task4, task5)
	time.Sleep(time.Second * 5)
	group.Run()
	fmt.Println(group.Set,group.Dead)
	group.Status()
	time.Sleep(time.Second * 2)
	group.Close(task3, task5)
	fmt.Println(group.Set,group.Dead)
	group.Status()
	time.Sleep(time.Second * 2)
	group.Run(task3)
	fmt.Println(group.Set,group.Dead)
	group.Status()
	time.Sleep(time.Second * 2)
	group.Close()
	fmt.Println(group.Set,group.Dead)
	group.Status()

}

type GoodsTask struct {
	Value string
}

func (t *GoodsTask) TaskRun() {
	fmt.Println("goods task")
}

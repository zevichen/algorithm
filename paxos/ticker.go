package paxos

import (
	"time"
	"sync"
	"fmt"
	"log"
)

const (
	_     = iota
	CLOSE
	DONE
)

type Task interface {
	TaskRun()
}

type TickerGroup struct {
	Set   map[string]*TickerTask
	Dead  map[string]*TickerTask
	Name  string
	mutex sync.Mutex
}
type TickerTask struct {
	Task     Task
	CloseSig chan int
	DoneSig  chan int
	Name     string
	Time     time.Duration
}

func (g *TickerGroup) Close(tt ...*TickerTask) {
	if len(tt) == 0 {
		for k, v := range g.Set {
			g.mutex.Lock()
			if v.Close() {
				g.Dead[k] = v
				delete(g.Set, k)
			} else {
				//retry
				fmt.Println(v.Name, " didn't close")
			}
			g.mutex.Unlock()
		}
	} else {

		for i := range tt {
			g.mutex.Lock()
			t := tt[i]
			if t.Close() {
				g.Dead[t.Name] = t
				delete(g.Set, t.Name)
			} else {
				//retry
				fmt.Println(t.Name, " didn't close")
			}
			g.mutex.Unlock()
		}
	}
}
func DetectClose(t *TickerTask, g *TickerGroup) {

}

func (g *TickerGroup) Status() {
	fmt.Println(g.Name, "\n\tactive:", len(g.Set), "\n\tdead:", len(g.Dead))
}

func (g *TickerGroup) Run(tt ...*TickerTask) {
	if len(tt) == 0 {
		for _, v := range g.Set {
			v.Run()
		}
	}

	for i := range tt {
		t := g.Set[tt[i].Name]
		if t != nil {
			t.Run()
			continue
		} else {
			t := g.Set[tt[i].Name]
			if t != nil {
				t.Run()
				continue
			}
		}
		log.Println("the task don't exist")
	}
}

func SetTask(task Task, time time.Duration) *TickerTask {
	return &TickerTask{
		Task:     task,
		CloseSig: make(chan int),
		DoneSig:  make(chan int),
		Name:     RdmString(10),
		Time:     time,
	}
}
func SetGroup(ts ...*TickerTask) *TickerGroup {
	set := make(map[string]*TickerTask)
	for i := range ts {
		set[ts[i].Name] = ts[i]
	}

	return &TickerGroup{
		Set:  set,
		Dead: make(map[string]*TickerTask),
		Name: RdmString(10),
	}
}

func (tt *TickerTask) Close() bool {
	tt.CloseSig <- CLOSE

	after := time.After(time.Second * 2)

	select {
	case <-tt.DoneSig:
		return true
	case <-after:
		return false
	}
}

func (tt *TickerTask) Run() {
	//one time
	if tt.Time == 0 {
		go func() {
			for {
				tt.Task.TaskRun()
			}
		}()
	} else {
		go func() {
			ticker := time.NewTicker(tt.Time)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					log.Println(tt.Name, " is running")
					tt.Task.TaskRun()
				case <-tt.CloseSig:
					log.Println(tt.Name, " is end")
					tt.DoneSig <- DONE
					return
				}
			}
		}()
	}
}

/*

组
1.go所有组
2.组状态关闭所有任务
3.组定时任务改状态

任务
1.状态关闭
2.循环，周期执行


 */

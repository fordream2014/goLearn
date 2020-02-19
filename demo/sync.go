package demo

import (
	"bytes"
	"io"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
	"sync/atomic"
	"errors"
)

const (
	gosize = 10  //goroutine个数
	writesize = 10  //每个goroutine写入数据块的大小
)

//同步机制
func TestSyncWrite() {
	var buf bytes.Buffer
	var mu sync.Mutex
	var sign = make(chan struct{}, gosize)
	for i:=0; i<gosize; i++ {

		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()

			mu.Lock()
			defer mu.Unlock()

			for j:=0; j<writesize; j++ {
				header := fmt.Sprintf("id: %d, current step: %d \n", id, j)
				_, err := writer.Write([]byte(header))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}(i, &buf)
	}

	for i:=0; i<gosize; i++ {
		<- sign
	}

	//读出buffer内容
	data, err := ioutil.ReadAll(&buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

type counter struct {
	num int
	rwlock sync.RWMutex
}

func (c *counter) count() int {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()

	return c.num
}

func (c *counter) add() {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()
	time.Sleep(time.Microsecond)
	c.num = c.num + 1
	//fmt.Println(c.num)
}

//读写锁
func TestRWsynLock() {
	var c = &counter{}
	//fmt.Println(c.num)
	//计数器
	sign := make(chan struct{}, 3)
	go func(id int) {
		defer func() {
			sign<-struct {}{}
		}()

		for i:=0; i<10; i++ {
			time.Sleep(time.Microsecond)
			fmt.Println(id, c.count())
		}
	}(1)

	go func(id int) {
		defer func() {
			sign<-struct {}{}
		}()

		for i:=0; i<10; i++ {
			time.Sleep(time.Microsecond)
			fmt.Println(id, c.count())
		}
	}(2)

	go func(id int) {
		defer func() {
			sign<-struct {}{}
		}()

		for i:=0; i<10; i++ {
			c.add()
			//fmt.Println(3)
		}
	}(3)

	for i:=0; i<3; i++ {
		<-sign
	}
}

//发信收信流程
func TestSendReceiveMail() {

	var lock = sync.RWMutex{}
	sendCond := sync.NewCond(&lock)
	receiveCond := sync.NewCond(lock.RLocker())

	var mail = 0
	var sign = make(chan struct{}, 2)
	//发信
	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for i:=0; i<5; i++ {
			lock.Lock()
			for mail == 1 {
				sendCond.Wait()
			}
			mail = 1
			fmt.Println("信箱中信件空了，放进入信件")
			lock.Unlock()
			receiveCond.Signal()
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for i:=0; i<5; i++ {
			lock.RLock()
			for mail == 0 {
				receiveCond.Wait()
			}
			mail = 0
			fmt.Println("信箱中还有信件，读取信件")
			lock.RUnlock()
			sendCond.Signal()
		}
	}()

	<- sign
	<- sign

}

//原子操作
func TestAtmic() {
	var num int32
	atomic.AddInt32(&num, 12)
	fmt.Println(num)

	var sign = make(chan struct{}, 5)

	for i:=0; i<5; i++ {
		go func(num *int32) {
			defer func() {
				sign <- struct{}{}
			}()
			for {
				curNum := atomic.LoadInt32(num)
				newNum := curNum + 2
				if atomic.CompareAndSwapInt32(num, curNum, newNum) {
					fmt.Println(newNum)
					break
				}
			}
		}(&num)
	}

	for i:=0; i<5; i++ {
		<-sign
	}
	fmt.Println("最终结果", num)
}

//不能用原子值存储nil
//我们向原子值存储的第一个值，决定了它今后能且只能存储哪一个类型的值。
//存储一个接口类型的值，然后再存储这个接口的某个实现类型的值，这样是不是可以呢？这是不可以的。原子值内部是依据被存储值的实际类型来做判断的
func TestAtmicValue() {
	var box atomic.Value
	box2 := box
	v1 := []int{1,2,3,4}
	box.Store(v1)
	fmt.Println(box.Load())
	fmt.Println(box2.Load())
	fmt.Println()

	v2 := "hello"
	box2.Store(v2)
	fmt.Println(box2.Load())
	fmt.Println()

	box3 := box
	fmt.Println(box3.Load())
	v3 := []int{1,2}
	box3.Store(v3)
	fmt.Println(box3.Load(), box.Load())
	fmt.Println()
}

func TestSynOnce() {
	var counter uint32
	var once sync.Once
	once.Do(func() {
		atomic.AddUint32(&counter, 1)
	})
	fmt.Printf("The counter: %d\n", counter)
	once.Do(func() {
		atomic.AddUint32(&counter, 2)
	})
	fmt.Printf("The counter: %d\n", counter)
	fmt.Println()

	once = sync.Once{}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		once.Do(func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("Do task. [1-%d]\n", i)
				time.Sleep(time.Second)
			}
		})
		fmt.Println("Done. [1]")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [2]")
		})
		fmt.Println("Done. [2]")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [3]")
		})
		fmt.Println("Done. [3]")
	}()
	wg.Wait()
	fmt.Println()

	// 示例3。
	once = sync.Once{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("fatal error: %v\n", p)
			}
		}()
		once.Do(func() {
			fmt.Println("Do task. [4]")
			panic(errors.New("something wrong"))
			fmt.Println("Done. [4]")
		})
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [5]")
		})
		fmt.Println("Done. [5]")
	}()
	wg.Wait()
}

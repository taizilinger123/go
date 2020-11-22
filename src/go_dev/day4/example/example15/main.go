package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//互斥锁：写很多用互斥锁
var lock sync.Mutex

//读写锁，读多写少用读写锁
var rwlock sync.RWMutex

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}

	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
}

func testRwlock() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			//rwlock.Lock()
			lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
			// rwlock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				lock.Lock()
				// rwlock.RLock()
				time.Sleep(time.Millisecond)
				// fmt.Println(a)
				// rwlock.RUnlock()
				lock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	//testMap()
	testRwlock()
}

//PS D:\project> go build -race  go_dev/day4/example/example15
//PS D:\project> example15.exe
//go  get github.com/go-sql-driver/mysql   //下载包

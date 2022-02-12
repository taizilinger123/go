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
//go  get github.com/go-sql-driver/mysql   //下载包,这个下载包都会下载到D:\project\src目录里
/*
D:\project>go  build -race go_dev/day4/example/example15

D:\project>dir
 驱动器 D 中的卷是 Data
 卷的序列号是 DC8C-C090

 D:\project 的目录

2020/11/23  03:28    <DIR>          .
2020/11/23  03:28    <DIR>          ..
2020/11/14  22:46    <DIR>          .vscode
2020/11/14  22:39    <DIR>          bin
2020/11/23  03:26         2,137,600 example1.exe
2020/11/23  03:31         3,195,392 example15.exe
2020/11/23  01:51         3,193,856 main.exe
2020/11/23  02:57    <DIR>          pkg
2020/11/23  03:10    <DIR>          src
2020/11/23  03:21    <DIR>          vender
               3 个文件      8,526,848 字节
               7 个目录 77,601,824,768 可用字节

D:\project>example15.exe
1726
---------------------------------------------------------
win10环境变量设置:
shenglei的用户变量
GOPATH   D:\project
系统变量
PATH:新增
D:\go\bin
---------------------------------------------------------
D:\project>go  env
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\shenglei\AppData\Local\go-build
set GOENV=C:\Users\shenglei\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=D:\project\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=D:\project      ***********************
set GOPRIVATE=
set GOPROXY=https://goproxy.cn,direct
set GOROOT=D:\go           *********************
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=D:\go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\shenglei\AppData\Local\Temp\go-build182248274=/tmp/go-build -gno-record-gcc-switches
---------------------------------------------------------------------
D:\project 的目录
2020/11/23  03:32    <DIR>          .
2020/11/23  03:32    <DIR>          ..
2020/11/14  22:46    <DIR>          .vscode
2020/11/14  22:39    <DIR>          bin
2020/11/23  02:57    <DIR>          pkg
2020/11/23  03:10    <DIR>          src
2020/11/23  03:21    <DIR>          vender
--------------------------------------------------------------------
D:\go>dir
 D:\go 的目录
2020/11/14  21:17    <DIR>          bin
2020/11/14  21:17    <DIR>          pkg
2020/11/14  21:17    <DIR>          src
*/

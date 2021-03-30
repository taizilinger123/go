package main

import (
	"flag"
	"fmt"
)

func main() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&logLevel, "d", 10, "please input log level")

	flag.Parse()
	fmt.Println("path:", confPath)
	fmt.Println("log level:", logLevel)
}

// PS D:\project> go run "d:\project\src\go_dev\day7\example\example7\main.go"
// path:
// log level: 10
// PS D:\project> go run "d:\project\src\go_dev\day7\example\example7\main.go" -c  c:/test.conf  -d  16
// path: c:/test.conf
// log level: 16

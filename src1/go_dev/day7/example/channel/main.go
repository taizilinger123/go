package main

func main() {
	var intChan chan int
	intChan = make(chan int, 10) //上面写的是chan int,所以下面这里也是chan int
	intChan <- 10
}

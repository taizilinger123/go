package main  

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i 
	}
}

func recv(ch chan int) {
    for {
		 
	}
}

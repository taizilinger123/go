package main

func main() {
	var mapChan chan map[string]string         //chan map[string]string
	mapChan = make(chan map[string]string, 10) //因为上面写的是chan map[string]string，所以这里也是写chan map[string]string，
	//map[string]map[string]map[string]string,这里的意思是:这里是一个map,键是string，值是map[string]map[string]string,值又分键是string值是map[string]string,键是string值是string
	m := make(map[string]string, 16) //这里写16不只是存16个元素，map自己会扩容的
	m["stu01"] = "001"
	m["stu02"] = "002"

	mapChan <- m
}

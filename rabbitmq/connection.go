package main

import "fmt"

func Run() {
	list := []int{1, 2, 3, 4, 5, 6}

	ch := make(chan bool, 10)
	for _, v := range list {
		fmt.Println(123)
		ch <- true
		go func () {
			fmt.Println(v)

			<- ch
		}()
	}
}

func main() {
	go Run()


}

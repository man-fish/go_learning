package main

import "fmt"

func chanDemo() {
	//var c chan int
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2
}

func main(){
	chanDemo()
}
package main

import "fmt"

func main(){
	defer func() {
		if r := recover();r != nil {
			//fmt.Println(r)
		}
	}()

	fmt.Println(1)
}

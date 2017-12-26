package main

import "fmt"

func main() {

	c := make(chan int)
	a := []int{1, 2, 3,}
	for _, v := range a {
		go func(val int) {
			fmt.Println(val)
		}(v)
		fmt.Println(v)
	}

}

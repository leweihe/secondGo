package main

import "fmt"

func PingPong(s []int) {
	s = append(s, 3)
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)
	PingPong(s)
	fmt.Println(s)
}

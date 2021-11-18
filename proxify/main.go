package main

import (
	"fmt"
	"reflect"
)

func Comp(array1 []int, array2 []int) bool {
	// your code
	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}
	a1 := make(map[int]int)
	a2 := make(map[int]int)

	for _, a := range array1 {
		if _, ok := a1[a*a]; ok {
			a1[a*a] += 1
		} else {
			a1[a*a] = 1
		}
	}

	for _, a := range array2 {
		if _, ok := a2[a]; ok {
			a2[a] += 1
		} else {
			a2[a] = 1
		}
	}
	eq := reflect.DeepEqual(a1, a2)
	if eq {
		return true
	}
	return false
}

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	Comp(a1, a2)
	/*c := make(chan int)
	l := []int{10, 2, 3}
	for _, v := range l {
		go writeToChannel(c, v)
	}
	fmt.Println("The Value 4 of X: ", l)
	fmt.Println("The Value 5 of X: ", <-c)
	fmt.Println("Ended")
	time.Sleep(1 * time.Second)*/
	//close(c)
}

func writeToChannel(c chan int, x int) {
	for {
		fmt.Println("The Value 1 is: ", x)
		if x < 10 {
			c <- x
		} else {
			x -= 1
		}
		fmt.Println("The Value 2 is: ", x)
	}
}

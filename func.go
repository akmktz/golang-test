package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	s := sum(arr[0:])
	fmt.Println(s)

	fmt.Println("======================================")

	n, result := half(4)
	fmt.Println(n, result)
	n, result = half(5)
	fmt.Println(n, result)

	fmt.Println("======================================")

	s = add(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(s)

	fmt.Println("======================================")

	defer func() {
		str := recover()
		fmt.Println("---------------------------")
		fmt.Println(str)
		fmt.Println("---------------------------")
	}()
	panic("PANIC")
}

func add(args ...int) int {
	sum := 0
	for _, num := range args {
		sum += num
	}

	return sum
}

func half(num int) (int, bool) {

	return num / 2, num%2 == 0
}

func sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

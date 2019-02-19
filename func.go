package main

import "fmt"

func main() {
	fmt.Println("============ SUM ===============")

	arr := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	s := sum(arr[0:])
	fmt.Println(s)

	fmt.Println("============ HALF ==============")

	n, result := half(4)
	fmt.Println(n, result)
	n, result = half(5)
	fmt.Println(n, result)

	fmt.Println("============ MAX ===============")

	n = max(1, 2, 31, 4, 5, 6, 7)
	fmt.Println(n)

	fmt.Println("============ ODD ===============")

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println("============ FIB ===============")

	n = fib(5)
	fmt.Println(n)
	fmt.Println("============ PANIC =============")

	defer func() {
		str := recover()
		fmt.Println("---------------------------")
		fmt.Println(str)
		fmt.Println("---------------------------")
	}()
	panic("PANIC")
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}

	return n + fib(n-1)
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func max(args ...uint) uint {
	max := args[0]
	for _, num := range args {
		if num > max {
			max = num
		}
	}

	return max
}

func half(num uint) (uint, bool) {

	return num / 2, num%2 == 0
}

func sum(nums []uint) uint {
	sum := uint(0)

	for _, num := range nums {
		sum += num
	}

	return sum
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("============ FOR1 ==============")
	var nums1 []uint

	for i := uint(0); i < 100; i += 1 {
		if i%3 == 0 {
			nums1 = append(nums1, i)
		}
	}
	fmt.Println(nums1)

	fmt.Println("============ FOR2 ==============")

	nums2 := make(map[uint]string)

	for i := uint(0); i < 100; i += 1 {
		if i%3 == 0 {
			nums2[i] = "Fizz"
		} else if i%5 == 0 {
			nums2[i] = "Buzz"
		} else {
			nums2[i] = fmt.Sprint(i)
		}
	}
	fmt.Println(nums2)

	fmt.Println("============ MAP ===============")

	flights := InitData()
	fmt.Println(&flights)

	id := FindFlight("Dublin", "Kherson", &flights)
	fmt.Println(id)
}

func InitData() map[string]map[string]string {

	flights := map[string]map[string]string{
		"1": map[string]string{
			"from": "Dublin",
			"to":   "Tsurich",
		},
		"2": map[string]string{
			"from": "Dublin",
			"to":   "Kherson",
		},
	}

	return flights
}

func FindFlight(from, to string, flights *map[string]map[string]string) string {

	for id, value := range *flights {
		if value["from"] == from && value["to"] == to {
			return id
		}
	}

	return ""
}

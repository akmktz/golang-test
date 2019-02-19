package main

import "fmt"

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

func main() {
	flights := InitData()
	fmt.Println(&flights)
	fmt.Println("---------------------------------------------")

	id := FindFlight("Dublin", "Kherson", &flights)
	fmt.Println(id)
}

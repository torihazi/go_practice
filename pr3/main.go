package main

import "fmt"

type MyIntSlice []int

func main() {
	m := MyIntSlice{4, 2, 2, 3, 3, 3, 4, 5, 1, 2, 4, 5}
	fmt.Println(m)
	fmt.Println(m.Unique())

}

func (m MyIntSlice) Unique() []int {
	var prevValueMap map[int]bool = make(map[int]bool)
	var result []int
	for _, currentValue := range m {
		if _, ok := prevValueMap[currentValue]; !ok {
			prevValueMap[currentValue] = true
			result = append(result, currentValue)
		}
	}
	return result
}

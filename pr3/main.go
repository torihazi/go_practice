package main

import "fmt"

type MyIntSlice []int

func main() {
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m)
	fmt.Println(m.Unique())

}

func (m MyIntSlice) Unique() []int {
	var prevValue int
	var result []int
	for index, currentValue := range m {
		if index == 0 || prevValue != currentValue {
			prevValue = currentValue
			result = append(result, currentValue)
		}
	}
	return result
}

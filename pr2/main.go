package main

import (
	"errors"
	"fmt"
)

func main() {
	m := map[int]string{
		1: "01",
		2: "02",
		3: "03",
	}

	key1, err1 := findKeyByValue(m, "03") // key→3, err→nil
	key2, err2 := findKeyByValue(m, "05") // key→0にすること(初期値なので), errはある

	fmt.Println(key1, err1)
	fmt.Println(key2, err2)
}

func findKeyByValue(m map[int]string, value string) (int, error) {
	for key, val := range m {
		if value == val {
			return key, nil
		}
	}
	return 0, errors.New("No key!!")
}

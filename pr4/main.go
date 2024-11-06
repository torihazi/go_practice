package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	// for _, user := range users {
	// 	user.Age = 44
	// }

	for i, user := range users {
		fmt.Printf("&users[%d]: %p\n", i, &users[i])
		fmt.Printf("&user: %p\n", &user)
	}

	for i := range users {
		users[i].Age = 44
	}

	fmt.Printf("%v", users) // どうなる？
}

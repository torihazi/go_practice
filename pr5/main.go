package main

import "fmt"
var appConfig = Config{Env: "test"}

type Config struct {
	Env string
}

func getConfig() *Config {
	return &appConfig
}

func main() {
	c := getConfig()
	c.Env = "production"

	// fmt.Printf("&c: %p\n", c)
	// fmt.Printf("&appconfig: %p\n", &appConfig)

	fmt.Println(c.Env)     // production
	fmt.Println(appConfig.Env) // testではなくproducionへ
}
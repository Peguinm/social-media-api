package main

import (
	"fmt"

	dotenv "github.com/joho/godotenv"
	"sm.com/m/app"
)

func main() {
	err := LoadEnvironment()
	if err != nil {
		fmt.Printf("error loading env file: %v\n", err)
	}

	app.Run()
}

func LoadEnvironment() (err error) {
	err = dotenv.Load()
	if err != nil {
		return
	}
	return
}

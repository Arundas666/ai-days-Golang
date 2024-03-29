package main

import (
	"fmt"
	"genAi-with-go/helper"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	for {

		helper.Selection()
	}


}

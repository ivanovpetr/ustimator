package main

import (
	"fmt"
	"os"

	"ivanovpetr/ustimator/cmd"
)

func main() {
	err := cmd.New().Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

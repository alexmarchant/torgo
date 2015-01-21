package main

import (
	"fmt"
	"github.com/alexmarchant/torgo/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

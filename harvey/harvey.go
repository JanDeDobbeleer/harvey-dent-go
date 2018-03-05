package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println(getLogo())
}

func getLogo() (logo string) {
	b, err := ioutil.ReadFile("harvey.txt")
	if err != nil {
		fmt.Print(err)
	}
	logo = string(b)
	return
}

package main

import "fmt"

type account struct {
	id            int
	branch        string
	accountNumber int
	name          string
	balance       int
}

func main() {
	fmt.Printf("hello, world\n")
	a1 := account{
		id: 1,
		branch: "TenjinCity",
		accountNumber: 1234567,
		name: "Test Account",
		balance: 1000000,
	}
	fmt.Println(a1)
}

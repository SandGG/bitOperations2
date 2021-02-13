package main

import "fmt"

func main() {
	//int8 = 0 to 2^7 - 1 = (-127, 127)
	var res int8

	//And
	res = 0 & 0
	fmt.Println("0 & 0: ", res)
	res = 0 & 1
	fmt.Println("0 & 1: ", res)
	res = 1 & 0
	fmt.Println("1 & 0: ", res)
	res = 1 & 1
	fmt.Println("1 & 1: ", res) //1

	//Or
	res = 0 | 0
	fmt.Println("0 | 0: ", res)
	res = 0 | 1
	fmt.Println("0 | 1: ", res) //1
	res = 1 | 0
	fmt.Println("1 | 0: ", res) //1
	res = 1 | 1
	fmt.Println("1 | 1: ", res) //1

	//XOr
	res = 0 ^ 0
	fmt.Println("0 ^ 0: ", res)
	res = 0 ^ 1
	fmt.Println("0 ^ 1: ", res) //1
	res = 1 ^ 0
	fmt.Println("1 ^ 0: ", res) //1
	res = 1 ^ 1
	fmt.Println("1 ^ 1: ", res)

	//Shift left
	res = 5
	fmt.Printf("Number: %08b\n", res)
	for i := 1; i < 6; i++ { // i = 5, res = 160? No, res = -96: 160 - 256 = -96
		fmt.Printf("Shift left: %08b\n", res<<i)
	}

	var resN int = 5 //11111111 -> 255
	fmt.Printf("Number (int): %08b\n", res)
	fmt.Printf("Shift left (int): %08b\n", resN<<5) // res = 160? Yes

	//Shift right (keep the sign)
	res = 45
	fmt.Printf("Number: %08b\n", res)
	for i := 1; i < 6; i++ {
		fmt.Printf("Shift right: %08b\n", res>>i)
	}
}

package main

import "fmt"

func main() {
	fmt.Printf("Enter data: ")
	var data string
	fmt.Scanln(&data)

	crc := CalculateChecksum([]byte(data))
	fmt.Println(crc)
}

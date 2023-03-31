package main

import "fmt"

func main() {
	kata := "selamat malam"

	for _, x := range kata {
		fmt.Printf("%c\n", x)
	}
	huruf := make(map[string]int)

	for _, x := range kata {
		huruf[string(x)] += 1
	}
	fmt.Println(huruf)
}

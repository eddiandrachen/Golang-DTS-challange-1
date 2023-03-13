package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}
	for num := 0; num < 10; num++ {
		if num == 5 {
			var ucode = []map[string]string{
				{"code": "U+0421", "char": "С", "num": "0"},
				{"code": "U+0410", "char": "А", "num": "2"},
				{"code": "U+0428", "char": "Ш", "num": "4"},
				{"code": "U+0410", "char": "А", "num": "6"},
				{"code": "U+0420", "char": "Р", "num": "8"},
				{"code": "U+0412", "char": "В", "num": "10"},
				{"code": "U+041E", "char": "О", "num": "12"},
			}
			for _, ucode := range ucode { //; i = 0; i <= 10; i++{
				fmt.Printf("character %s '%s' start as byte position %s\n", ucode["code"], ucode["char"], ucode["num"])
			}
			//fmt.Println("Nilai j =", num)

		} else {
			fmt.Println("Nilai j =", num)
		}

	}
	//for j := 6; j < 11; j++ {
	//	fmt.Println("Nilai j =", j)
	//}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			d := []rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}

			for tx := 0; tx < len(d); tx++ {
				char := d[tx]
				fmt.Printf("character %U '%s' start as byte position %d\n", char, string(char), tx*2)
			}
		} else {
			fmt.Println("Nilai J = ", j)
		}
	}
}

//ucode := []rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'} //"САШАРВО"
//[]rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'} //[]rune{"\u0421", "\u0410", "\u0428", "\u0410", "\u0420", "\u0412", "\u041E"}
//fmt.Printf("Unicode codepoint: %U\n", []rune(ucode))
//for i := range ucode {
//fmt.Printf("Character %s '%U'", i, ucode)

//}

package main

import "fmt"

//menampilkan tipe data dari variabel i menampilkan tanda %
//menampilkan nilai boolean j : true
//menampilkan nilai boolean j : true
//menampilkan unicode russia : Я (ya)
//menampilkan nilai base 10 : 21
//menampilkan nilai base 8 :25
//menampilkan nilai base 16 : f
//menampilkan nilai base 16 : F 13
//menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
//menampilkan float : 123.456000
//menampilkan float scientific : 1.234560E+02
func main() {

	b := 21
	h := 21
	i := "%"
	j := true
	a := "Я"
	e := 15
	t := 15
	n := 19
	d := 123.456000
	fmt.Printf("%T\n", i)
	s := "Я"
	fmt.Printf("Unicode codepoint: %U\n", []rune(s))

	fmt.Println(b)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(j)
	fmt.Println(a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%o\n", h)
	fmt.Printf("%x\n", e)
	fmt.Printf("%X\n", t)
	fmt.Printf("%x\n", n)
	fmt.Printf("%f\n", d)
	fmt.Printf("%e\n", d)

}

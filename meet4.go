package main

import (
	"fmt"
	"os"
)

// use os.Args
func main() {
	defer errInfo()

	argc := len(os.Args)
	if argc <= 0 {
		fmt.Println("Default", os.Args[0:], "default")
		os.Exit(1)
	}
	fmt.Println("Default", os.Args[0])
	fmt.Println("Nama : Calman\n. Alamat : sungai nil\n. pekerjaan : Mentor jago\n. Alasan memilih kelas golang : Suka suki", os.Args[1])
	fmt.Println("Nama : masa iaa\n. Alamat : indonesia\n. pekerjaan : Mahasiswa\n. Alasan memilih kelas golang unknown.", os.Args[2])
	fmt.Println("Nama : Digidaw\n. Alamat : indonesia\n. pekerjaan : Mahasiswa\n. Alasan memilih kelas golang unknown.", os.Args[3])
	fmt.Println("Nama : Aselole\n. Alamat : indonesia\n. pekerjaan : Mahasiswa\n. Alasan memilih kelas golang unknown.", os.Args[4])
	fmt.Println("The program argc is:", argc)
}
func errInfo() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("No error occured")
	}
}

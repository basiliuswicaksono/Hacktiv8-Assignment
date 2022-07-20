package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type students []student

func main() {
	arg := os.Args[1]
	stnds := students{
		{
			Nama:      "Lius",
			Alamat:    "Bogor",
			Pekerjaan: "Backend",
			Alasan:    "Karena golang sedang populer",
		},
		{
			Nama:      "Adrian",
			Alamat:    "Bekasi",
			Pekerjaan: "Backend",
			Alasan:    "Karena golang sedang populer",
		},
		{
			Nama:      "Abdul",
			Alamat:    "Tangerang",
			Pekerjaan: "Backend",
			Alasan:    "Karena golang sedang populer",
		},
	}

	stnds.Print(arg)

}

func (st students) Print(arg string) {
	//string to int
	intVar, _ := strconv.Atoi(arg)

	if intVar > len(st) {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Println("Nama :", st[intVar-1].Nama)
	fmt.Println("Alamat :", st[intVar-1].Alamat)
	fmt.Println("Pekerjaan :", st[intVar-1].Pekerjaan)
	fmt.Println("Alasan memilih kelas golang :", st[intVar-1].Alasan)
}

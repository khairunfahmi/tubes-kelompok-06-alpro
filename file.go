package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func simpanKeFile() {
	file, err := os.Create("pengeluaran.json")
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membuat file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(pengeluarans)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat menyimpan data ke file:", err)
	}
	fmt.Println("Data berhasil disimpan ke file.")
}

func muatDariFile() {
	file, err := os.Open("pengeluaran.json")
	if err != nil {
		fmt.Println("File tidak ditemukan atau terjadi kesalahan:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pengeluarans)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat memuat data dari file:", err)
	}
	fmt.Println("Data berhasil dimuat dari file.")
}

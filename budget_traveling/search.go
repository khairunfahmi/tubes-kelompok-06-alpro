package main

import (
	"fmt"
	"sort"
	"strings"
)

func cariKategoriSequential() {
	var cari string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scan(&cari)
	cari = strings.ToLower(cari)

	found := false
	for _, p := range pengeluarans {
		if strings.ToLower(p.Kategori) == cari {
			fmt.Printf("Ditemukan: [%s] Rp %.2f\n", p.Kategori, p.Jumlah)
			found = true
		}
	}
	if !found {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func cariKategoriBinary() {
	if len(pengeluarans) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	sort.Slice(pengeluarans, func(i, j int) bool {
		return pengeluarans[i].Kategori < pengeluarans[j].Kategori
	})

	var cari string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scan(&cari)

	low, high := 0, len(pengeluarans)-1
	for low <= high {
		mid := (low + high) / 2
		if pengeluarans[mid].Kategori == cari {
			fmt.Printf("Ditemukan: [%s] Rp %.2f\n", pengeluarans[mid].Kategori, pengeluarans[mid].Jumlah)
			return
		} else if pengeluarans[mid].Kategori < cari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Kategori tidak ditemukan.")
}

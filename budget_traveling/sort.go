package main

import "fmt"

func urutkanKategoriSelection() {
	n := len(pengeluarans)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if pengeluarans[j].Kategori < pengeluarans[min].Kategori {
				min = j
			}
		}
		pengeluarans[i], pengeluarans[min] = pengeluarans[min], pengeluarans[i]
	}
	fmt.Println("Pengeluaran berhasil diurutkan berdasarkan kategori.")
}

func urutkanJumlahInsertion() {
	for i := 1; i < len(pengeluarans); i++ {
		key := pengeluarans[i]
		j := i - 1
		for j >= 0 && pengeluarans[j].Jumlah > key.Jumlah {
			pengeluarans[j+1] = pengeluarans[j]
			j--
		}
		pengeluarans[j+1] = key
	}
	fmt.Println("Pengeluaran berhasil diurutkan berdasarkan jumlah.")
}

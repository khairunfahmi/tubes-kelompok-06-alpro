package main

import (
	"fmt"
)

var pengeluarans []Pengeluaran
var totalBudget float64

func main() {
	var pilihan int
	fmt.Print("Masukkan total budget perjalanan Anda: Rp ")
	fmt.Scan(&totalBudget)

	for {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Tambah Data Pengeluaran")
		fmt.Println("2. Tampilkan Data Pengeluaran")
		fmt.Println("3. Edit Data Pengeluaran")
		fmt.Println("4. Hapus Data Pengeluaran")
		fmt.Println("5. Cari Data Pengeluaran (Sequential Search)")
		fmt.Println("6. Cari Data Pengeluaran (Binary Search)")
		fmt.Println("7. Urutkan Berdasarkan Kategori (Selection Sort)")
		fmt.Println("8. Urutkan Berdasarkan Jumlah (Insertion Sort)")
		fmt.Println("9. Tampilkan Laporan")
		fmt.Println("10. Simpan Data ke File")
		fmt.Println("11. Muat Data dari File")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			tampilkanData()
		case 3:
			editData()
		case 4:
			hapusData()
		case 5:
			cariKategoriSequential()
		case 6:
			cariKategoriBinary()
		case 7:
			urutkanKategoriSelection()
		case 8:
			urutkanJumlahInsertion()
		case 9:
			laporan()
		case 10:
			simpanKeFile()
		case 11:
			muatDariFile()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

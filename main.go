package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pengeluaran struct {
	Kategori string
	Jumlah   int
}

var pengeluaran []Pengeluaran
var totalBudget int

func simpanData() {
	file, _ := os.Create("data.json")
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(pengeluaran)
}

func bacaData() {
	file, err := os.Open("data.json")
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode(&pengeluaran)
}

func tambahData() {
	var kategori string
	var jumlah int
	fmt.Print("Masukkan kategori pengeluaran: ")
	fmt.Scan(&kategori)
	fmt.Print("Masukkan jumlah pengeluaran: ")
	fmt.Scan(&jumlah)
	pengeluaran = append(pengeluaran, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	simpanData()
	fmt.Println("Data berhasil ditambahkan!")
}

func tampilkanData() {
	fmt.Println("\nDaftar Pengeluaran:")
	for i, p := range pengeluaran {
		fmt.Printf("%d. %s - Rp%d\n", i+1, p.Kategori, p.Jumlah)
	}
}

func editData() {
	tampilkanData()
	var index int
	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&index)
	if index <= 0 || index > len(pengeluaran) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	var kategori string
	var jumlah int
	fmt.Print("Kategori baru: ")
	fmt.Scan(&kategori)
	fmt.Print("Jumlah baru: ")
	fmt.Scan(&jumlah)
	pengeluaran[index-1] = Pengeluaran{Kategori: kategori, Jumlah: jumlah}
	simpanData()
	fmt.Println("Data berhasil diubah!")
}

func hapusData() {
	tampilkanData()
	var index int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&index)
	if index <= 0 || index > len(pengeluaran) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	pengeluaran = append(pengeluaran[:index-1], pengeluaran[index:]...)
	simpanData()
	fmt.Println("Data berhasil dihapus!")
}

func cariKategoriSequential(kategori string) {
	ditemukan := false
	for _, p := range pengeluaran {
		if strings.ToLower(p.Kategori) == strings.ToLower(kategori) {
			fmt.Printf("Kategori: %s - Rp%d\n", p.Kategori, p.Jumlah)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func cariKategoriBinary(kategori string) {
	sort.Slice(pengeluaran, func(i, j int) bool {
		return strings.ToLower(pengeluaran[i].Kategori) < strings.ToLower(pengeluaran[j].Kategori)
	})
	kiri := 0
	kanan := len(pengeluaran) - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		kat := strings.ToLower(pengeluaran[tengah].Kategori)
		if kat == strings.ToLower(kategori) {
			fmt.Printf("Kategori: %s - Rp%d\n", pengeluaran[tengah].Kategori, pengeluaran[tengah].Jumlah)
			return
		} else if kat < strings.ToLower(kategori) {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println("Kategori tidak ditemukan.")
}

func urutkanKategoriSelection() {
	n := len(pengeluaran)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(pengeluaran[j].Kategori) < strings.ToLower(pengeluaran[min].Kategori) {
				min = j
			}
		}
		pengeluaran[i], pengeluaran[min] = pengeluaran[min], pengeluaran[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan kategori (Selection Sort).")
}

func urutkanJumlahInsertion() {
	for i := 1; i < len(pengeluaran); i++ {
		key := pengeluaran[i]
		j := i - 1
		for j >= 0 && pengeluaran[j].Jumlah > key.Jumlah {
			pengeluaran[j+1] = pengeluaran[j]
			j--
		}
		pengeluaran[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan jumlah (Insertion Sort).")
}

func laporan() {
	tampil := make(map[string]int)
	total := 0
	for _, p := range pengeluaran {
		tampil[p.Kategori] += p.Jumlah
		total += p.Jumlah
	}
	fmt.Println("\nLaporan Pengeluaran:")
	for kat, jumlah := range tampil {
		fmt.Printf("- %s: Rp%d\n", kat, jumlah)
	}
	fmt.Printf("Total pengeluaran: Rp%d\n", total)
	selisih := totalBudget - total
	fmt.Printf("Anggaran tersedia: Rp%d\n", totalBudget)
	fmt.Printf("Selisih: Rp%d\n", selisih)
	if selisih < 0 {
		fmt.Println("⚠️  Pengeluaran melebihi anggaran, coba kurangi hiburan atau makanan.")
	} else {
		fmt.Println("👍  Kamu masih dalam batas anggaran.")
	}
}

func main() {
	bacaData()
	fmt.Print("Masukkan total budget kamu: Rp")
	fmt.Scan(&totalBudget)
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Lihat Pengeluaran")
		fmt.Println("3. Edit Pengeluaran")
		fmt.Println("4. Hapus Pengeluaran")
		fmt.Println("5. Cari Kategori (Sequential)")
		fmt.Println("6. Cari Kategori (Binary)")
		fmt.Println("7. Urutkan Kategori (Selection Sort)")
		fmt.Println("8. Urutkan Jumlah (Insertion Sort)")
		fmt.Println("9. Lihat Laporan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		var pilih int
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahData()
		case 2:
			tampilkanData()
		case 3:
			editData()
		case 4:
			hapusData()
		case 5:
			var kat string
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kat)
			cariKategoriSequential(kat)
		case 6:
			var kat string
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kat)
			cariKategoriBinary(kat)
		case 7:
			urutkanKategoriSelection()
		case 8:
			urutkanJumlahInsertion()
		case 9:
			laporan()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

package main

import "fmt"

type Pengeluaran struct {
	Kategori string
	Jumlah   float64
}

func tambahData() {
	var kategori string
	var jumlah float64

	fmt.Print("Masukkan kategori pengeluaran: ")
	fmt.Scan(&kategori)
	fmt.Print("Masukkan jumlah pengeluaran: Rp ")
	fmt.Scan(&jumlah)

	pengeluarans = append(pengeluarans, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	fmt.Println("Data berhasil ditambahkan.")
}

func tampilkanData() {
	fmt.Println("\n--- Daftar Pengeluaran ---")
	for i, p := range pengeluarans {
		fmt.Printf("%d. [%s] Rp %.2f\n", i+1, p.Kategori, p.Jumlah)
	}
}

func editData() {
	var index int
	tampilkanData()
	fmt.Print("Pilih nomor data yang ingin diedit: ")
	fmt.Scan(&index)

	if index <= 0 || index > len(pengeluarans) {
		fmt.Println("Index tidak valid.")
		return
	}

	var kategoriBaru string
	var jumlahBaru float64
	fmt.Print("Kategori baru: ")
	fmt.Scan(&kategoriBaru)
	fmt.Print("Jumlah baru: Rp ")
	fmt.Scan(&jumlahBaru)

	pengeluarans[index-1] = Pengeluaran{Kategori: kategoriBaru, Jumlah: jumlahBaru}
	fmt.Println("Data berhasil diubah.")
}

func hapusData() {
	var index int
	tampilkanData()
	fmt.Print("Pilih nomor data yang ingin dihapus: ")
	fmt.Scan(&index)

	if index <= 0 || index > len(pengeluarans) {
		fmt.Println("Index tidak valid.")
		return
	}

	pengeluarans = append(pengeluarans[:index-1], pengeluarans[index:]...)
	fmt.Println("Data berhasil dihapus.")
}

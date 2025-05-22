# Aplikasi Pengelolaan Budget Traveling (Upgrade)

Aplikasi ini dirancang untuk membantu traveler mengelola anggaran perjalanan secara efektif dan juga praktis. Cocok untuk mahasiswa atau siapa pun yang ingin mengontrol pengeluaran selama bepergian.

## Fitur Utama

1. Tambah, ubah, dan hapus pengeluaran perjalanan  
2. Hitung total pengeluaran & sisa anggaran secara otomatis  
3. Cari pengeluaran berdasarkan kategori dengan:
- Sequential Search
- Binary Search
4. Urutkan data pengeluaran menggunakan:
- Selection Sort (kategori)
- Insertion Sort (jumlah)  
5. Laporan akhir pengeluaran vs budget  

---

## Pemberitahuan Upgrade

Versi terbaru aplikasi ini telah menambahkan fitur untuk **menyimpan dan memuat data pengeluaran dari file**. Pengguna dapat menyimpan data pengeluaran dalam file JSON dan memuatnya kembali ke aplikasi untuk melanjutkan pengelolaan anggaran perjalanan tanpa kehilangan data.

---

## Struktur File

| File         | Deskripsi                          |
|--------------|------------------------------------|
| `main.go`    | Entry point aplikasi               |
| `data.go`    | CRUD data pengeluaran              |
| `search.go`  | Fitur pencarian kategori           |
| `sort.go`    | Fitur pengurutan kategori/jumlah   |
| `laporan.go` | Menampilkan laporan akhir perjalanan |
| `go.mod`     | Modul Go                           |

---

## Cara Menjalankan

```bash
go run main.go data.go search.go sort.go laporan.go file.go

package main

import "fmt"

func laporan() {
	fmt.Println("\n--- Laporan Pengeluaran ---")
	kategoriMap := make(map[string]float64)
	total := 0.0

	for _, p := range pengeluarans {
		kategoriMap[p.Kategori] += p.Jumlah
		total += p.Jumlah
	}

	for kategori, jumlah := range kategoriMap {
		fmt.Printf("- %s: Rp %.2f\n", kategori, jumlah)
	}

	fmt.Printf("\nTotal Pengeluaran: Rp %.2f\n", total)
	fmt.Printf("Total Budget: Rp %.2f\n", totalBudget)
	fmt.Printf("Sisa Budget: Rp %.2f\n", totalBudget-total)

	if total > totalBudget {
		fmt.Println("⚠️  Anda telah melebihi budget!")
	} else if totalBudget-total < 0.2*totalBudget {
		fmt.Println("⚠️  Sisa budget sangat sedikit, pertimbangkan penghematan.")
	} else {
		fmt.Println("✅  Budget Anda aman.")
	}
}

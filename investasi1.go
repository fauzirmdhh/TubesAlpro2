package main

import (
	"fmt"
	"strings"
)

const NMAX int = 1024

const (
	hargaEmas      = 2000000 // Harga per gram emas
	hargaSaham     = 500000  // Harga per unit saham
	hargaReksaDana = 150000  // Harga per unit reksa dana
)

type Investasi struct {
	Nama   string
	Harga  float64
	Jumlah int
}

var daftarInvestasi []Investasi

func tambahInvestasi(nama string, jumlah int) {
	var harga float64
	switch strings.ToLower(nama) { // Menggunakan ini karena setiap inputan pengguna akan menjadi kecil jadi tidak akan ada bias
	case "emas":
		harga = hargaEmas
	case "saham":
		harga = hargaSaham
	case "reksa dana":
		harga = hargaReksaDana
	default:
		fmt.Println("🚫 Jenis investasi tidak dikenal.")
		return
	}

	daftarInvestasi = append(daftarInvestasi, Investasi{Nama: nama, Harga: harga, Jumlah: jumlah}) // append digunakan untuk menanbahkan elemen baru dalam list
	fmt.Println("✅ Investasi berhasil ditambahkan!")
}

func cetakData() {
	if len(daftarInvestasi) == 0 {
		fmt.Println("📭  Tidak ada data investasi.")
		return
	}
	fmt.Println("\n📁  Daftar Investasi Saat Ini:")
	fmt.Println("------------------------------------------------")
	fmt.Printf("%-15s %-10s %-10s\n", "📌 Nama", "💰 Harga", "📦 Jumlah")
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(daftarInvestasi); i++ {
		inv := daftarInvestasi[i]
		fmt.Printf("%-15s %-10.2f %-10d\n", inv.Nama, inv.Harga, inv.Jumlah)
	}
	fmt.Println("------------------------------------------------")
}

func linearSearch(nama string) int {
	for i := 0; i < len(daftarInvestasi); i++ {
		inv := daftarInvestasi[i]
		if strings.ToLower(inv.Nama) == strings.ToLower(nama) {
			return i
		}
	}
	return -1
}

func menu() {
	var M int
	var nama string

	for {
		fmt.Println("\n=========================================")
		fmt.Println("📊  APLIKASI INVESTASI SEDERHANA")
		fmt.Println("=========================================")
		fmt.Println("1️⃣  Tambah Investasi")
		fmt.Println("2️⃣  Tampilkan Semua Investasi")
		fmt.Println("3️⃣  Jual / Tarik Investasi (❌ Hapus)")
		fmt.Println("0️⃣  Keluar")
		fmt.Println("-----------------------------------------")
		fmt.Print("🧭  Pilih menu: ")

		fmt.Scan(&M)

		switch M {
		case 1:
			// Menu tambah investasi
			fmt.Print("Masukkan nama investasi (Emas/Saham/Reksa Dana): ")
			fmt.Scan(&nama)
			nama = strings.TrimSpace(nama) // untuk menghilangkan spasi atau enter dalam inputan nama

			fmt.Print("Masukkan jumlah investasi: ")
			var jumlah int
			fmt.Scan(&jumlah)

			tambahInvestasi(nama, jumlah)
		case 2:
			// Tampilkan semua investasi
			cetakData()
		case 3:
			// Jual atau tarik investasi
			fmt.Print("Masukkan nama investasi yang ingin dijual/tarik: ")
			fmt.Scan(&nama)
			nama = strings.TrimSpace(nama)

			index := linearSearch(nama)
			if index != -1 {
				// Hapus data investasi dari daftar
				daftarInvestasi = append(daftarInvestasi[:index], daftarInvestasi[index+1:]...)
				fmt.Println("✅ Investasi berhasil dihapus.")
			} else {
				fmt.Println("🚫 Investasi tidak ditemukan.")
			}
		case 0:
			// Keluar dari program
			fmt.Println("👋 Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("⚠️  Menu tidak valid. Coba lagi.")
		}
	}
}

func main() {
	menu()
}

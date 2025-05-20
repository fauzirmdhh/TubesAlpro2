package main

import (
	"fmt"
	"strings"
)

const NMAX int = 20

var akun Akun
var daftarInvestasi [NMAX]Investasi
var nData int
var isLoggedIn bool = false
var idCounter int = 100

type Akun struct {
	Username string
	Password string
	Deposit float64
}

type Investasi struct {
	ID     int
	Nama   string
	Jumlah float64
	Total  float64
}

func main() {
	var pilihan int

	for {
		if !isLoggedIn {
			tampilkanHeader("🟢 FinVest - Aplikasi Investasi Sederhana")
			fmt.Println("1. Buat Akun")
			fmt.Println("2. Login")
			fmt.Println("3. Keluar")
			fmt.Print(">> Pilihan Anda: ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case 1:
				buatAkun()
			case 2:
				login()
			case 3:
				fmt.Println("Terima kasih telah menggunakan FinVest!")
				return
			default:
				fmt.Println("🚫 Pilihan tidak valid.\n")
			}
		} else {
			menuUtama()
		}
	}
}

func buatAkun() {
	fmt.Println("\n📌 Buat Akun Baru")
	fmt.Print("Username: ")
	fmt.Scanln(&akun.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&akun.Password)
	fmt.Println("✅ Akun berhasil dibuat!\n")
}

func login() {
	var LoginUsn, LoginPass string

	fmt.Println("\n🔐 Login")
	fmt.Print("Username: ")
	fmt.Scanln(&LoginUsn)
	fmt.Print("Password: ")
	fmt.Scanln(&LoginPass)

	if LoginUsn == akun.Username && LoginPass == akun.Password {
		isLoggedIn = true
		fmt.Println("✅ Login berhasil!\n")
		fmt.Print("Deposit : ")
		fmt.Scanln(&akun.Deposit)
	} else {
		fmt.Println("🚫 Username atau password salah\n")
	}
}

func tampilkanHeader(judul string) {
	fmt.Println(strings.Repeat("=", 45))
	fmt.Println(judul)
	fmt.Println(strings.Repeat("=", 45))
}

func menuUtama() {
	var pilihan int

	tampilkanHeader("📊 FinVest - Menu Utama")
	fmt.Println("1. Tambah Investasi")
	fmt.Println("2. Jual/Tarik Investasi")
	fmt.Println("3. Urutkan Berdasarkan Harga (Ascending)")
	fmt.Println("4. Urutkan Berdasarkan Harga (Descending)")
	fmt.Println("5. Urutkan Berdasarkan ID (Ascending)")
	fmt.Println("6. Lihat Daftar Investasi")
	fmt.Println("7. Logout")
	fmt.Print(">> Pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahInvestasi(&daftarInvestasi, &nData, akun)
	case 2:
		tarikInvestasi(&daftarInvestasi, &nData)
	case 3:
		sortByHargaAsc(&daftarInvestasi, nData)
	case 4:
		sortByHargaDsc(&daftarInvestasi, nData)
	case 5:
		sortByID(&daftarInvestasi, nData)
	case 6:
		tampilkanInvestasi(&daftarInvestasi, nData)
	case 7:
		isLoggedIn = false
		fmt.Println("🔓 Logout berhasil.\n")
	default:
		fmt.Println("🚫 Pilihan tidak valid.\n")
	}
}

func tambahInvestasi(data *[NMAX]Investasi, n *int, akun Akun) {
	if *n >= NMAX {
		fmt.Println("🚫 Data investasi penuh!")
		return
	}

	var pilihan int
	var jumlah float64
	var nama string
	var hargaSatuan float64
	var beli float64

	fmt.Println("\n💼 Pilih Jenis Investasi:")
	fmt.Println("1. Emas (Rp 2.000.000/gram)")
	fmt.Println("2. Dolar (Rp 16.500/lembar)")
	fmt.Println("3. Saham BBCA (Rp 900.000/lot)")
	fmt.Println("4. Saham BMRI (Rp 477.000/lot)")
	fmt.Println("5. Saham BBRI (Rp 384.000/lot)")
	fmt.Println("6. Saham TLKM (Rp 260.000/lot)")
	fmt.Print(">> Pilihan: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		nama = "Emas"
		hargaSatuan = 2000000
	case 2:
		nama = "Dollar"
		hargaSatuan = 16500
	case 3:
		nama = "BBCA"
		hargaSatuan = 900000
	case 4:
		nama = "BMRI"
		hargaSatuan = 477000
	case 5:
		nama = "BBRI"
		hargaSatuan = 384000
	case 6:
		nama = "TLKM"
		hargaSatuan = 260000
	default:
		fmt.Println("🚫 Pilihan tidak valid.")
		return
	}

	fmt.Print("Masukkan jumlah (gram/lembar/lot): ")
	fmt.Scanln(&jumlah)
	
	beli = jumlah * hargaSatuan
	if beli <= akun.Deposit {
		data[*n] = Investasi{
			ID:     idCounter,
			Nama:   nama,
			Jumlah: jumlah,
			Total:  beli,
		}
		idCounter++
		*n++
		fmt.Println("✅ Investasi berhasil ditambahkan!\n")
		akun.Deposit = akun.Deposit - beli
		fmt.Printf("Saldo anda saat ini: %.0f\n", akun.Deposit)
	} else {
		fmt.Println("🚫 Investasi gagal, saldo tidak cukup!\n")
	}
}

func tarikInvestasi(data *[NMAX]Investasi, n *int) {
	var id int
	tampilkanInvestasi(data, *n)
	fmt.Print("\nMasukkan ID investasi yang ingin ditarik: ")
	fmt.Scanln(&id)

	index := -1
	i := 0

	for i < *n {
		if data[i].ID == id {
			index = i
			i = *n // keluar dari loop tanpa break
		} else {
			i++
		}
	}

	if index != -1 {
		for j := index; j < *n-1; j++ {
			data[j] = data[j+1]
		}
		*n--
		fmt.Println("✅ Investasi berhasil ditarik!\n")
	} else {
		fmt.Println("🚫 ID tidak ditemukan.\n")
	}
}

func sortByHargaAsc(data *[NMAX]Investasi, n int) { // Ascending
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].Total < data[minIdx].Total {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	fmt.Println("✅ Investasi berhasil diurutkan berdasarkan total harga (terendah ke tertinggi).\n")
	tampilkanInvestasi(data, n)
}

func sortByHargaDsc(data *[NMAX]Investasi, n int) { // Descending
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].Total > data[maxIdx].Total {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
	fmt.Println("✅ Investasi berhasil diurutkan berdasarkan total harga (tertinggi ke terendah).\n")
	tampilkanInvestasi(data, n)
}

func sortByID(data *[NMAX]Investasi, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].ID < data[minIdx].ID {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	fmt.Println("✅ Investasi berhasil diurutkan berdasarkan ID (terkecil ke terbesar).\n")
	tampilkanInvestasi(data, n)
}

func tampilkanInvestasi(data *[NMAX]Investasi, n int) {
	tampilkanHeader("Daftar Investasi FinVest")
	if n == 0 {
		fmt.Println("Belum ada data investasi.")
		return
	}
	fmt.Printf("%-5s %-10s %-10s %-15s\n", "ID", "Nama", "Jumlah", "Total Harga")
	fmt.Println(strings.Repeat("-", 45))
	for i := 0; i < n; i++ {
		inv := data[i]
		fmt.Printf("%-5d %-10s %-10.2f Rp %-15.2f\n", inv.ID, inv.Nama, inv.Jumlah, inv.Total)
	}
	fmt.Println()
}

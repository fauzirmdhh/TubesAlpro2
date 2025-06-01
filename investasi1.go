package main

import (
	"fmt"
	"strings"
)

const NMAX int = 20

var daftarInvestasi [NMAX]Investasi

var saldo float64 = 0
var nData int
var isLoggedIn bool = false
var idCounter int = 100

type Akun struct {
	Username string
	Password string
}

var akun Akun

type Investasi struct {
	ID     int
	Nama   string
	Jumlah float64
	Total  float64
}

func main() {
	var pilihan int
	var isRunning bool

	isRunning = true

	for isRunning {
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
				isRunning = false
			default:
				fmt.Println("Pilihan tidak valid.\n")
			}
		} else {
			menuUtama()
		}
	}
}

func buatAkun() {
	fmt.Println("\nBuat Akun Baru")
	fmt.Print("Username: ")
	fmt.Scanln(&akun.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&akun.Password)
	fmt.Println("Akun berhasil dibuat!\n")
}

func login() {
	var LoginUsn, LoginPass string

	fmt.Println("\nLogin")
	fmt.Print("Username: ")
	fmt.Scanln(&LoginUsn)
	fmt.Print("Password: ")
	fmt.Scanln(&LoginPass)

	if LoginUsn == akun.Username && LoginPass == akun.Password {
		isLoggedIn = true
		fmt.Println("Login berhasil!\n")
	} else {
		fmt.Println("Username atau password salah\n")
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
	fmt.Printf("Saldo Anda: Rp %.2f\n", saldo)
	fmt.Println("1. Deposit Saldo")
	fmt.Println("2. Tambah Investasi")
	fmt.Println("3. Lihat Daftar Investasi")
	fmt.Println("4. Urutkan Berdasarkan Harga (Ascending)")
	fmt.Println("5. Urutkan Berdasarkan Harga (Descending)")
	fmt.Println("6. Urutkan Berdasarkan ID (Ascending)")
	fmt.Println("7. Cari Investasi Berdasarkan ID")
	fmt.Println("8. Jual/Tarik Investasi")
	fmt.Println("9. Logout")
	fmt.Print(">> Pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		depositSaldo()
	case 2:
		tambahInvestasi(&daftarInvestasi, &nData)
	case 3:
		tampilkanInvestasi(&daftarInvestasi, nData)
	case 4:
		sortByHargaAsc(&daftarInvestasi, nData)
	case 5:
		sortByHargaDsc(&daftarInvestasi, nData)
	case 6:
		sortByID(&daftarInvestasi, nData)
		fmt.Println("✅ Investasi berhasil diurutkan berdasarkan ID (terkecil ke terbesar).\n")
		tampilkanInvestasi(&daftarInvestasi, nData)
	case 7:
		cariInvestasiByID(&daftarInvestasi, nData)
		fmt.Println()
	case 8:
		tarikInvestasi(&daftarInvestasi, &nData)
	case 9:
		isLoggedIn = false
		fmt.Println("Logout berhasil.\n")
	default:
		fmt.Println("🚫 Pilihan tidak valid.\n")
	}
}

func tambahInvestasi(data *[NMAX]Investasi, n *int) {
	var pilihan int
	var jumlah float64
	var nama string
	var hargaSatuan float64
	var totalHarga float64

	if *n >= NMAX {
		fmt.Println("🚫 Data investasi penuh!")
	} else {
		fmt.Println("\n💼 Pilih Jenis Investasi:")
		fmt.Println("1. Emas (Rp 2.000.000/gram)")
		fmt.Println("2. Dolar (Rp 16.500/lembar)")
		fmt.Println("3. Saham BBCA (Rp 900.000/lot)")
		fmt.Println("4. Saham BMRI (Rp 477.000/lot)")
		fmt.Println("5. Saham BBRI (Rp 384.000/lot)")
		fmt.Println("6. Saham TLKM (Rp 260.000/lot)")
		fmt.Println("7. Saham ASII (Rp 471.000/lot)")
		fmt.Println("8. Saham GOTO (Rp 670.000/lot)")
		fmt.Println("9. Reksadana Pasar Uang (Rp 100.000/unit)")
		fmt.Println("10. Obligasi Ritel Indonesia (Rp 1.000.000/unit)")
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
		case 7:
			nama = "ASII"
			hargaSatuan = 471000
		case 8:
			nama = "GOTO"
			hargaSatuan = 670000
		case 9:
			nama = "Reksadana Pasar Uang"
			hargaSatuan = 100000
		case 10:
			nama = "ORI"
			hargaSatuan = 1000000
		default:
			fmt.Println("🚫 Pilihan tidak valid.")
		}

		fmt.Print("Masukkan jumlah (gram/lembar/lot/unit): ")
		fmt.Scanln(&jumlah)

		totalHarga = jumlah * hargaSatuan
		if totalHarga > saldo {
			fmt.Println("🚫 Saldo tidak cukup untuk melakukan investasi.\n")
		} else {
			data[*n] = Investasi{
				ID:     idCounter,
				Nama:   nama,
				Jumlah: jumlah,
				Total:  totalHarga,
			}

			saldo = saldo - totalHarga
			fmt.Println("✅ Investasi berhasil ditambahkan!\n")
		}

		idCounter++
		*n++
	}
}

func tarikInvestasi(data *[NMAX]Investasi, n *int) {
	var id, i, index, j int
	tampilkanInvestasi(data, *n)
	fmt.Print("\nMasukkan ID investasi yang ingin ditarik: ")
	fmt.Scanln(&id)

	index = -1
	i = 0

	for i < *n {
		if data[i].ID == id {
			index = i
			i = *n // keluar dari loop tanpa break
		} else {
			i++
		}
	}

	if index != -1 {
		for j = index; j < *n-1; j++ {
			data[j] = data[j+1]
		}
		*n--
		fmt.Println("✅ Investasi berhasil ditarik!\n")
	} else {
		fmt.Println("🚫 ID tidak ditemukan.\n")
	}

	tampilkanInvestasi(data, *n)
}

func sortByHargaAsc(data *[NMAX]Investasi, n int) { // Ascending
	// Selection Sort
	var i, minIdx, j int

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
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
	// Selection Sort
	var i, maxIdx, j int

	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
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
	// Selection Sort
	var i, minIdx, j int

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if data[j].ID < data[minIdx].ID {
				minIdx = j
			}
		}
		data[i] = data[minIdx]
		data[minIdx] = data[i]
	}
}

func tampilkanInvestasi(data *[NMAX]Investasi, n int) {
	var totalInvestasi float64

	tampilkanHeader("Daftar Investasi FinVest")

	if n == 0 {
		fmt.Println("Belum ada data investasi.")
	} else {
		fmt.Printf("%-5s %-10s %-10s %-15s\n", "ID", "Nama", "Jumlah", "Total Harga")
		fmt.Println(strings.Repeat("-", 45))
		for i := 0; i < n; i++ {
			inv := data[i]
			fmt.Printf("%-5d %-10s %-10.2f Rp %-15.2f\n", inv.ID, inv.Nama, inv.Jumlah, inv.Total)
			totalInvestasi += inv.Total
		}
		fmt.Printf("\nTotal Investasi Saat Ini: Rp %.2f\n\n", totalInvestasi)
	}
}

func depositSaldo() {
	var jumlah float64

	fmt.Print("Masukkan jumlah saldo yang ingin dideposit: Rp ")
	fmt.Scanln(&jumlah)
	if jumlah > 0 {
		saldo += jumlah
		fmt.Printf("✅ Deposit berhasil. Saldo saat ini: Rp %.2f\n\n", saldo)
	} else {
		fmt.Println("🚫 Jumlah tidak valid.\n")
	}
}

func cariInvestasiByID(data *[NMAX]Investasi, n int) {
	// Binary Search
	var id, low, high, mid int
	var found bool

	fmt.Print("Masukkan ID yang ingin dicari: ")
	fmt.Scanln(&id)

	if n > 0 {
		sortByID(data, n)
		low = 0
		high = n - 1

		for low <= high && !found {
			mid = (low + high) / 2
			if data[mid].ID == id {
				found = true
			} else if data[mid].ID < id {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if found {
			inv := data[mid]
			fmt.Println("✅ Investasi ditemukan:\n")
			fmt.Printf("%-5s %-10s %-10s %-15s\n", "ID", "Nama", "Jumlah", "Total Harga")
			fmt.Println(strings.Repeat("-", 45))
			fmt.Printf("%-5d %-10s %-10.2f Rp %-15.2f", inv.ID, inv.Nama, inv.Jumlah, inv.Total)
		} else {
			fmt.Println("🚫 Investasi dengan ID tersebut tidak ditemukan.\n")
		}
	} else {
		fmt.Println("🚫 Belum ada data investasi.\n")
	}
}

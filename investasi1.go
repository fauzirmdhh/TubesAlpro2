package main

import (
	"fmt"
	"strings"
)

const NMAX int = 20

var akun Akun
var daftarInvestasi [NMAX]Investasi
var nData = 0
var isLoggedIn bool = false
var idCounter int = 100

type Akun struct {
	Username string
	Password string
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
	} else {
		fmt.Println("🚫 Username atau password salah\n")
	}
}

func tampilkanHeader(judul string) {
	fmt.Println(strings.Repeat("=", len(judul)))
	fmt.Println(judul)
	fmt.Println(strings.Repeat("=", len(judul)))
}

func menuUtama() {
	var pilihan int
	tampilkanHeader("📊 FinVest - Menu Utama")
	fmt.Println("1. Tambah Investasi")
	fmt.Println("2. Jual/Tarik Investasi")
	fmt.Println("3. Urutkan Berdasarkan Harga (Ascending)")
	fmt.Println("4. Urutkan Berdasarkan ID (Ascending)")
	fmt.Println("5. Lihat Daftar Investasi")
	fmt.Println("6. Logout")
	fmt.Print(">> Pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahInvestasi()
	case 2:
		tarikInvestasi()
	case 3:
		sortByHarga()
	case 4:
		sortByID()
	case 5:
		tampilkanInvestasi()
	case 6:
		isLoggedIn = false
		fmt.Println("🔓 Logout berhasil.\n")
	default:
		fmt.Println("🚫 Pilihan tidak valid.")
	}
}

func sortByID() {
	for i := 0; i < nData-1; i++ {
		minIdx := i
		for j := i + 1; j < nData; j++ {
			if daftarInvestasi[j].ID < daftarInvestasi[minIdx].ID {
				minIdx = j
			}
		}
		daftarInvestasi[i], daftarInvestasi[minIdx] = daftarInvestasi[minIdx], daftarInvestasi[i]
	}
	fmt.Println("✅ Investasi berhasil diurutkan berdasarkan ID (terkecil ke terbesar).")
	tampilkanInvestasi()
}

func sortByHarga() {
	for i := 0; i < nData-1; i++ {
		minIdx := i
		for j := i + 1; j < nData; j++ {
			if daftarInvestasi[j].Total < daftarInvestasi[minIdx].Total {
				minIdx = j
			}
		}
		daftarInvestasi[i], daftarInvestasi[minIdx] = daftarInvestasi[minIdx], daftarInvestasi[i]
	}
	fmt.Println("✅ Investasi berhasil diurutkan berdasarkan total harga (terendah ke tertinggi).")
	tampilkanInvestasi() // langsung tampilkan hasil
}

func tambahInvestasi() {
	if nData >= NMAX {
		fmt.Println("🚫 Data investasi penuh!")
		return
	}

	var pilihan int
	var jumlah float64
	var nama string
	var hargaSatuan float64

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

	daftarInvestasi[nData] = Investasi{
		ID:     idCounter,
		Nama:   nama,
		Jumlah: jumlah,
		Total:  jumlah * hargaSatuan,
	}
	idCounter++
	nData++
	fmt.Println("✅ Investasi berhasil ditambahkan!")
}

func tarikInvestasi() {
	var id int
	tampilkanInvestasi()
	fmt.Print("\n Masukkan ID investasi yang ingin ditarik: ")
	fmt.Scanln(&id)

	// Inisialisasi index sebagai -1
	index := -1
	i := 0

	for i < nData {
		if daftarInvestasi[i].ID == id {
			index = i
			i = nData // langsung keluar dari loop dengan menyelesaikan kondisi
		} else {
			i++
		}
	}

	if index != -1 {
		// Geser elemen ke kiri untuk menghapus data pada index
		for j := index; j < nData-1; j++ {
			daftarInvestasi[j] = daftarInvestasi[j+1]
		}
		nData--
		fmt.Println("✅ Investasi berhasil ditarik!\n")
	} else {
		fmt.Println("🚫 ID tidak ditemukan.\n")
	}
}

func tampilkanInvestasi() {
	tampilkanHeader("\n📄 Daftar Investasi FinVest")
	if nData == 0 {
		fmt.Println("\n📭 Belum ada data investasi.")
		return
	}
	fmt.Printf("%-5s %-10s %-10s %-15s\n", "ID", "Nama", "Jumlah", "Total Harga")
	fmt.Println(strings.Repeat("-", 45))
	for i := 0; i < nData; i++ {
		inv := daftarInvestasi[i]
		fmt.Printf("%-5d %-10s %-10.2f Rp %-15.2f\n",
			inv.ID, inv.Nama, inv.Jumlah, inv.Total)
	}
}

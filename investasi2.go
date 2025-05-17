package main
import (

"fmt"
"os"
)

const NMAX int = 50
const MAXbuy int = 20

type investasi struct {
	Nama string
	Harga float64
	Jumlah float64
	Kategori string
}

type selfData struct {
	name string
	age int
	deposit float64
	gender string
}

type Menu struct {
	opsi int
}


type cari struct {
	Nama string
	Harga float64
	Jumlah float64
	Kategori string
}

var inv [NMAX]investasi
//var akanJual investasi

var pembelianSaya [MAXbuy]investasi
var jumlahPembelian int

var PenjualanSaya [MAXbuy]investasi
var jumlahPenjualan int

func main(){
	
	inv[0] = investasi{"RD Mandiri", 741576300.0, 2.5244, "Reksadana"}
	inv[1] = investasi{"Bitcoin", 676731800.0, 8.9326, "Crypto"}
	inv[2] = investasi{"RD Sucorinvest", 421979600.0, 0.3950, "Reksadana"}
	inv[3] = investasi{"Ethereum", 505404800.0, 0.3627, "Crypto"}
	inv[4] = investasi{"Ripple", 419577900.0, 4.5472, "Crypto"}
	inv[5] = investasi{"Emas Antam", 758831500.0, 1.6806, "Emas"}
	inv[6] = investasi{"BBRI", 277943600.0, 2.2316, "Saham"}
	inv[7] = investasi{"Emas Antam", 92836570.0, 1.0575, "Emas"}
	inv[8] = investasi{"Logam Mulia", 807147600.0, 7.3243, "Emas"}
	inv[9] = investasi{"RD Schroder", 78892320.0, 3.0025, "Reksadana"}
	inv[10] = investasi{"Emas UBS", 704601400.0, 0.5537, "Emas"}
	inv[11] = investasi{"Solana", 985223000.0, 8.5676, "Crypto"}
	inv[12] = investasi{"RD Schroder", 278045800.0, 6.3933, "Reksadana"}
	inv[13] = investasi{"Emas UBS", 370243900.0, 2.1741, "Emas"}
	inv[14] = investasi{"Emas Antam", 609170100.0, 1.7943, "Emas"}
	inv[15] = investasi{"Ethereum", 462314000.0, 2.7725, "Crypto"}
	inv[16] = investasi{"Solana", 842867600.0, 7.7824, "Crypto"}
	inv[17] = investasi{"Bitcoin", 805065300.0, 4.0715, "Crypto"}
	inv[18] = investasi{"RD BCA", 913144800.0, 5.7151, "Reksadana"}
	inv[19] = investasi{"Emas UBS", 655473100.0, 4.0168, "Emas"}
	inv[20] = investasi{"TLKM", 264953700.0, 2.5416, "Saham"}
	inv[21] = investasi{"Emas Pegadaian", 897833100.0, 4.0541, "Emas"}
	inv[22] = investasi{"Ethereum", 509575300.0, 1.0000, "Crypto"}
	inv[23] = investasi{"RD Mandiri", 152926000.0, 1.6838, "Reksadana"}
	inv[24] = investasi{"UNVR", 63621350.0, 3.8780, "Saham"}
	inv[25] = investasi{"UNVR", 251489000.0, 5.5769, "Saham"}
	inv[26] = investasi{"RD Mandiri", 681742200.0, 5.4160, "Reksadana"}
	inv[27] = investasi{"Logam Mulia", 111641000.0, 4.4042, "Emas"}
	inv[28] = investasi{"BBCA", 953820500.0, 8.7709, "Saham"}
	inv[29] = investasi{"Emas UBS", 507730900.0, 1.1535, "Emas"}
	inv[30] = investasi{"Emas UBS", 152924000.0, 7.6489, "Emas"}
	inv[31] = investasi{"RD Sucorinvest", 324223600.0, 0.2928, "Reksadana"}
	inv[32] = investasi{"Logam Mulia", 239528400.0, 2.4846, "Emas"}
	inv[33] = investasi{"RD Mandiri", 731934400.0, 8.1786, "Reksadana"}
	inv[34] = investasi{"Ethereum", 659784800.0, 9.4738, "Crypto"}
	inv[35] = investasi{"Solana", 527721000.0, 6.1053, "Crypto"}
	inv[36] = investasi{"Ripple", 755289700.0, 6.9297, "Crypto"}
	inv[37] = investasi{"Emas Pegadaian", 995149800.0, 6.5338, "Emas"}
	inv[38] = investasi{"UNVR", 451541000.0, 2.5543, "Saham"}
	inv[39] = investasi{"RD Panin", 21132160.0, 5.5838, "Reksadana"}
	inv[40] = investasi{"Bitcoin", 71085990.0, 6.3479, "Crypto"}
	inv[41] = investasi{"Bitcoin", 905429500.0, 8.6104, "Crypto"}
	inv[42] = investasi{"RD Sucorinvest", 238080800.0, 6.7229, "Reksadana"}
	inv[43] = investasi{"Ripple", 132398600.0, 9.3616, "Crypto"}
	inv[44] = investasi{"TLKM", 784641000.0, 8.0942, "Saham"}
	inv[45] = investasi{"Bitcoin", 97021120.0, 4.3674, "Crypto"}
	inv[46] = investasi{"ASII", 467078000.0, 7.3179, "Saham"}
	inv[47] = investasi{"RD Mandiri", 402681000.0, 3.4591, "Reksadana"}
	inv[48] = investasi{"RD BCA", 191670900.0, 5.4092, "Reksadana"}
	inv[49] = investasi{"Cardano", 183569700.0, 4.6800, "Crypto"}
	
	var d selfData
	var un, pw, unM, pwM string
	var k int = 0
	
	fmt.Println()
	fmt.Println("=====================================================")
	fmt.Println("Anda saat ini berada di aplikasi investasi sederhana, ")
	fmt.Println("ini adalah awal bagi anda, silahkan masukkan informasi")
	fmt.Println("daftar akun untuk memulai penjelajahan anda...")
	fmt.Println("=====================================================")
	fmt.Printf(">> Daftar akun \n")
	fmt.Printf("username%5s ", ":")
	fmt.Scan(&un)
	fmt.Printf("password%5s ", ":")
	fmt.Scan(&pw)
	fmt.Println()
	fmt.Printf("Pendaftaran berhasil \n")
	fmt.Println()
		fmt.Println("=====================================================")
		fmt.Printf(">> Silahkan login \n")
		fmt.Printf("username%5s ", ":")
		fmt.Scan(&unM)
		fmt.Printf("password%5s ", ":")
		fmt.Scan(&pwM)
		fmt.Println()
	for k < 3 {
		if unM == un && pwM == pw { 
			fmt.Printf("Login berhasil \n")
			fmt.Println("+---------------------------------------------------+")
			fmt.Printf("|%52s\n", "|")
			fmt.Printf("%-14sMasukkan Data Diri Anda%16s\n", "|", "|")
			fmt.Printf("|%52s\n", "|")
			fmt.Println("+---------------------------------------------------+")
			fmt.Println("+---------------------------------------------------+")
			fmt.Println("Lengkapi data diri anda sebelum ke halaman utama...")
			fmt.Println()
			dataPengguna(&d)
		} else {
			fmt.Printf("Login gagal \n")
			fmt.Printf("Masukkan informasi login yang benar \n")
			fmt.Println("=====================================================")
			fmt.Printf(">> Silahkan login \n")
			fmt.Printf("username%5s ", ":")
			fmt.Scan(&unM)
			fmt.Printf("password%5s ", ":")
			fmt.Scan(&pwM)
			fmt.Println()
			k++
		}
	}
	if k == 3 {
			fmt.Println()
			fmt.Printf("Gagal dilakukan login, silahkan mulai dari awal \n")
	}
}


func menu(D *selfData, done *[1]investasi){ //======================================================== (3)
	var M Menu
	var key cari
	var id string
	fmt.Printf("Apa yang akan anda lakukan sekarang 1/2/3/4 :\n")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Tampilkan semua investasi")
	fmt.Println("2. Buy asset")
	fmt.Println("3. Sell asset")
	fmt.Println("4. Keluar")
	fmt.Println()
	fmt.Printf("Jawaban anda%4s ", ":")
	fmt.Scan(&M.opsi)
	pemanggilan(M, &key, &id, D, done)
}

func dataPengguna(D *selfData){ //===================================== (1)
	fmt.Printf("Nama%7s ", ":")
	fmt.Scan(&D.name)
	fmt.Printf("Usia%7s ", ":")
	fmt.Scan(&D.age)
	fmt.Printf("Gender%5s ", ":")
	fmt.Scan(&D.gender)
	fmt.Printf("Negara%5s %s\n", ":", "Indonesia")
	fmt.Println("-----------------------------------------------------")
	fmt.Println()
	deposit(D)
}

func deposit(D *selfData){ //=========================================== (2)
	var done [1]investasi
	fmt.Printf("Deposit%4s ", ":")
	fmt.Scan(&D.deposit)
	fmt.Println()
	fmt.Println("-----------------------------------------------------")
	menu(D, &done)
	
}

func pemanggilan(M Menu, target *cari, kode *string, D *selfData, done *[1]investasi){ //=============== (4)
	var pilih int
	var decide string
	switch M.opsi{
		case 1:
			tampilkanInv(inv)
		case 2:
			fmt.Println("=====================================================")
			fmt.Println("Temukan Aset :")
			fmt.Println("-----------------------------------------------------")
			fmt.Printf("Dibawah ini anda akan diminta untuk memasukkan kode\npencarian berdasarkan [N : Nama], [H : Harga],\n[J : Jumlah], [K : Kategori]\n\nPastikan jika anda memilih kode [N], mulailah\npencarian dengan penggunaan kapital diawal kata kunci\njuga kategori yang tersedia pada aplikasi ini adalah:\n -Crypto\n -Reksadana\n -Saham\n -Emas\n\n --silahkan lakukan pencarian anda--\n")
			fmt.Println("=====================================================")
			fmt.Printf("Masukkan kode pencarian [N, H, J, K] :")
			fmt.Scan(kode)
			if *kode == "N"{
				fmt.Printf("Nama Aset%9s ", ":")
				fmt.Scan(&target.Nama)
			} else if *kode == "H" {
				fmt.Printf("Harga maksimal%4s ", ":")
				fmt.Scan(&target.Harga)
			} else if *kode == "J" {
				fmt.Printf("Jumlah maksimal%3s ", ":")
				fmt.Scan(&target.Jumlah)
			} else {
				fmt.Printf("Kategori%10s ", ":")
				fmt.Scan(&target.Kategori)
			}
			Buy(inv, *target, *kode)
			printPencarian(inv, &pilih, done, target, kode, D)
		case 3:
			investasiSaya(D)
			pilihSell(done, D)
			fmt.Println(Sell(inv, done))
			fmt.Println("-----------------------------------------------------")
		case 4:
			endOfProgram(&decide, D)
	}
}

//buy program================================================================

func printPencarian(inv [NMAX]investasi, pilih *int, done *[1]investasi, target *cari, kode *string, D *selfData) { //option I
	var i, j, miN, tukar, y int
	var temukan [20]investasi
	var panjang int
	var totalBuyVal *investasi
	temukan, panjang = Buy(inv, *target, *kode)
	if panjang > 0 {
		fmt.Println()
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Data yang berhasil ditemukan :")
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Data telah diurutkan mulai dari harga yang terendah:")
		fmt.Println()
		for i = 0; i < panjang; i++{
			
			//selection sort
			miN = i
			for y = i+1; y < panjang; y++ {
				tukar = y
				if temukan[tukar].Harga < temukan[miN].Harga {
					miN = tukar
				}
			}
			temukan[i], temukan[miN] = temukan[miN], temukan[i]
			fmt.Printf("%d. \n", i+1)
			fmt.Printf("%4s Aset%10s %s\n", "|", ":", temukan[i].Nama)
			fmt.Printf("%4s Kategori%6s %s\n", "|", ":", temukan[i].Kategori)
			fmt.Printf("%4s Jumlah%8s %.4f\n", "|", ":", temukan[i].Jumlah)
			fmt.Printf("%4s Total Harga%3s %.2f\n", "|", ":", temukan[i].Harga)
			fmt.Println("-----------------------------------------------------")
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Mana yang akan anda beli 1/2/3/...?")
		fmt.Printf("Urutan ke%4s ", ":")
		fmt.Scan(pilih)
		fmt.Println()		
		for j = 0; j < panjang; j++ {
			if temukan[j] == temukan[*pilih-1] {
				(*done)[0] = temukan[j]
				fmt.Printf("%4s Aset%10s %s\n", "|", ":", (*done)[0].Nama)
				fmt.Printf("%4s Kategori%6s %s\n", "|", ":",(*done)[0].Kategori)
				fmt.Printf("%4s Jumlah%8s %.4f\n", "|", ":", (*done)[0].Jumlah)
				fmt.Printf("%4s Total Harga%3s %.2f\n", "|", ":", (*done)[0].Harga)
			}
		}
		
		*done = [1]investasi{temukan[*pilih-1]}
		fmt.Println()
		totalBuyVal = totalBuy(inv, done, D)
		if totalBuyVal == nil {
			fmt.Println("Pembelian gagal:")
			fmt.Println("lakukan pembelian dengan nominal dibawah deposit anda")
			fmt.Println("-----------------------------------------------------")
			menu(D, done)
		}
		fmt.Printf("%4s Aset%10s %s\n", "|", ":", totalBuyVal.Nama)
		fmt.Printf("%4s Kategori%6s %s\n", "|", ":",totalBuyVal.Kategori)
		fmt.Printf("%4s Jumlah%8s %.4f\n", "|", ":", totalBuyVal.Jumlah)
		fmt.Printf("%4s Total Harga%3s %.2f\n", "|", ":", totalBuyVal.Harga)
		fmt.Println("-----------------------------------------------------")
		menu(D, done)
		return
	} else {
		fmt.Println("+---------------------------------------------------+")
		fmt.Printf("%-22sNot Found%21s\n", "|", "|")
		fmt.Println("+---------------------------------------------------+")
		fmt.Println("Buat pilihan anda kembali:")
		menu(D, done)
	}
}

//linear saerch
func Buy(inv [NMAX]investasi, target cari, kode string) ([20]investasi, int) { //proses I.I
	var i int
	var idx int = 0
	var klasifikasi [20]investasi
	if kode == "N" {
		for i = 0; i < NMAX; i++ {
			if inv[i].Nama == target.Nama {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	} else if kode == "H" {
		for i = 0; i < NMAX; i++ {
			if inv[i].Harga <= target.Harga && i < 20 {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	} else if kode == "J" {
		for i = 0; i < NMAX; i++ {
			if inv[i].Jumlah <= target.Jumlah && i < 20 {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	} else if kode == "K" {
		for i = 0; i < NMAX; i++ {
			if inv[i].Kategori == target.Kategori {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	}
	
	return klasifikasi, idx //mengembalikan daftar array berdasarkan target dan banyak daftar
}

//mengubah nilai di array investasi secara non-permanen
func totalBuy(inv [NMAX]investasi, done *[1]investasi, D *selfData) *investasi{ //proses I.II
	var akanBeli investasi
	var totalBeli float64
	var i int = 0
	var j int
	if done[i].Nama != "" {
		
		//harga perunit Jumlah
		akanBeli.Harga = done[i].Harga / done[i].Jumlah
		fmt.Printf("Sejumlah%9s ", ":")
		fmt.Scan(&akanBeli.Jumlah)
		if akanBeli.Jumlah <= done[i].Jumlah {
			totalBeli = akanBeli.Harga * akanBeli.Jumlah
			if totalBeli > D.deposit {
				fmt.Println("Saldo tidak cukup untuk membeli aset ini.")
					return nil
			}
			fmt.Printf("Totalnya%9s %.2f\n", ":", totalBeli)
			
			//memastikan jumlah pembelian hanya diproses jika dibawah MAXbuy
			if jumlahPembelian < MAXbuy {
				D.deposit -= totalBeli
				fmt.Printf("Sisa saldo%7s %.2f\n", ":", D.deposit)
				akanBeli.Harga = totalBeli
				akanBeli.Nama = done[i].Nama
				akanBeli.Kategori = done[i].Kategori
				
				//jika aset sudah ada dengan nama dan kategori yang sama, bisa digabung saja tanpa pemborosan list baru
				for j = 0; j < jumlahPembelian; j++ {
					if pembelianSaya[j].Nama == akanBeli.Nama && pembelianSaya[j].Kategori == akanBeli.Kategori {
						pembelianSaya[j].Jumlah += akanBeli.Jumlah
						pembelianSaya[j].Harga += akanBeli.Harga
						return &pembelianSaya[j]
					}
				}
				
				//jika tidak ada langsung return nilainya saja
				pembelianSaya[jumlahPembelian] = akanBeli
				fmt.Println()
				fmt.Printf("Investasi saya%1s\n", ":")
				jumlahPembelian++
				return &akanBeli
			} 
		}
	} else {
		return nil
	}
	return nil
}

//sell program===============================================================

func investasiSaya(D *selfData){
	var done [1]investasi	
	var i int
	var aset investasi
	
	if jumlahPembelian == 0 {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Anda belum memiliki aset investasi")
		fmt.Println("Silahkan lakukan pembelian dulu")
		fmt.Println("-----------------------------------------------------")
		menu(D, &done)
		return //memastikan tidak dilakukan pengecekan lagi di pengkondisian dibawah
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("%14s\n", "Daftar Aset investasi Saya")
		fmt.Println("-----------------------------------------------------")
		for i = 0; i < jumlahPembelian; i++ {
			aset = pembelianSaya[i]
			fmt.Println("+---------------------------------------------------+")
			fmt.Printf("%d.%2s Aset%10s %s\n", i+1, "|", ":", aset.Nama)
			fmt.Printf("%4s Kategori%6s %s\n", "|", ":", aset.Kategori)
			fmt.Printf("%4s Jumlah%8s %.4f\n", "|", ":", aset.Jumlah)
			fmt.Printf("%4s Total Harga%3s %.2f\n", "|", ":", aset.Harga)
		}
		fmt.Println("-----------------------------------------------------")
	}
	
}

func pilihSell(done *[1]investasi, D *selfData){ 
	var j int
	var temukan [20]investasi
	var SellVal *investasi
	var panjang int = 20
	var pilih int
	PenjualanSaya = pembelianSaya
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Mana yang akan anda jual 1/2/3/...?")
		fmt.Printf("Urutan ke%4s ", ":")
		fmt.Scan(&pilih)
		fmt.Println()
		temukan[0] = PenjualanSaya[pilih-1]
		if pilih >= 0 && pilih < panjang {
			for j = 0; j < panjang; j++ {
				if PenjualanSaya[j].Nama == temukan[0].Nama &&
					PenjualanSaya[j].Kategori == temukan[0].Kategori &&
					PenjualanSaya[j].Harga == temukan[0].Harga &&
					PenjualanSaya[j].Jumlah == temukan[0].Jumlah {
					(*done)[0] = PenjualanSaya[j]
					fmt.Printf("%4s Aset%10s %s\n", "|", ":", (*done)[0].Nama)
					fmt.Printf("%4s Kategori%6s %s\n", "|", ":",(*done)[0].Kategori)
					fmt.Printf("%4s Jumlah%8s %.4f\n", "|", ":", (*done)[0].Jumlah)
					fmt.Printf("%4s Total Harga%3s %.2f\n", "|", ":", (*done)[0].Harga)
				}
			}
		}
		//mengisi done dengan data 'PenjualanSaya'
		*done = [1]investasi{PenjualanSaya[pilih-1]}
		SellVal = Sell(inv, done)
		if SellVal == nil {
			fmt.Println("Penjualan gagal:")
			fmt.Println("tampaknya anda tidak punya aset yang dapat dijual")
			fmt.Println("-----------------------------------------------------")
			menu(D, done)
			return
		}
		fmt.Printf("%4s Nama%10s %s\n", "|", ":", SellVal.Nama)
		fmt.Printf("%4s Kategori%6s %s\n", "|", ":",SellVal.Kategori)
		fmt.Printf("%4s Sejumlah%6s %.4f\n", "|", ":", SellVal.Jumlah)
		fmt.Printf("%4s Harga%9s %.2f\n", "|", ":", SellVal.Harga)
		fmt.Println("-----------------------------------------------------")
		menu(D, done)
}

func Sell(inv [NMAX]investasi, done *[1]investasi)*investasi{
	var akanJual investasi
	var totalJual float64
	var i int = 0
	var j int = 0
	var k int 
	var sudahDiproses bool = false
	if done[i].Nama != "" {
		akanJual.Harga = done[i].Harga / done[i].Jumlah
		fmt.Println()
		fmt.Printf("Sejumlah%9s ", ":")
		fmt.Scan(&akanJual.Jumlah)
		if akanJual.Jumlah <= done[i].Jumlah {
			totalJual = akanJual.Harga * akanJual.Jumlah
			fmt.Printf("Totalnya%9s %.2f\n", ":", totalJual)
			
			if jumlahPenjualan < MAXbuy {
				for !sudahDiproses && j < jumlahPembelian {
					if pembelianSaya[j].Nama == done[i].Nama && pembelianSaya[j].Kategori == done[i].Kategori {
						pembelianSaya[j].Jumlah -= akanJual.Jumlah
						if pembelianSaya[j].Jumlah <= 0 {
							for k = j; k < jumlahPembelian-1; k++ {
								pembelianSaya[k] = pembelianSaya[k+1]
							}
							jumlahPembelian--
						} else {
							j++
						}
						sudahDiproses = true
					} else {
						j++
					}
				}
				akanJual.Harga = totalJual
				akanJual.Nama = done[i].Nama
				akanJual.Kategori = done[i].Kategori
				PenjualanSaya[jumlahPenjualan] = akanJual
				fmt.Println()
				fmt.Printf("Aset yang saya jual%1s\n", ":")
				jumlahPenjualan++
				return &akanJual
			} 
		}
	} else {
		return nil
	}
	return nil
}

func tampilkanInv(inv [NMAX]investasi){//, akanJual *investasi){
	var n, i int
	var D selfData
	var done [1]investasi
	fmt.Println("Sistem memiliki 50 baris data")
	fmt.Println("Berapa baris data yang anda ingin tampilkan?")
	fmt.Printf("Jawaban anda : ")
	fmt.Scan(&n)
	if n <= NMAX {
		fmt.Println("+===================================================+")
		fmt.Printf("%-12sDAFTAR INFESTASI YANG TERSEDIA%12s\n", "|", "|")
		fmt.Println("+---------------------------------------------------+")
		for i = 0; i < n; i++ {
			fmt.Printf("%d. \n", i+1)
			fmt.Printf("%4s Nama%10s %s\n", "|", ":", inv[i].Nama)
			fmt.Printf("%4s Kategori%6s %s\n", "|", ":", inv[i].Kategori)
			fmt.Printf("%4s Jumlah%8s %.4f\n", "|", ":", inv[i].Jumlah)
			fmt.Printf("%4s Harga%9s %.2f\n", "|", ":", inv[i].Harga)
			fmt.Println("-----------------------------------------------------")
		}
		menu(&D, &done)
	}
	
}


func endOfProgram(putuskan *string, D *selfData){
	var done [1]investasi
	fmt.Printf("%13s\n", "Yakin Ingin Keluar?")
	fmt.Printf("+-------+   +-------+\n")
	fmt.Printf("|  %-5s|   | %-5s |\n", "Ya", "Tidak")
	fmt.Printf("+-------+   +-------+\n")
	fmt.Println()
	fmt.Printf("Pilih mana? :")
	fmt.Scan(putuskan)
	switch *putuskan {
	
	case "Ya":
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("%-17s--Program Berakhir--%16s\n", "|", "|")
		fmt.Println("-----------------------------------------------------")
		os.Exit(0)
	case "Tidak":
		fmt.Println("-----------------------------------------------------")
		menu(D, &done)
	}
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<(PERLU DI SELESAIKAN)<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

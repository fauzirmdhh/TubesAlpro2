package main
import "fmt"

const NMAX int = 50
const MAXbuy int = 10

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

var pembelianSaya [MAXbuy]investasi
var jumlahPembelian int

func main(){
	
	var inv [NMAX]investasi
	
	inv[0] = investasi{"RD Mandiri", 741576300, 2.5244, "Reksadana"}
	inv[1] = investasi{"Bitcoin", 676731800, 8.9326, "Crypto"}
	inv[2] = investasi{"RD Sucorinvest", 421979600, 0.3950, "Reksadana"}
	inv[3] = investasi{"Ethereum", 505404800, 0.3627, "Crypto"}
	inv[4] = investasi{"Ripple", 419577900, 4.5472, "Crypto"}
	inv[5] = investasi{"Emas Antam", 758831500, 1.6806, "Emas"}
	inv[6] = investasi{"BBRI", 277943600, 2.2316, "Saham"}
	inv[7] = investasi{"Emas Antam", 92836570, 1.0575, "Emas"}
	inv[8] = investasi{"Logam Mulia", 807147600, 7.3243, "Emas"}
	inv[9] = investasi{"RD Schroder", 78892320, 3.0025, "Reksadana"}
	inv[10] = investasi{"Emas UBS", 704601400, 0.5537, "Emas"}
	inv[11] = investasi{"Solana", 985223000, 8.5676, "Crypto"}
	inv[12] = investasi{"RD Schroder", 278045800, 6.3933, "Reksadana"}
	inv[13] = investasi{"Emas UBS", 370243900, 2.1741, "Emas"}
	inv[14] = investasi{"Emas Antam", 609170100, 1.7943, "Emas"}
	inv[15] = investasi{"Ethereum", 462314000, 2.7725, "Crypto"}
	inv[16] = investasi{"Solana", 842867600, 7.7824, "Crypto"}
	inv[17] = investasi{"Bitcoin", 805065300, 4.0715, "Crypto"}
	inv[18] = investasi{"RD BCA", 913144800, 5.7151, "Reksadana"}
	inv[19] = investasi{"Emas UBS", 655473100, 4.0168, "Emas"}
	inv[20] = investasi{"TLKM", 264953700, 2.5416, "Saham"}
	inv[21] = investasi{"Emas Pegadaian", 897833100, 4.0541, "Emas"}
	inv[22] = investasi{"Ethereum", 509575300, 1.0000, "Crypto"}
	inv[23] = investasi{"RD Mandiri", 152926000, 1.6838, "Reksadana"}
	inv[24] = investasi{"UNVR", 63621350, 3.8780, "Saham"}
	inv[25] = investasi{"UNVR", 251489000, 5.5769, "Saham"}
	inv[26] = investasi{"RD Mandiri", 681742200, 5.4160, "Reksadana"}
	inv[27] = investasi{"Logam Mulia", 111641000, 4.4042, "Emas"}
	inv[28] = investasi{"BBCA", 953820500, 8.7709, "Saham"}
	inv[29] = investasi{"Emas UBS", 507730900, 1.1535, "Emas"}
	inv[30] = investasi{"Emas UBS", 152924000, 7.6489, "Emas"}
	inv[31] = investasi{"RD Sucorinvest", 324223600, 0.2928, "Reksadana"}
	inv[32] = investasi{"Logam Mulia", 239528400, 2.4846, "Emas"}
	inv[33] = investasi{"RD Mandiri", 731934400, 8.1786, "Reksadana"}
	inv[34] = investasi{"Ethereum", 659784800, 9.4738, "Crypto"}
	inv[35] = investasi{"Solana", 527721000, 6.1053, "Crypto"}
	inv[36] = investasi{"Ripple", 755289700, 6.9297, "Crypto"}
	inv[37] = investasi{"Emas Pegadaian", 995149800, 6.5338, "Emas"}
	inv[38] = investasi{"UNVR", 451541000, 2.5543, "Saham"}
	inv[39] = investasi{"RD Panin", 21132160, 5.5838, "Reksadana"}
	inv[40] = investasi{"Bitcoin", 71085990, 6.3479, "Crypto"}
	inv[41] = investasi{"Bitcoin", 905429500, 8.6104, "Crypto"}
	inv[42] = investasi{"RD Sucorinvest", 238080800, 6.7229, "Reksadana"}
	inv[43] = investasi{"Ripple", 132398600, 9.3616, "Crypto"}
	inv[44] = investasi{"TLKM", 784641000, 8.0942, "Saham"}
	inv[45] = investasi{"Bitcoin", 97021120, 4.3674, "Crypto"}
	inv[46] = investasi{"ASII", 467078000, 7.3179, "Saham"}
	inv[47] = investasi{"RD Mandiri", 402681000, 3.4591, "Reksadana"}
	inv[48] = investasi{"RD BCA", 191670900, 5.4092, "Reksadana"}
	inv[49] = investasi{"Cardano", 183569700, 4.6800, "Crypto"}
	inv[50] = investasi{"RD Schroder", 808041442, 8.5741, "Reksadana"}
	
	var d selfData
	//var m Menu
	//var key cari
	//var id string
	
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("|%52s\n", "|")
	fmt.Printf("%-14sMasukkan Data Diri Anda%14s\n", "|", "|")
	fmt.Printf("|%52s\n", "|")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")
	fmt.Println()
	dataPengguna(&d)
}

func menu(){ //======================================================== (3)
	var M Menu
	var key cari
	var id string
	fmt.Printf("Apa yang akan anda lakukan sekarang :\n")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Buy asset")
	fmt.Println("2. Sell asset")
	fmt.Println("3. Tampilkan semua investasi")
	fmt.Println("4. Keluar")
	fmt.Scan(&M.opsi)
	pemanggilan(M, &key, &id)
}

func dataPengguna(D *selfData){ //===================================== (1)
	fmt.Printf("Nama%7s", ":")
	fmt.Scan(&D.name)
	fmt.Printf("Usia%7s", ":")
	fmt.Scan(&D.age)
	fmt.Println("-----------------------------------------------------")
	fmt.Println()
	deposit(&D)
}

func pemanggilan(M Menu, target *cari, kode *string){ //=============== (4)
	var pilih int
	var done [1]investasi
	switch M.opsi{
		case 1:
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Temukan Aset :")
			fmt.Println("-----------------------------------------------------")
			fmt.Printf("Masukkan kode pencarian [N, H, J, K] :")
			fmt.Scan(&kode)
			if kode == "N"{
				fmt.Scan(&target.Nama)
			} else if kode == "H" {
				fmt.Scan(&target.Harga)
			} else if kode == "J" {
				fmt.Scan(&target.Jumlah)
			} else {
				fmt.Scan(&target.Kategori)
			}
			printPencarian(inv, &pilih, &done, *target, *kode)
		case 2:
			investasiSaya()
			Sell()
		case 3:
			tampilkanInv()
		case 4:
			endOfProgram()
	}
}

func deposit(D *selfData){ //=========================================== (2)
	fmt.Printf("Deposit%4s", ":")
	fmt.Scan(&D.deposit)
	menu()
}

//buy program================================================================

func printPencarian(inv [NMAX]investasi, pilih *int, done *[1]investasi, target cari, kode string) { //option I
	var i, j int
	var target cari
	var kode string
	var temukan [10]investasi
	var panjang int
	temukan, panjang = Buy(inv, target, kode)
	if panjang > 0 {
		fmt.Println()
		fmt.Println("Data yang berhasil ditemukan :")
		fmt.Println("-----------------------------------------------------")
		for i = 0; i < panjang; i++{
			fmt.Println(temukan[i])
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Mana yang akan anda beli?")
		fmt.Println()
		fmt.Scan(&pilih) 
		for j = 0; j < panjang; j++ {
			if temukan[j].Nama == temukan[pilih].Nama {
				*done[0] = temukan[j]
				fmt.Println(done[0])
			}
		}
		fmt.Println()
		fmt.Println(totalBuy(inv, done,))
		fmt.Println("-----------------------------------------------------")
		menu()
	} else {
		fmt.Println("Not Found")
	}
}

func Buy(inv [NMAX]investasi, target cari, kode string) ([10]investasi, int) { //proses I.I
	var i int
	var idx int = 0
	var klasifikasi [10]investasi
	if kode == "N" {
		for i = 0; i <= NMAX; i++ {
			if inv[i].Nama == target.Nama {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	} else if kode == "H" {
		for i = 0; i <= NMAX; i++ {
			if inv[i].Harga == target.Harga {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	} else if kode == "J" {
		for i = 0; i <= NMAX; i++ {
			if inv[i].Jumlah == target.Jumlah {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	} else if kode == "K" {
		for i = 0; i <= NMAX; i++ {
			if inv[i].Kategori == target.Kategori {
				klasifikasi[idx] = inv[i]
				idx++
			}
		}
	}
	
	return klasifikasi, idx //mengembalikan daftar array berdasarkan target dan banyak daftar
}

//mengubah nilai di array investasi secara non-permanen
func totalBuy(inv [NMAX]investasi, done [1]investasi) *investasi{ //proses I.II
	var akanBeli investasi
	var D selfData
	var totalBeli float64
	var i int = 0
	if done[i].Nama != "" {
		akanBeli.Harga = done[i].Harga / done[i].Jumlah
		fmt.Printf("Sejumlah%9s ", ":")
		fmt.Scan(&akanBeli.Jumlah)
		if akanBeli.Jumlah <= done[i].Jumlah {
			totalBeli = akanBeli.Harga * akanBeli.Jumlah
			fmt.Printf("Totalnya%9s %d\n", ":", totalBeli)
			
			if jumlahPembelian < MAXbuy {
				D.deposit = D.deposit - totalBeli
				fmt.Printf("Sisa saldo%5s %d\n", ":", D.deposit)
				akanBeli.Harga = totalBeli
				akanBeli.Nama = done[i].Nama
				akanBeli.Kategori = done[i].Kategori
				pembelianSaya[jumlahPembelian] = akanBeli
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

func investasiSaya(){ // optionn II
	var i int
	var aset investasi
	
	if jumlahPembelian == 0 {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Anda belum memiliki aset investasi")
		fmt.Println("Silahkan lakukan pembelian dulu")
		fmt.Println("-----------------------------------------------------")
		menu()
		return
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("%14s\n", "Daftar Aset investasi Saya")
		fmt.Println("-----------------------------------------------------")
		for i = 0; i < jumlahPembelian; i++ {
			aset = pembelianSaya[i]
			fmt.Println("+---------------------------------------------------+")
			fmt.Printf("%d.%2s Aset%10s %s\n", i+1, "|", ":", aset.Nama)
			fmt.Printf("%4s Kategori%6s %s\n", "|", ":", aset.Kategori)
			fmt.Printf("%4s Jumlah%8s %s\n", "|", ":", aset.Jumlah)
			fmt.Printf("%4s Total Harga%3s %s\n", "|", ":", aset.Harga)
		}
		fmt.Println("-----------------------------------------------------")
	}
	
}

func pilihSell(){
	
}

func Sell(){
	
}

func tampilkanInv(){
	
}

func endOfProgram(){
	
}

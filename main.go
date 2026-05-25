package main

import "fmt"

// --- 1. DEKLARASI KONSTANTA DAN STRUCT ---
const NMAX = 100

// Struct untuk Data Master Barang
type Barang struct {
	Kode  string
	Nama  string
	Harga int
	Stok  int
}

// Struct untuk Log Transaksi
type Transaksi struct {
	KodeBarang string
	Jenis      string // "Masuk" atau "Keluar"
	Jumlah     int
}

// Alias untuk Array
type TabBarang [NMAX]Barang
type TabTransaksi [NMAX]Transaksi


// --- 2. PROSEDUR TAMBAH BARANG (Draf Awal) ---
// Prosedur ini menggunakan pointer (*) agar data tersimpan ke array asli
func tambahBarang(data *TabBarang, n *int) {
	fmt.Println("\n--- TAMBAH DATA BARANG ---")
	// Cek apakah gudang (array) sudah penuh
	if *n >= NMAX {
		fmt.Println("Kapasitas gudang penuh! Tidak bisa menambah barang baru.")
		return
	}

	// Meminta input data dari user
	fmt.Print("Masukkan Kode Barang  : ")
	fmt.Scan(&data[*n].Kode)
	fmt.Print("Masukkan Nama Barang  : ")
	fmt.Scan(&data[*n].Nama)
	fmt.Print("Masukkan Harga Barang : ")
	fmt.Scan(&data[*n].Harga)
	fmt.Print("Masukkan Stok Awal    : ")
	fmt.Scan(&data[*n].Stok)

	*n++ // Menambah counter jumlah barang
	fmt.Println("Data barang berhasil ditambahkan!")
}


// --- 3. FUNGSI UTAMA (MAIN) ---
func main() {
	var gudang TabBarang
	var nBarang int = 0
	var pilihan int
	var jalan bool = true

	for jalan {
		fmt.Println("\n=========================================")
		fmt.Println("   SISTEM INVENTARIS GUDANG (GUDANGIN)   ")
		fmt.Println("=========================================")
		fmt.Println("1. Tambah Data Barang")
		fmt.Println("2. Tampilkan Data Barang")
		fmt.Println("0. Keluar")
		fmt.Println("=========================================")
		fmt.Print("Pilih menu (0-2): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// Memanggil prosedur tambahBarang
			tambahBarang(&gudang, &nBarang)
		case 2:
			// Menampilkan data sementara (langsung di dalam main untuk test)
			fmt.Println("\n--- DAFTAR BARANG ---")
			if nBarang == 0 {
				fmt.Println("Gudang masih kosong.")
			} else {
				for i := 0; i < nBarang; i++ {
					fmt.Printf("%d. [%s] %s - Harga: %d | Stok: %d\n", i+1, gudang[i].Kode, gudang[i].Nama, gudang[i].Harga, gudang[i].Stok)
				}
			}
		case 0:
			fmt.Println("Keluar dari program...")
			jalan = false
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
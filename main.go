package main

import "fmt"

const NMAX = 100

type Barang struct {
	Kode string
	Nama string 
	Harga int
	Stok int
}

type Transaksi struct {
	Kode string 
	Jenis string 
	Jumlah int
}

type dataBarang [NMAX]Barang
type dataTransaksi [NMAX]Transaksi 

func tambahBarang(data *dataBarang, n * int) {
	fmt.Println("\n== Tambah Data Barang ==")

	if *n >= NMAX {
		fmt.Println("Tidak bisa menambah barang baru lagi.")
		return 
	}

	fmt.Print("Masukkan Kode Barang : ") 
	fmt.Scan(&data[*n].Kode)

	fmt.Print("Masukkan Nama Barang : ") 
	fmt.Scan(&data[*n].Nama)

	fmt.Print("Masukkan Harga Barang: ") 
	fmt.Scan(&data[*n].Harga)

	fmt.Print("Masukkan Stok Awal   : ") 
	fmt.Scan(&data[*n].Stok)

	(*n)++
	
	fmt.Println("Data barang berhasil ditambahkan") 

}

func tampilkanBarang(data dataBarang, n int) { 
	fmt.Println("\n== Daftar Barang di Gudang")

	if n == 0 {
		fmt.Println("Gudang saat ini masih kosong")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. [%s] %s\n", i+1, data[i].Kode, data[i].Nama)
		fmt.Printf(" Harga: Rp%d | Stok: %d\n", data[i].Harga, data[i].Stok)
	}
}

func ubahdatabarang(data *dataBarang, n int ) {
	var kode string 
	var ketemu bool = false 

	fmt.Println("\n== Ubah Data Barang ==")

	if n == 0 {
		fmt.Println("Gudang Kosong")
		return
	}

	fmt.Print("Masukan Kode Barang yang akan diubah; ")
	fmt.Scan(&kode)

	for i := 0; i < n; i++ {
		if data[i].Kode == kode {
			ketemu = true 

			fmt.Printf("Barang ditemukan: %s\n", data[i].Nama)

			fmt.Print("Nama Baru  : ") 
			fmt.Scan(&data[i].Nama)

			fmt.Print("Harga Baru : ") 
			fmt.Scan(&data[i].Harga)

			fmt.Println("Sukses: Data berhasil diperbarui!")
			break 
		}
	}

	if !ketemu {
		fmt.Println("Barang dengan kode tersebut tidak ditemukan.")
	}
}

func hapusBarang(data *dataBarang, n *int) { 
	var kode string 
	var ketemu bool = false 
	var indeks int 

	fmt.Println("\n== HAPUS DATA BARANG ==") 

	if *n == 0 { 
		fmt.Println("Gudang kosong.")
		return
	}

	fmt.Print("Masukkan Kode Barang yang akan dihapus: ") 
	fmt.Scan(&kode)

	for i := 0; i < *n; i++ { 
		if data[i].Kode == kode { 
			ketemu = true
			indeks = i 
			break
		}
	}
	if ketemu { 
		fmt.Printf("Barang [%s] %s berhasil dihapus!\n", data[indeks].Kode, data[indeks].Nama)

		for i := indeks; i < *n-1; i++ { // Menggeser data array ke kiri
			data[i] = data[i+1]
		}
		(*n)-- 
	} else { 
		fmt.Println("Gagal: Barang dengan kode tersebut tidak ditemukan.")
	}
}

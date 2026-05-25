package main

const NMAX = 100

type Barang struct {
	Kode  string
	Nama  string
	Harga int
	Stok  int
}

type Transaksi struct {
	Kode   string
	Jenis  string
	Jumlah int
}

type dataBarang [NMAX]Barang
type dataTransaksi [NMAX]Transaksi

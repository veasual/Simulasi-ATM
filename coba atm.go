package main

import (
	"fmt"
)

// Deklarasi Variabel
var (
	username = "muhammad.thorieq" // Nama Akun Mahasiswa
	password = "2406441184"       // NPM Mahasiswa
	saldo    = 350000.0           // Saldo Awal
	riwayat  []string             // Riwayat Transaksi
)

// Lihat saldo
func lihatSaldo() {
	fmt.Println("Saldo Anda saat ini:", saldo)
}

// Tambah Saldo
func tambahSaldo(amount float64) {
	if amount > 0 {
		saldo += amount
		riwayat = append(riwayat, fmt.Sprintf("Tambah saldo: %.2f", amount))
		fmt.Println("Saldo berhasil ditambahkan.")
	} else {
		fmt.Println("Jumlah yang ditambahkan harus positif.")
	}
}

// Tarik Saldo
func tarikSaldo(amount float64) {
	if amount < 0 {
		fmt.Println("Jumlah tarik tidak boleh negatif.")
	} else if amount > saldo {
		fmt.Println("Tarik saldo melebihi saldo yang ada.")
	} else {
		saldo -= amount
		riwayat = append(riwayat, fmt.Sprintf("Tarik saldo: %.2f", amount))
		fmt.Println("Tarik saldo berhasil.")
	}
}

// Fungsi untuk melihat akun
func lihatAkun() {
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}

// Riwayat Transaksi
func historyTransaksi() {
	if len(riwayat) == 0 {
		fmt.Println("Tidak ada riwayat transaksi.")
	} else {
		fmt.Println("Riwayat Transaksi:")
		for _, trans := range riwayat {
			fmt.Println(trans)
		}
	}
}

// Fungsi main
func main() {
	var pilihan int
	var amount float64

	fmt.Println("Halo, selamat datang di ATM!")
	fmt.Println("Silakan login dengan username dan password.")
	fmt.Print("Masukkan Username: ")
	var inputUsername string
	fmt.Scanln(&inputUsername)

	fmt.Print("Masukkan Password: ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	if inputUsername != username || inputPassword != password {
		fmt.Println("Username atau password salah. Program dihentikan.")
		return
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tambah Saldo")
		fmt.Println("3. Tarik Saldo")
		fmt.Println("4. Lihat Akun")
		fmt.Println("5. History Transaksi")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatSaldo()
		case 2:
			fmt.Print("Masukkan jumlah yang ingin ditambahkan: ")
			fmt.Scanln(&amount)
			tambahSaldo(amount)
		case 3:
			fmt.Print("Masukkan jumlah yang ingin ditarik: ")
			fmt.Scanln(&amount)
			tarikSaldo(amount)
		case 4:
			lihatAkun()
		case 5:
			historyTransaksi()
		case 6:
			fmt.Println("Terima kasih telah menggunakan ATM. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

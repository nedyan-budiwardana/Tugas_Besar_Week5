package main

import (
	"fmt"
)

var (
	password = "2406356952"
	saldo    = [1]int{3500000}
	histori  []string
)

func main() {
	var username, inputPassword string

	fmt.Println("=== SELAMAT DATANG DI ATM ===")
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&inputPassword)

	if inputPassword != password {
		fmt.Println("Password salah.")
		return
	}

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Lihat Informasi Akun")
		fmt.Println("2. Lihat Saldo")
		fmt.Println("3. Tambahkan Saldo")
		fmt.Println("4. Tarik Saldo")
		fmt.Println("5. Histori Transaksi")
		fmt.Println("6. Keluar dari Program")

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			lihatAkun(username)
		case 2:
			lihatSaldo()
		case 3:
			tambahSaldo()
		case 4:
			tarikSaldo()
		case 5:
			historiTransaksi()
		case 6:
			fmt.Println("Terima kasih telah menggunakan layanan ATM. Selamat tinggal!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func lihatAkun(username string) {
	fmt.Println("=== INFORMASI AKUN ===")
	fmt.Println("Username:", username)
	fmt.Println("Saldo: Rp", saldo[0])
}

func lihatSaldo() {
	fmt.Println("=== SALDO AKUN ===")
	fmt.Printf("Saldo Anda: Rp %d\n", saldo[0])
}

func tambahSaldo() {
	var amount int
	fmt.Print("Masukkan jumlah yang ingin ditambahkan: Rp ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah harus lebih besar dari 0.")
		return
	}

	saldo[0] += amount
	histori = append(histori, fmt.Sprintf("Menambahkan saldo: Rp %d", amount))
	fmt.Printf("Saldo berhasil ditambahkan. Saldo baru: Rp %d\n", saldo[0])
}

func tarikSaldo() {
	var amount int
	fmt.Print("Masukkan jumlah yang ingin ditarik: Rp ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah tarik harus lebih besar dari 0.")
		return
	}

	if amount > saldo[0] {
		fmt.Println("Saldo tidak cukup untuk menarik Rp", amount)
		return
	}

	saldo[0] -= amount
	histori = append(histori, fmt.Sprintf("Tarik saldo: Rp %d", amount))
	fmt.Printf("Tarik saldo berhasil. Saldo baru: Rp %d\n", saldo[0])
}

func historiTransaksi() {
	if len(histori) == 0 {
		fmt.Println("Tidak ada transaksi.")
		return
	}

	fmt.Println("=== HISTORI TRANSAKSI ===")
	for _, entry := range histori {
		fmt.Println(entry)
	}
}

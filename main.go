package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type User struct {
	username string
	password string
}

type TouristSpot struct {
	name       string
	attributes map[string]string
}

var users []User
var spots []TouristSpot

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// Initialize tourist spots with data
	initializeTouristSpots()

	for {
		fmt.Println("*** Menu Awal ***")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			register()
		case "2":
			login()
		case "3":
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func initializeTouristSpots() {
	spots = []TouristSpot{
		{name: "Monumen Nasional", attributes: map[string]string{"lokasi": "Jakarta Pusat", "jarak": "0 km", "biaya": "Rp 20.000", "rating": "4.5"}},
		{name: "Taman Mini Indonesia Indah", attributes: map[string]string{"lokasi": "Jakarta Timur", "jarak": "17 km", "biaya": "Rp 10.000", "rating": "4.2"}},
		{name: "Museum Nasional Indonesia", attributes: map[string]string{"lokasi": "Jakarta Pusat", "jarak": "2 km", "biaya": "Rp 5.000", "rating": "4.3"}},
		{name: "Ancol Dreamland", attributes: map[string]string{"lokasi": "Jakarta Utara", "jarak": "13 km", "biaya": "Rp 25.000", "rating": "4.0"}},
		{name: "Kota Tua Jakarta", attributes: map[string]string{"lokasi": "Jakarta Barat", "jarak": "7 km", "biaya": "Gratis", "rating": "4.4"}},
		{name: "Museum Bank Indonesia", attributes: map[string]string{"lokasi": "Jakarta Barat", "jarak": "7 km", "biaya": "Rp 5.000", "rating": "4.5"}},
		{name: "Jakarta Aquarium", attributes: map[string]string{"lokasi": "Jakarta Barat", "jarak": "10 km", "biaya": "Rp 200.000", "rating": "4.6"}},
		{name: "Dunia Fantasi (Dufan)", attributes: map[string]string{"lokasi": "Jakarta Utara", "jarak": "13 km", "biaya": "Rp 300.000", "rating": "4.1"}},
		{name: "Ragunan Zoo", attributes: map[string]string{"lokasi": "Jakarta Selatan", "jarak": "15 km", "biaya": "Rp 4.000", "rating": "3.9"}},
		{name: "Setu Babakan", attributes: map[string]string{"lokasi": "Jakarta Selatan", "jarak": "20 km", "biaya": "Gratis", "rating": "4.0"}},
	}
}

func register() {
	fmt.Print("Masukkan nama pengguna baru: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Masukkan kata sandi baru: ")
	scanner.Scan()
	password := scanner.Text()

	for _, user := range users {
		if user.username == username {
			fmt.Println("Nama pengguna sudah ada, silakan gunakan nama lain.")
			return
		}
	}

	users = append(users, User{username: username, password: password})
	fmt.Println("Registrasi berhasil!")
}

func login() {
	fmt.Print("Masukkan nama pengguna: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Masukkan kata sandi: ")
	scanner.Scan()
	password := scanner.Text()

	if username == "admin_kece" && password == "admin135" {
		fmt.Println("Login sebagai admin berhasil!")
		adminMenu(scanner)
		return
	}

	for _, user := range users {
		if user.username == username && user.password == password {
			fmt.Println("Login berhasil!")
			userMenu()
			return
		}
	}

	fmt.Println("Login gagal. Nama pengguna atau kata sandi salah.")
}

func userMenu() {
	for {
		fmt.Println("Menu Utama - User")
		fmt.Println("1. Lihat Tempat Wisata")
		fmt.Println("2. Cari Tempat Wisata")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Print("Pilihan: ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &choice)

		switch choice {
		case 1:
			displayTouristSpots()
		case 2:
			showTouristSpot()
		case 3:
			fmt.Println("Keluar dari Menu User")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func showTouristSpot() {
	if len(spots) == 0 {
		fmt.Println("Belum ada data tempat wisata.")
		return
	}

	fmt.Println("Cari tempat wisata berdasarkan kategori: jarak, biaya, rating")
	fmt.Print("Masukkan kategori: ")
	scanner.Scan()
	category := scanner.Text()

	// Sort the spots based on the category input
	sort.Slice(spots, func(i, j int) bool {
		switch category {
		case "jarak":
			return spots[i].attributes["jarak"] < spots[j].attributes["jarak"]
		case "biaya":
			return spots[i].attributes["biaya"] < spots[j].attributes["biaya"]
		case "rating":
			return spots[i].attributes["rating"] > spots[j].attributes["rating"]
		default:
			return false
		}
	})

	// Display all tourist spots after sorting
	fmt.Println("Data Tempat Wisata:")
	for i, spot := range spots {
		fmt.Printf("%d. %s - Jarak: %s, Biaya: %s, Rating: %s\n", i+1, spot.name, spot.attributes["jarak"], spot.attributes["biaya"], spot.attributes["rating"])
	}
}

func displayTouristSpots() {
	fmt.Println("Data Tempat Wisata")
	for i, spot := range spots {
		fmt.Printf("%d. %s\n", i+1, spot.name)
		fmt.Printf("Jarak: %s\n", spot.attributes["jarak"])
		fmt.Printf("Biaya: %s\n", spot.attributes["biaya"])
		fmt.Printf("Rating: %s\n", spot.attributes["rating"])
	}
}

func adminMenu(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n*** Menu Admin ***")
		fmt.Println("1. Mengubah Data")
		fmt.Println("2. Menuju Menu User")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Print("Pilihan: ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &choice)

		switch choice {
		case 1:
			changeData()
		case 2:
			userMenu()
		case 3:
			fmt.Println("Keluar dari Admin")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func addTouristSpot() {
	fmt.Print("Nama Tempat Wisata: ")
	scanner.Scan()
	name := scanner.Text()

	attributes := make(map[string]string)
	fmt.Print("Masukkan lokasi: ")
	scanner.Scan()
	attributes["lokasi"] = scanner.Text()

	fmt.Print("Masukkan jarak: ")
	scanner.Scan()
	attributes["jarak"] = scanner.Text()

	fmt.Print("Masukkan biaya: ")
	scanner.Scan()
	attributes["biaya"] = scanner.Text()

	fmt.Print("Masukkan rating: ")
	scanner.Scan()
	attributes["rating"] = scanner.Text()

	spots = append(spots, TouristSpot{name: name, attributes: attributes})
	fmt.Println("Data tempat wisata berhasil ditambahkan!")
}

func updateTouristSpot() {
	if len(spots) == 0 {
		fmt.Println("Belum ada data tempat wisata.")
		return
	}

	fmt.Println("Pilih tempat wisata yang ingin diubah:")
	for i, spot := range spots {
		fmt.Printf("%d. %s\n", i+1, spot.name)
	}

	var choice int
	fmt.Print("Pilihan: ")
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &choice)

	if choice < 1 || choice > len(spots) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Print("Nama Tempat Wisata: ")
	scanner.Scan()
	name := scanner.Text()

	attributes := make(map[string]string)
	for {
		fmt.Print("Masukkan atribut (contoh: 'lokasi: Jakarta'): ")
		scanner.Scan()
		attribute := scanner.Text()

		if attribute == "" {
			break
		}

		var key, value string
		fmt.Sscanf(attribute, "%s: %s", &key, &value)
		attributes[key] = value
	}

	spots[choice-1] = TouristSpot{name: name, attributes: attributes}
	fmt.Println("Data tempat wisata berhasil diubah!")
}

func deleteTouristSpot() {
	if len(spots) == 0 {
		fmt.Println("Belum ada data tempat wisata.")
		return
	}

	fmt.Println("Pilih tempat wisata yang ingin dihapus:")
	for i, spot := range spots {
		fmt.Printf("%d. %s\n", i+1, spot.name)
	}

	var choice int
	fmt.Print("Pilihan: ")
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &choice)

	if choice < 1 || choice > len(spots) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	spots = append(spots[:choice-1], spots[choice:]...)
	fmt.Println("Data tempat wisata berhasil dihapus!")
}

func changeData() {
	for {
		fmt.Println("1. Tambah Tempat Wisata")
		fmt.Println("2. Update Tempat Wisata")
		fmt.Println("3. Hapus Tempat Wisata")
		fmt.Println("4. Lihat Data Tempat Wisata")
		fmt.Println("5. Kembali ke Menu Admin")

		var choice int
		fmt.Print("Pilihan: ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &choice)

		switch choice {
		case 1:
			addTouristSpot()
		case 2:
			updateTouristSpot()
		case 3:
			deleteTouristSpot()
		case 4:
			displayTouristSpots()
		case 5:
			fmt.Println("Kembali ke Menu Admin")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

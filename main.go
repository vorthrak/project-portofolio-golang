package main

import "fmt"

func main() {
	var names []string
	for {
		fmt.Println("\n=== CRUD Sederhana ===")
		fmt.Println("1. Tambah Nama")
		fmt.Println("2. Lihat Nama")
		fmt.Println("3. Ubah Nama")
		fmt.Println("4. Hapus Nama")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var c, i int
		fmt.Scan(&c)

		switch c {
		case 1:
			var n string
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&n)
			names = append(names, n)
			fmt.Println("✅ Nama berhasil ditambahkan!")
		case 2:
			if len(names) == 0 {
				fmt.Println("⚠️  Tidak ada data!")
			} else {
				fmt.Println("\n📜 Daftar Nama:")
				for i, n := range names {
					fmt.Printf("%d. %s\n", i+1, n)
				}
			}
		case 3:
			if len(names) == 0 {
				fmt.Println("⚠️  Tidak ada data untuk diubah!")
				continue
			}
			fmt.Print("Masukkan nomor nama yang ingin diubah: ")
			fmt.Scan(&i)
			if i > 0 && i <= len(names) {
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&names[i-1])
				fmt.Println("✅ Nama berhasil diubah!")
			} else {
				fmt.Println("❌ Nomor tidak valid!")
			}
		case 4:
			if len(names) == 0 {
				fmt.Println("⚠️  Tidak ada data untuk dihapus!")
				continue
			}
			fmt.Print("Masukkan nomor nama yang ingin dihapus: ")
			fmt.Scan(&i)
			if i > 0 && i <= len(names) {
				names = append(names[:i-1], names[i:]...)
				fmt.Println("✅ Nama berhasil dihapus!")
			} else {
				fmt.Println("❌ Nomor tidak valid!")
			}
		case 5:
			fmt.Println("🚪 Keluar...")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid! Coba lagi.")
		}
	}
}

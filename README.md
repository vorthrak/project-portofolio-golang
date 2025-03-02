# CRUD Sederhana Golang

Project ini memungkinkan user menambah, melihat, mengubah, dan menghapus nama dalam daftar.

## Cara Menjalankan
1. Install Golang
2. Simpan kode dalam `main.go`
3. Jalankan dengan:
   ```sh
   go run main.go
   ```

## Fitur
- **Tambah Nama**: Tambah nama ke daftar.
- **Lihat Nama**: Lihat semua nama yang sudah disimpan.
- **Ubah Nama**: Edit nama berdasarkan nomor urut.
- **Hapus Nama**: Hapus nama berdasarkan nomor urut.
- **Keluar**: Menghentikan program.

## Penjelasan Kode Secara Detail

### 1. Import Library
```go
import "fmt"
```
- `import "fmt"` → Mengimpor paket `fmt` untuk menangani input dan output (print ke layar, baca dari input pengguna, dll.).

### 2. Fungsi Utama `main()`
```go
func main() {
```
- `func main()` → Fungsi utama program, dieksekusi pertama kali saat program dijalankan.

### 3. Deklarasi Variabel
```go
var names []string
```
- `var names []string` → Mendeklarasikan slice kosong `names` untuk menyimpan daftar nama.

### 4. Looping Menu
```go
for {
```
- `for {}` → Loop tanpa kondisi, berjalan terus-menerus hingga program dihentikan.

#### Menampilkan Menu
```go
fmt.Println("=== CRUD Sederhana ===")
fmt.Println("1. Tambah Nama")
fmt.Println("2. Lihat Nama")
fmt.Println("3. Ubah Nama")
fmt.Println("4. Hapus Nama")
fmt.Println("5. Keluar")
fmt.Print("Pilih menu: ")
```
- `fmt.Println(...)` → Menampilkan teks menu ke layar.
- `fmt.Print(...)` → Mencetak teks tanpa baris baru agar user bisa langsung mengetik input.

#### Menerima Input Menu
```go
var c int
fmt.Scan(&c)
```
- `var c int` → Variabel `c` menyimpan pilihan menu dalam bentuk angka.
- `fmt.Scan(&c)` → Membaca input user dan menyimpannya di variabel `c`.

### 5. Switch Case untuk Menu
```go
switch c {
```
- `switch c {}` → Memeriksa nilai `c` untuk mengeksekusi perintah yang sesuai.

#### a) Menambah Nama
```go
case 1:
    var n string
    fmt.Print("Masukkan nama: ")
    fmt.Scan(&n)
    names = append(names, n)
    fmt.Println("✅ Nama ditambahkan!")
```
- `case 1:` → Menangani opsi penambahan nama.
- `var n string` → Variabel `n` untuk menyimpan nama baru.
- `fmt.Scan(&n)` → Membaca input nama dari user.
- `names = append(names, n)` → Menambahkan nama ke dalam slice `names`.

#### b) Menampilkan Nama
```go
case 2:
    if len(names) == 0 {
        fmt.Println("⚠️ Tidak ada data!")
    } else {
        for i, n := range names {
            fmt.Printf("%d. %s\n", i+1, n)
        }
    }
```
- `if len(names) == 0` → Mengecek apakah `names` kosong.
- `for i, n := range names {}` → Loop untuk mencetak daftar nama dengan nomor urut.
- `fmt.Printf("%d. %s\n", i+1, n)` → Menampilkan nomor urut (`i+1`) dan nama (`n`).

#### c) Mengubah Nama
```go
case 3:
    if len(names) == 0 {
        fmt.Println("⚠️ Tidak ada data untuk diubah!")
        continue
    }
    fmt.Print("Nomor nama yang ingin diubah: ")
    fmt.Scan(&i)
    if i > 0 && i <= len(names) {
        fmt.Print("Masukkan nama baru: ")
        fmt.Scan(&names[i-1])
        fmt.Println("✅ Nama diubah!")
    } else {
        fmt.Println("❌ Nomor tidak valid!")
    }
```
- `if len(names) == 0` → Cek apakah ada data yang bisa diubah.
- `fmt.Scan(&i)` → User memasukkan nomor nama yang ingin diubah.
- `if i > 0 && i <= len(names)` → Cek apakah nomor valid.
- `fmt.Scan(&names[i-1])` → Mengubah elemen pada indeks tertentu.

#### d) Menghapus Nama
```go
case 4:
    if len(names) == 0 {
        fmt.Println("⚠️ Tidak ada data untuk dihapus!")
        continue
    }
    fmt.Print("Nomor nama yang ingin dihapus: ")
    fmt.Scan(&i)
    if i > 0 && i <= len(names) {
        names = append(names[:i-1], names[i:]...)
        fmt.Println("✅ Nama dihapus!")
    } else {
        fmt.Println("❌ Nomor tidak valid!")
    }
```
- `if len(names) == 0` → Cek apakah ada data untuk dihapus.
- `fmt.Scan(&i)` → User memasukkan nomor nama yang ingin dihapus.
- `names = append(names[:i-1], names[i:]...)` → Menghapus elemen slice dengan cara menggabungkan dua bagian slice tanpa elemen yang dihapus.

#### e) Keluar Program
```go
case 5:
    fmt.Println("🚪 Keluar...")
    return
```
- `return` → Menghentikan program.

#### f) Jika Input Tidak Valid
```go
default:
    fmt.Println("❌ Pilihan tidak valid! Coba lagi.")
```
- `default:` → Menangani input yang tidak valid.

## Contoh Output
```
=== CRUD Sederhana ===
1. Tambah Nama
2. Lihat Nama
3. Ubah Nama
4. Hapus Nama
5. Keluar
Pilih menu: 1
Masukkan nama: Alice
✅ Nama ditambahkan!
```

CRUD ini simpel, cocok untuk latihan Golang. 🚀


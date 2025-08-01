# HSI-Sandbox Week 1 - Pemrograman GO

**Nama:** Muhammad Fadhil Zurani  
**NIP:** ARN251-06144  
**Tugas:** Week 1 - Pemrograman GO

## Deskripsi

Repository ini berisi solusi untuk 3 soal pemrograman GO yang diberikan pada tugas HSI-Sandbox Week 1. Setiap soal dikerjakan dalam folder terpisah dengan implementasi yang sesuai dengan requirement yang diminta.

## Struktur Proyek

```
week-1-go/
├── README.md
├── Soal-1/
│   └── faktorial.go
├── Soal-2/
│   └── palindrome.go
└── Soal-3/
    └── bilangan_besar.go
```

## Soal yang Dikerjakan

### Soal 1: Faktorial

**File:** `Soal-1/faktorial.go`

**Deskripsi:** Program untuk menghitung faktorial dari sebuah bilangan bulat non-negatif.

**Fitur:**

- Menghitung faktorial menggunakan iterasi
- Validasi input untuk bilangan non-negatif
- Menangani kasus khusus faktorial 0 dan 1

**Cara Menjalankan:**

```bash
cd Soal-1
go run faktorial.go
```

**Contoh Output:**

```
Masukkan angka: 5
Faktorial dari 5 adalah 120
```

### Soal 2: Palindrome Checker

**File:** `Soal-2/palindrome.go`

**Deskripsi:** Program untuk mengecek apakah sebuah string adalah palindrome (sama jika dibaca dari depan dan belakang).

**Fitur:**

- Mengecek palindrome dengan case-insensitive
- Menggunakan algoritma two-pointer untuk efisiensi
- Output yang informatif dengan penjelasan

**Cara Menjalankan:**

```bash
cd Soal-2
go run palindrome.go
```

**Contoh Output:**

```
Input: madam
Output: true (karena "madam" dibaca sama dari depan dan belakang)
```

### Soal 3: Bilangan Terbesar

**File:** `Soal-3/bilangan_besar.go`

**Deskripsi:** Program untuk mencari bilangan terbesar dalam sebuah slice/array.

**Fitur:**

- Input dinamis untuk jumlah elemen
- Validasi input untuk jumlah elemen yang valid
- Pencarian nilai maksimum menggunakan algoritma linear search

**Cara Menjalankan:**

```bash
cd Soal-3
go run bilangan_besar.go
```

**Contoh Output:**

```
Masukkan jumlah elemen: 5
Masukkan elemen-elemen slice:
Elemen ke-1: 10
Elemen ke-2: 25
Elemen ke-3: 5
Elemen ke-4: 30
Elemen ke-5: 15
Output: 30 (karena 30 adalah bilangan terbesar dalam slice tersebut)
```

## Teknologi yang Digunakan

- **Bahasa Pemrograman:** Go (Golang)
- **Versi Go:** 1.x atau lebih baru
- **Package yang Digunakan:**
  - `fmt` - untuk input/output
  - `strings` - untuk manipulasi string (Soal 2)

## Cara Menjalankan Semua Program

1. **Pastikan Go terinstall di sistem Anda:**

   ```bash
   go version
   ```

2. **Clone atau download repository ini**

3. **Jalankan setiap program sesuai dengan instruksi di atas**

## Konsep Pemrograman yang Diimplementasikan

### Soal 1 - Faktorial

- **Iterasi:** Menggunakan loop for untuk perhitungan faktorial
- **Conditional:** Pengecekan kondisi untuk bilangan negatif
- **Function:** Pembuatan fungsi factorial yang reusable

### Soal 2 - Palindrome

- **String Manipulation:** Menggunakan `strings.ToLower()` untuk case-insensitive
- **Algorithm:** Two-pointer technique untuk efisiensi
- **Boolean Logic:** Return boolean value dengan penjelasan

### Soal 3 - Bilangan Terbesar

- **Dynamic Arrays:** Menggunakan slice dengan ukuran dinamis
- **Loop:** Iterasi untuk input dan pencarian maksimum
- **Array Manipulation:** Akses dan perbandingan elemen array

## Pembelajaran

Melalui tugas ini, saya telah mempelajari:

- Sintaks dasar bahasa pemrograman Go
- Penggunaan package `fmt` untuk I/O operations
- Implementasi algoritma dasar (iterasi, pencarian, string processing)
- Best practices dalam penulisan kode Go
- Error handling dan validasi input

## Kontribusi

Tugas ini dikerjakan secara mandiri sebagai bagian dari pembelajaran HSI-Sandbox Week 1.

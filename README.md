# 🧮 Kalkulator Sederhana

Aplikasi kalkulator web sederhana yang dibangun menggunakan Go dan HTML/CSS murni.

## ✨ Fitur

- **Operasi Matematika Dasar**: Penjumlahan (+), pengurangan (-), perkalian (*), dan pembagian (/)
- **Interface Modern**: Desain responsif dengan gradient background dan efek visual yang menarik
- **Keyboard Support**: Mendukung input keyboard untuk angka dan operator
- **Error Handling**: Menampilkan pesan error yang informatif
- **Responsive Design**: Bekerja dengan baik di desktop dan mobile

## 🚀 Cara Menjalankan

### Prerequisites
- Go 1.21 atau lebih baru

### Langkah-langkah

1. **Clone atau download project ini**

2. **Masuk ke direktori project**
   ```bash
   cd Kalkulator1
   ```

3. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```

4. **Buka browser dan kunjungi**
   ```
   http://localhost:8080
   ```

## 📁 Struktur Project

```
Kalkulator1/
├── main.go              # File utama aplikasi Go
├── go.mod               # File konfigurasi module Go
├── README.md            # Dokumentasi project
└── templates/
    └── index.html       # Template HTML untuk interface kalkulator
```

## 🛠️ Teknologi yang Digunakan

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3, JavaScript
- **Server**: HTTP server bawaan Go
- **Template Engine**: Go HTML templates

## 📝 Cara Penggunaan

1. **Input Manual**: Ketik langsung di input field
   - Contoh: `5 + 3`, `10 * 2`, `15 / 3`

2. **Menggunakan Tombol**: Klik tombol angka dan operator
   - Klik tombol angka dan operator untuk membangun ekspresi
   - Klik tombol "= Hitung" untuk melihat hasil

3. **Keyboard Shortcuts**:
   - Angka: `0-9`
   - Operator: `+`, `-`, `*`, `/`
   - Enter: Hitung hasil
   - Backspace/Delete: Hapus karakter

## 🔧 Operasi yang Didukung

| Operasi | Simbol | Contoh |
|---------|--------|--------|
| Penjumlahan | `+` | `5 + 3 = 8` |
| Pengurangan | `-` | `10 - 4 = 6` |
| Perkalian | `*` | `6 * 7 = 42` |
| Pembagian | `/` | `15 / 3 = 5` |

## 🎨 Fitur UI/UX

- **Gradient Background**: Background gradient yang menarik
- **Glassmorphism Effect**: Efek kaca buram pada kalkulator
- **Hover Effects**: Animasi hover pada tombol
- **Responsive Grid**: Layout grid yang responsif
- **Color Coding**: Warna berbeda untuk tiap jenis tombol
- **Error Display**: Pesan error yang jelas dan informatif

## 🐛 Troubleshooting

### Port 8080 sudah digunakan
Jika port 8080 sudah digunakan, Anda dapat mengubah port di file `main.go`:
```go
log.Fatal(http.ListenAndServe(":8081", nil)) // Ganti 8080 dengan port lain
```

### Error "template not found"
Pastikan file `templates/index.html` ada dan berada di direktori yang benar.

## 📄 License

Project ini dibuat untuk tujuan pembelajaran dan dapat digunakan secara bebas.

## 🤝 Kontribusi

Silakan berkontribusi dengan:
- Melaporkan bug
- Menambahkan fitur baru
- Memperbaiki dokumentasi

---

**Dibuat dengan ❤️ menggunakan Go dan HTML/CSS** 
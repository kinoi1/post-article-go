# post-article-go
# Go Post Article API

API sederhana dan cepat yang dibangun menggunakan **Golang**, **Gin Gonic (Web Framework)**, dan **GORM (ORM)** untuk manajemen artikel (Post Article) dengan database **PostgreSQL**. Proyek ini sudah dilengkapi dengan fitur validasi input otomatis, CORS terkonfigurasi, dan arsitektur kode yang terpisah (*Clean Structure*).

kenapa pake postgre? 
jyujur lupa gara2 project belajar sebelumnya yang pake golang dbnya postgre >_<!!
---

## 🚀 Fitur Utama

- **RESTful API Architecture** dengan Router Gin Gonic.
- **ORM PostgreSQL** menggunakan GORM (termasuk *Auto Migration* tabel).
- **Data Validation Layer** terpusat di `utils/validator.go` untuk mengecek aturan minimal karakter dan status pilihan.
- **CORS Enabled** untuk integrasi aman dengan frontend (Next.js/React/Vue).
- **Pagination Feature** (Limit & Offset) untuk endpoint penarikan semua data artikel.

---

## 🛠️ Struktur Kolom Database (`posts`)

| Nama Kolom | Tipe Data | Keterangan / Validasi API |
| :--- | :--- | :--- |
| **id** | `Int` | Primary Key, Auto Increment |
| **title** | `Varchar(200)` | Required, Minimal 20 Karakter |
| **content** | `Text` | Required, Minimal 200 Karakter |
| **category** | `Varchar(100)` | Required, Minimal 3 Karakter |
| **status** | `Varchar(100)` | Pilihan wajib: `Publish`, `Draft`, `Thrash` |
| **created_date** | `Timestamp` | Diisi otomatis saat data dibuat |
| **updated_date** | `Timestamp` | Diperbarui otomatis saat terjadi perubahan |

---

## 📦 Persyaratan Awal (Prerequisites)

Sebelum menjalankan aplikasi, pastikan komputer Anda telah terinstal:
- [Go](https://go.dev/doc/install) (versi 1.18 atau lebih baru)
- [PostgreSQL](https://www.postgresql.org/download/)

---

## ⚙️ Instalasi & Menjalankan Aplikasi

1. **Clone Repositori**
   ```bash
   git clone [https://github.com/kinoi1/post-article-go.git](https://github.com/kinoi1/post-article-go.git)
   cd go-post-article
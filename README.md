# Ticketing System

## Deskripsi
Ticketing System adalah aplikasi backend yang dirancang untuk mengelola penjualan tiket, pemesanan event, dan laporan penjualan. Aplikasi ini mendukung otentikasi berbasis JWT serta role-based access control untuk memastikan keamanan akses data.

---

## Fitur Utama

### 1. Otentikasi Pengguna
- Registrasi dan login pengguna menggunakan email dan password.
- Token JWT digunakan untuk autentikasi.

### 2. Manajemen Event
- **Admin:**
  - Membuat, mengupdate, dan menghapus event.
- **User:**
  - Melihat daftar event.

### 3. Manajemen Tiket
- **User:**
  - Memesan tiket untuk event tertentu.
  - Melihat tiket yang dipesan.
- **Admin:**
  - Melihat semua tiket yang dipesan.

### 4. Laporan
- Ringkasan tiket terjual dan total pendapatan.
- Laporan berdasarkan event tertentu.

---

## Teknologi yang Digunakan

### Backend
- Golang dengan Gin Framework

### Database
- MySQL

### Autentikasi
- JSON Web Token (JWT)

---

## Struktur Folder
```
.
|-- config
|   |-- database.go          // Koneksi database
|   |-- test_db.go           // Database untuk testing
|
|-- controller
|   |-- event_controller.go  // Handler untuk event
|   |-- report_controller.go // Handler untuk laporan
|   |-- ticket_controller.go // Handler untuk tiket
|   |-- user_controller.go   // Handler untuk user
|
|-- entity
|   |-- event.go             // Model event
|   |-- reports.go           // Model laporan
|   |-- ticket.go            // Model tiket
|   |-- user.go              // Model user
|
|-- middleware
|   |-- auth_middleware.go   // Middleware autentikasi
|   |-- rbac_middleware.go   // Middleware otorisasi
|
|-- repository
|   |-- event_repository.go  // Query terkait event
|
|-- response
|   |-- response.go          // Struct respons untuk client
|
|-- service
|   |-- event_service.go     // Logika bisnis event
|   |-- report_service.go    // Logika bisnis laporan
|   |-- ticket_service.go    // Logika bisnis tiket
|
|-- transaction
|   |-- transaction.go       // Logika transaksi tiket
|
|-- main.go                  // Entry point aplikasi
```

---

## Cara Menjalankan Aplikasi

1. **Clone Repository**
   ```bash
   git clone https://github.com/username/ticketing-system.git
   cd ticketing-system
   ```

2. **Instalasi Dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup Database**
   - Buat database di MySQL.
   - Perbarui konfigurasi database di `config/database.go`.

4. **Jalankan Aplikasi**
   ```bash
   go run main.go
   ```

5. **Endpoint yang Tersedia**
   - **User Routes:**
     - `POST /register` - Registrasi user.
     - `POST /login` - Login user.
     - `GET /events` - Melihat semua event.
   - **Admin Routes:**
     - `POST /admin/events` - Membuat event.
     - `PUT /admin/events/:id` - Mengupdate event.
     - `DELETE /admin/events/:id` - Menghapus event.
     - `GET /admin/reports` - Mendapatkan laporan.

---

## Pengujian
1. Pastikan database testing sudah terkonfigurasi di `config/test_db.go`.
2. Jalankan test:
   ```bash
   go test ./...
   ```

---

## Kontribusi
Silakan ajukan pull request jika ingin berkontribusi pada pengembangan aplikasi ini.

---

## Lisensi
Aplikasi ini menggunakan lisensi [MIT](LICENSE).

# PRBCareAPI

PRBCareAPI adalah aplikasi REST API yang dibuat dengan bahasa Golang untuk manajemen Puskesmas, manajemen Apotek, pengambilan obat, kontrol balik, dan manajemen pasien. Aplikasi ini menyediakan fungsionalitas khusus berdasarkan peran pengguna yang berbeda, termasuk Super Admin, Admin Puskesmas, Admin Apotek, dan Calon Pasien.
PRBCareAPI dikembangkan dengan mengikuti prinsip-prinsip REST API untuk memastikan skalabilitas dan pemeliharaan yang mudah. Sistem autentikasi dilengkapi untuk memastikan keamanan data.

## Pengguna dan Hak Akses

- **Super Admin**: Dapat melakukan semua operasi manajemen dalam sistem, termasuk manajemen pengguna.
- **Admin Apotek**: Mengelola obat dan proses pengambilan obat.
- **Admin Puskesmas**: Mengelola pasien, kontrol balik, dan pengambilan obat sesuai resep, serta mengarahkan ke Admin Apotek yang terkait.
- **Calon Pasien**: Dapat mengakses informasi terkait dengan layanan Puskesmas dan Apotek yang mereka butuhkan.

## Fitur

- Manajemen pengguna dengan autentikasi yang berbeda untuk Super Admin, Admin Puskesmas, dan Admin Apotek.
- Manajemen pasien yang meliputi pendaftaran, pembaruan data, dan pencatatan medis.
- Manajemen obat oleh Admin Apotek, termasuk stok dan dispensasi obat.
- Kontrol balik untuk memonitor dan mengevaluasi pengobatan pasien.
- Sistem pembuatan jadwal kontrol balik dan pengambilan obat.

## Dokumentasi API

Untuk mendapatkan lebih detail mengenai endpoint dan cara penggunaan API, kunjungi dokumentasi API di link berikut:

[API Documentation](https://bump.sh/sckiddie/doc/prb-care-api)

## Contoh Implementasi Frontend

Lihat contoh implementasi frontend untuk aplikasi PRBCareAPI di link berikut:

[Frontend Implementation](https://github.com/RyanAprs/PRB-Care-Client.git)

## Aplikasi Scheduler Pendukung

Aplikasi scheduler mendukung pengingat melalui push notifikasi dan pembatalan jadwal secara otomatis. Informasi lebih lanjut dan dokumentasi aplikasi scheduler dapat diakses melalui link berikut:

[Scheduler Application](https://github.com/scrkiddie/PRBCareScheduler)


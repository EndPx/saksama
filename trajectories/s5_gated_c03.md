# Trajectory — s5_gated — c03

## checklist:PP35-12

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 12 — PKWT dilarang mensyaratkan masa percobaan: PKWT tidak dapat mensyaratkan adanya masa percobaan kerja. Apabila masa percobaan tetap disyaratkan dalam PKWT, masa percobaan tersebut batal demi hukum dan masa kerja pekerja tetap dihitung sejak awal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"Perjanjian ini tidak mensyaratkan masa percobaan.","deskripsi":"Kontrak secara eksplisit menyatakan tidak ada masa percobaan, sehingga patuh terhadap PP 35/2021 Pasal 12."}
```

## checklist:UU13-60-1

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (1) — PKWTT masa percobaan paling lama 3 bulan: PKWTT dapat mensyaratkan masa percobaan kerja paling lama 3 bulan. Ketentuan yang menetapkan masa percobaan lebih dari 3 bulan tidak berkekuatan hukum untuk kelebihannya, dan kelebihan waktu itu dihitung sebagai masa kerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"Perjanjian ini tidak mensyaratkan masa percobaan.","deskripsi":"Kontrak secara eksplisit menyatakan tidak ada masa percobaan, sehingga tidak melanggar Pasal 60 ayat (1) UU 13/2003."}
```

## checklist:UU13-60-2

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (2) — Upah masa percobaan tidak boleh di bawah upah minimum: Selama masa percobaan kerja, pengusaha dilarang membayar upah di bawah upah minimum yang berlaku di wilayah tempat pekerja bekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku… Perjanjian ini tidak mensyaratkan masa percobaan.","deskripsi":"Kontrak tidak melanggar Pasal 60 ayat (2) UU 13/2003. Upah ditetapkan di atas upah minimum dan perjanjian secara eksplisit menyatakan tidak ada masa percobaan, sehingga tidak ada klausa pembayaran upah di bawah upah minimum selama masa percobaan."}
```

## checklist:PP35-4-2

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 4 ayat (2) — PKWT tidak untuk pekerjaan bersifat tetap: PKWT didasarkan pada jangka waktu atau selesainya suatu pekerjaan tertentu, dan tidak dapat diadakan untuk pekerjaan yang bersifat tetap. Penilaian ini bergantung pada sifat pekerjaan, bukan semata-mata bunyi klausa.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"PKWT yang didasarkan pada selesai-nya suatu pekerjaan tertentu... bukan PKWT berdasarkan jangka waktu","deskripsi":"Kontrak menyatakan PKWT didasarkan pada selesainya pekerjaan tertentu (Proyek JPKIT), bukan pekerjaan tetap. Estimasi 6 tahun hanyalah perkiraan, dan perjanjian berakhir otomatis saat Berita Acara Serah Terima Akhir diterbitkan, sesuai Pasal 5 ayat (2) jo Pasal 9 PP 35/2021."}
```

## checklist:PP35-8

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 8 — PKWT jangka waktu paling lama 5 tahun termasuk perpanjangan: PKWT berdasarkan jangka waktu dibuat paling lama 5 tahun, dan jangka waktu keseluruhan termasuk perpanjangan tidak boleh lebih dari 5 tahun. Batas ini hanya berlaku untuk PKWT berdasarkan jangka waktu (Pasal 5 ayat 1); PKWT berdasarkan selesainya suatu pekerjaan diatur Pasal 9 dan tidak memiliki plafon 5 tahun. Ditegaskan pula dalam Putusan MK 168/PUU-XXI/2023.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"Perjanjian ini merupakan PKWT yang didasarkan pada selesainya suatu pekerjaan tertentu sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.","deskripsi":"Kontrak secara eksplisit menyatakan sebagai PKWT berdasarkan selesainya pekerjaan tertentu (Pasal 9 PP 35/2021), bukan PKWT berdasarkan jangka waktu, sehingga plafon 5 tahun Pasal 8 tidak berlaku."}
```

## checklist:PP35-16

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 16 — Besaran uang kompensasi PKWT: Besaran uang kompensasi: PKWT selama 12 bulan terus-menerus sebesar 1 bulan upah; PKWT 1 bulan atau lebih tetapi kurang dari 12 bulan dihitung secara proporsional dengan rumus masa kerja dibagi 12 dikali 1 bulan upah; PKWT lebih dari 12 bulan juga dihitung proporsional. Pada usaha mikro dan usaha kecil, besaran uang kompensasi ditentukan berdasarkan kesepakatan pengusaha dan pekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 8","kutipan":"Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.","deskripsi":"Klausa kompensasi mengacu pada PP 35/2021 tanpa mengurangi atau mengubah besaran, sehingga tidak melanggar ketentuan."}
```

## checklist:PP35-17

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 17 — Kompensasi tetap wajib bila hubungan kerja diakhiri lebih awal: Apabila salah satu pihak mengakhiri hubungan kerja sebelum berakhirnya jangka waktu PKWT, pengusaha tetap wajib memberikan uang kompensasi yang besarannya dihitung berdasarkan jangka waktu PKWT yang telah dijalani oleh pekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 8","kutipan":"Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.","deskripsi":"Pasal 8 kontrak menyatakan pengusaha wajib memberikan uang kompensasi sesuai PP 35/2021, sehingga tidak ada klausa yang melanggar Pasal 17 PP 35/2021."}
```

## checklist:PP35-26-31

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 26 dan Pasal 31 — Batas dan upah kerja lembur: Waktu kerja lembur paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu. Upah kerja lembur untuk jam pertama sebesar 1,5 kali upah sejam, dan untuk setiap jam berikutnya sebesar 2 kali upah sejam. Upah sejam dihitung sebesar 1/173 dari upah sebulan.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausa spesifik tentang kerja lembur yang melanggar PP 35/2021 Pasal 26 dan 31. Pasal 6 hanya mengatur waktu kerja 5 hari per minggu tanpa mengatur batas lembur, perhitungan upah lembur, maupun upah sejam, sehingga tidak ada klausa yang jelas melanggar ketentuan tersebut."}
```

## checklist:MK168-79-2b

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
Putusan MK 168/PUU-XXI/2023 atas Pasal 79 ayat (2) huruf b UU 13/2003 dalam Pasal 81 angka 25 UU 6/2023 — Lima hari kerja wajib memberi 2 hari istirahat mingguan: Mahkamah Konstitusi menyatakan norma tersebut inkonstitusional sepanjang tidak dimaknai mencakup frasa istirahat mingguan 2 hari untuk 5 hari kerja dalam 1 minggu. Konsekuensinya, kontrak dengan pola 5 hari kerja dalam 1 minggu wajib memberikan 2 hari istirahat mingguan.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"Pasal 6","kutipan":"Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.","deskripsi":"Pasal 6 secara eksplisit menetapkan pola 5 hari kerja dengan 2 hari istirahat mingguan, sesuai putusan MK 168/PUU-XXI/2023."}
```

## checklist:SE-M5-2025

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
Surat Edaran Menaker M/5/HK.04.00/V/2025 tertanggal 20 Mei 2025 — Larangan menahan ijazah dan dokumen pribadi sebagai jaminan: Pemberi kerja dilarang mensyaratkan atau menahan ijazah dan dokumen pribadi pekerja sebagai jaminan untuk bekerja, mencakup antara lain sertifikat kompetensi, paspor, akta kelahiran, buku nikah, dan BPKB. Pengecualian hanya berlaku jika ijazah atau sertifikat kompetensi diperoleh melalui pendidikan dan pelatihan yang dibiayai pemberi kerja berdasarkan perjanjian kerja tertulis, dengan kewajiban pemberi kerja menjamin keamanan dokumen dan memberi ganti rugi bila dokumen rusak atau hilang. Surat Edaran bukan peraturan perundang-undangan yang mengikat dan tidak memiliki sanksi, sehingga daya paksanya terbatas; memo harus menyatakan hal ini secara eksplisit dan tidak menyebut klausa penahanan ijazah sebagai ilegal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausul penahanan ijazah atau dokumen pribadi (sertifikat kompetensi, paspor, BPKB, dll.) sebagai jaminan kerja. Perihal tersebut tidak disinggung sama sekali."}
```

## absence:PP35-13

**Outcome:** 0 dari 9 item wajib tidak ada

_Prompt:_

```
PP 35/2021 Pasal 13 mewajibkan PKWT memuat sembilan hal berikut:
1. nama, alamat perusahaan, dan jenis usaha
2. nama, jenis kelamin, umur, dan alamat pekerja
3. jabatan atau jenis pekerjaan
4. tempat pekerjaan
5. besaran dan cara pembayaran upah
6. hak dan kewajiban pengusaha dan pekerja
7. mulai dan jangka waktu berlakunya PKWT
8. tempat dan tanggal PKWT dibuat
9. tanda tangan para pihak

Untuk setiap nomor, tentukan apakah hal itu ADA dalam kontrak. Balas HANYA JSON: {"missing":[daftar NOMOR item yang TIDAK ada]}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"missing":[]}
```

## absence:PP35-15

**Outcome:** memenuhi / tidak ada temuan

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 15 — Uang kompensasi wajib pada berakhirnya PKWT: Pengusaha wajib memberikan uang kompensasi kepada pekerja pada saat berakhirnya PKWT, bagi pekerja yang telah mempunyai masa kerja paling sedikit 1 bulan secara terus-menerus. Kewajiban ini tidak berlaku bagi tenaga kerja asing.

Apakah kontrak berikut MEMENUHI ketentuan di atas? Klausa yang membuat kewajiban menjadi bersyarat atau menghapusnya dihitung sebagai TIDAK memenuhi. Balas HANYA JSON: {"memenuhi":true|false,"section":"Pasal N atau ABSENT","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"memenuhi":true,"section":"Pasal 15","deskripsi":"Pasal 8 kontrak menyatakan PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA pada saat berakhirnya PKWT sesuai PP 35/2021, tanpa syarat pengecualian. Tidak ada klausa yang menghapuskan atau mensyaratkan kewajiban kompensasi. Dimas adalah WNI (bukan TKA) dan masa kerja telah melewati 1 bulan."}
```

## absence:PP35-27-5

**Outcome:** tidak ada pengecualian lembur / tidak ada temuan

_Prompt:_

```
Periksa ketentuan lembur dalam kontrak berikut. Apakah kontrak MEMUAT klausa yang menyatakan suatu jabatan dikecualikan atau TIDAK BERHAK atas upah lembur? Jika kontrak TIDAK memuat pernyataan pengecualian lembur seperti itu, jawab {"melanggar":false}. Jika kontrak MEMUAT pernyataan bahwa jabatan tidak berhak lembur TETAPI tidak mendefinisikan golongan jabatan tertentu (pemikir, perencana, pelaksana, atau pengendali jalannya perusahaan), jawab {"melanggar":true,"section":"Pasal N","kutipan":"kutipan harfiah"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)
## Berdasarkan Selesainya Suatu Pekerjaan Tertentu

Nomor: 007/PKWT-PROJ/II/2026

Pada hari ini, Senin, tanggal 3 Februari 2026, bertempat di Surabaya, yang bertanda tangan di bawah ini:

1. Nama: Hendra Wijaya; Jabatan: Direktur Operasional; dalam hal ini bertindak untuk dan atas nama **PT Bumi Konstruksi Mandiri**, berkedudukan di Jalan Basuki Rahmat Nomor 120, Surabaya, yang bergerak dalam bidang usaha jasa konstruksi dan rekayasa sipil. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Dimas Anggoro; Jenis Kelamin: Laki-laki; Umur: 25 tahun; Alamat: Jalan Kertajaya Nomor 30, Surabaya, Jawa Timur. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Dasar dan Jenis Perjanjian
Perjanjian ini merupakan PKWT yang didasarkan pada **selesainya suatu pekerjaan tertentu** sebagaimana dimaksud dalam Pasal 5 ayat (2) dan Pasal 9 Peraturan Pemerintah Nomor 35 Tahun 2021, bukan PKWT berdasarkan jangka waktu.

## Pasal 2 — Ruang Lingkup dan Ukuran Selesainya Pekerjaan
Pekerjaan tertentu yang dimaksud adalah perencanaan, pembangunan, pengujian, dan serah terima Jembatan Penghubung Kawasan Industri Timur (Proyek "JPKIT"). Pekerjaan dinyatakan selesai pada saat diterbitkannya Berita Acara Serah Terima Akhir (Final Hand Over) yang ditandatangani pemberi proyek. Perjanjian berakhir dengan sendirinya pada saat pekerjaan tersebut selesai.

## Pasal 3 — Perkiraan Jangka Waktu
Berdasarkan rencana kerja proyek, penyelesaian pekerjaan diperkirakan memerlukan waktu 6 (enam) tahun. Perkiraan ini bukan merupakan batas waktu perjanjian; apabila pekerjaan belum selesai, jangka waktu diperpanjang sampai pekerjaan selesai sesuai Pasal 9 PP 35/2021.

## Pasal 4 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Site Engineer dan ditempatkan di lokasi Proyek JPKIT serta kantor proyek PIHAK PERTAMA di Surabaya.

## Pasal 5 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp8.000.000,00 (delapan juta rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 28 melalui transfer ke rekening bank atas nama PIHAK KEDUA. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 6 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan, sesuai pengaturan jam kerja proyek.

## Pasal 7 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas sesuai standar keselamatan konstruksi.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini karena selesainya pekerjaan, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Surabaya pada tanggal 3 Februari 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Hendra Wijaya                         Dimas Anggoro
```

_Response:_

```
{"melanggar":false}
```

## Findings kept (0)


## Findings rejected by citation gate (0)


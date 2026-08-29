# Trajectory — s3_checklist — c12

## checklist:PP35-12

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 12 — PKWT dilarang mensyaratkan masa percobaan: PKWT tidak dapat mensyaratkan adanya masa percobaan kerja. Apabila masa percobaan tetap disyaratkan dalam PKWT, masa percobaan tersebut batal demi hukum dan masa kerja pekerja tetap dihitung sejak awal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":true,"section":"Pasal 3","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.","deskripsi":"Pasal 3 mensyaratkan masa percobaan selama 3 bulan dalam PKWT, yang dilarang dan batal demi hukum sesuai PP 35/2021 Pasal 12."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"Pasal 3","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.","deskripsi":"Masa percobaan 3 bulan sesuai batas maksimal Pasal 60 ayat (1) UU 13/2003, tidak ada kelebihan."}
```

## checklist:UU13-60-2

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (2) — Upah masa percobaan tidak boleh di bawah upah minimum: Selama masa percobaan kerja, pengusaha dilarang membayar upah di bawah upah minimum yang berlaku di wilayah tempat pekerja bekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":true,"section":"Pasal 4","kutipan":"Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku.","deskripsi":"Klausul Pasal 4 menetapkan upah masa percobaan (Rp3.150.000,00) di bawah upah minimum kota yang berlaku, yang secara eksplisit melanggar Pasal 60 ayat (2) UU 13/2003."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"Pasal 2","kutipan":"PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.","deskripsi":"Pekerjaan sebagai Staf Penjualan di kantor pusat bersifat tetap/permanen (bukan musiman atau pekerjaan tertentu yang penyelesaiannya terbatas), namun kontrak hanya menyatakan PKWT berdasarkan jangka waktu 12 bulan tanpa dasar pekerjaan tertentu. Meskipun ada indikasi sifat pekerjaan tetap, kontrak tidak secara eksplisit melanggar Pasal 4 ayat (2) PP 35/2021 karena PP 35/2021 Pasal 4 ayat (2) mensyaratkan penilaian berdasarkan sifat pekerjaan, namun di sini tidak ada klausa yang secara eksplisit menyatakan pekerjaan tersebut sesuai PKWT. Ambigu, sehingga dianggap tidak melanggar secara eksplisit."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027","deskripsi":"Kontrak PKWT berdasarkan jangka waktu selama 12 bulan, tidak melebihi plafon 5 tahun. Tidak ada klausa perpanjangan yang melebihi batas. Patuh terhadap PP 35/2021 Pasal 8."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"Pasal 7","kutipan":"Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.","deskripsi":"Klausul kompensasi mengacu pada ketentuan peraturan perundang-undangan, tidak menetapkan besaran yang bertentangan dengan PP 35/2021 Pasal 16."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"Pasal 7","kutipan":"Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.","deskripsi":"Pasal 7 hanya mengatur kompensasi pada saat berakhirnya PKWT, tidak secara eksplisit menolak atau mengecualikan kompensasi apabila hubungan kerja diakhiri lebih awal. Tidak ada klausa yang secara jelas melanggar Pasal 17 PP 35/2021."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausa spesifik mengenai waktu kerja lembur maupun upah lembur, sehingga tidak secara eksplisit melanggar PP 35/2021 Pasal 26 dan 31. Catatan: Pasal 4 menyatakan upah masa percobaan di bawah upah minimum, yang merupakan pelanggaran terpisah di luar cakupan pertanyaan ini."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.","deskripsi":"Kontrak menyatakan 5 hari kerja dengan 2 hari istirahat mingguan, sesuai Putusan MK 168/PUU-XXI/2023."}
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

Nomor: 093/HR-PKWT/VIII/2026

Pada hari ini, Jumat, tanggal 28 Agustus 2026, bertempat di Makassar, yang bertanda tangan di bawah ini:

1. Nama: Andi Baso; Jabatan: Manajer HRD; dalam hal ini bertindak untuk dan atas nama **PT Timur Raya Niaga**, berkedudukan di Jalan Boulevard Nomor 12, Makassar, yang bergerak dalam bidang usaha distribusi barang konsumen. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Nurul Aisyah; Jenis Kelamin: Perempuan; Umur: 22 tahun; Alamat: Jalan Perintis Kemerdekaan Nomor 66, Makassar, Sulawesi Selatan. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jangka Waktu
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan sejak 1 September 2026 sampai dengan 31 Agustus 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Penjualan dan ditempatkan di kantor PIHAK PERTAMA di Makassar.

## Pasal 3 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.

## Pasal 4 — Upah dan Cara Pembayaran
Upah pokok PIHAK KEDUA setelah masa percobaan adalah Rp4.500.000,00 (empat juta lima ratus ribu rupiah) per bulan. Selama masa percobaan, PIHAK KEDUA menerima upah sebesar 70 persen dari upah pokok tersebut, yaitu Rp3.150.000,00, yang berada di bawah upah minimum kota yang berlaku. Upah dibayarkan setiap tanggal 25 melalui transfer bank.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA sesuai ketentuan peraturan perundang-undangan.

## Pasal 8 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah untuk mufakat.

## Pasal 9 — Penutup
Perjanjian ini dibuat dan ditandatangani di Makassar pada tanggal 28 Agustus 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andi Baso                             Nurul Aisyah
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausa penahanan ijazah atau dokumen pribadi pekerja sebagai jaminan. Tidak ada pelanggaran terhadap SE Menaker M/5/HK.04.00/V/2025."}
```

## Findings kept (2)

- PP35-12 @ Pasal 3
- UU13-60-2 @ Pasal 4

## Findings rejected by citation gate (0)


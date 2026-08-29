# Trajectory — s3_checklist — c02

## checklist:PP35-12

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 12 — PKWT dilarang mensyaratkan masa percobaan: PKWT tidak dapat mensyaratkan adanya masa percobaan kerja. Apabila masa percobaan tetap disyaratkan dalam PKWT, masa percobaan tersebut batal demi hukum dan masa kerja pekerja tetap dihitung sejak awal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"Perjanjian ini tidak mensyaratkan masa percobaan.","deskripsi":"Kontrak secara eksplisit menyatakan tidak ada masa percobaan, sehingga patuh terhadap PP 35/2021 Pasal 12."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"Perjanjian ini tidak mensyaratkan masa percobaan.","deskripsi":"Kontrak tidak memuat klausa masa percobaan, sehingga tidak melanggar Pasal 60 ayat (1) UU 13/2003."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"Perjanjian ini tidak mensyaratkan masa percobaan.","deskripsi":"Kontrak tidak memuat klausul masa percobaan, sehingga Pasal 60 ayat (2) UU 13/2003 tidak berlaku. Upah Rp5.200.000 juga disebut di atas upah minimum yang berlaku."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"","kutipan":"","deskripsi":"Kontrak PKWT 12 bulan untuk jabatan Junior Graphic Designer bersifat jangka waktu tertentu dan tidak menunjukkan pekerjaan yang bersifat tetap/h持续性; tidak ada klausa yang melanggar Pasal 4 ayat (2) PP 35/2021."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027","deskripsi":"Kontrak PKWT berdasarkan jangka waktu selama 12 bulan, tidak melebihi batas 5 tahun termasuk perpanjangan, sehingga patuh terhadap PP 35/2021 Pasal 8."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 8","kutipan":"PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021","deskripsi":"Klausul uang kompensasi merujuk pada PP 35/2021 dan tidak mengatur besaran yang menyimpang, sehingga tidak melanggar ketentuan."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 8","kutipan":"Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.","deskripsi":"Klausua kompensasi merujuk pada PP 35/2021 dan tidak mencabut kewajiban kompensasi bila hubungan kerja diakhiri lebih awal, sehingga tidak melanggar Pasal 17 PP 35/2021."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.","deskripsi":"Pasal 5 secara eksplisit menetapkan batas lembur 4 jam/hari dan 18 jam/minggu serta merujuk pada ketentuan peraturan perundang-undangan untuk upah lembur, sehingga sesuai dengan PP 35/2021 Pasal 26 dan 31."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"Pasal 4","kutipan":"Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.","deskripsi":"Pasal 4 telah memberikan 2 hari istirahat mingguan (Sabtu dan Minggu) untuk pola 5 hari kerja, sesuai putusan MK 168/PUU-XXI/2023."}
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

Nomor: 022/HR-PKWT/IV/2026

Pada hari ini, Rabu, tanggal 1 April 2026, bertempat di Bandung, yang bertanda tangan di bawah ini:

1. Nama: Siti Marlina; Jabatan: Manajer Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Cakrawala Kreatif Indonesia**, berkedudukan di Jalan Asia Afrika Nomor 88, Bandung, yang bergerak dalam bidang usaha jasa desain dan periklanan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Bagas Nugroho; Jenis Kelamin: Laki-laki; Umur: 24 tahun; Alamat: Jalan Cihampelas Nomor 45, Bandung, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 1 April 2026 sampai dengan 31 Maret 2027. Perjanjian ini tidak mensyaratkan masa percobaan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Graphic Designer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Asia Afrika Nomor 88, Bandung.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.200.000,00 (lima juta dua ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 1 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 08.00 sampai 16.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 5 — Kerja Lembur
Apabila diperlukan kerja lembur, pelaksanaannya paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu, dan PIHAK KEDUA berhak atas upah lembur sesuai ketentuan peraturan perundang-undangan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu, mengikutsertakan PIHAK KEDUA dalam jaminan sosial, dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan.

## Pasal 8 — Uang Kompensasi
Pada saat berakhirnya PKWT ini, PIHAK PERTAMA memberikan uang kompensasi kepada PIHAK KEDUA yang besarnya dihitung sesuai ketentuan Peraturan Pemerintah Nomor 35 Tahun 2021.

## Pasal 9 — Pencatatan
PIHAK PERTAMA mencatatkan Perjanjian ini secara daring paling lama 3 (tiga) hari kerja sejak ditandatangani.

## Pasal 10 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 11 — Penutup
Perjanjian ini dibuat dan ditandatangani di Bandung pada tanggal 1 April 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Siti Marlina                          Bagas Nugroho
```

_Response:_

```
{"melanggar":false,"section":"","kutipan":"","deskripsi":"Kontrak tidak memuat klausa apa pun yang menahan atau mensyaratkan ijazah maupun dokumen pribadi pekerja sebagai jaminan bekerja, sehingga tidak melanggar ketentuan dalam Surat Edaran Menaker M/5/HK.04.00/V/2025."}
```

## Findings kept (0)


## Findings rejected by citation gate (0)


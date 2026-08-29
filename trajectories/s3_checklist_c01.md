# Trajectory — s3_checklist — c01

## checklist:PP35-12

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 12 — PKWT dilarang mensyaratkan masa percobaan: PKWT tidak dapat mensyaratkan adanya masa percobaan kerja. Apabila masa percobaan tetap disyaratkan dalam PKWT, masa percobaan tersebut batal demi hukum dan masa kerja pekerja tetap dihitung sejak awal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":true,"section":"Pasal 4","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.","deskripsi":"Pasal 4 mensyaratkan masa percobaan kerja selama 3 bulan dalam PKWT, yang dilarang oleh PP 35/2021 Pasal 12. Klausa ini batal demi hukum dan masa kerja pekerja dihitung sejak awal."}
```

## checklist:UU13-60-1

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (1) — PKWTT masa percobaan paling lama 3 bulan: PKWTT dapat mensyaratkan masa percobaan kerja paling lama 3 bulan. Ketentuan yang menetapkan masa percobaan lebih dari 3 bulan tidak berkekuatan hukum untuk kelebihannya, dan kelebihan waktu itu dihitung sebagai masa kerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":false,"section":"Pasal 4","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.","deskripsi":"Masa percobaan ditetapkan tepat 3 bulan, sesuai batas maksimal yang diperbolehkan Pasal 60 ayat (1) UU 13/2003."}
```

## checklist:UU13-60-2

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (2) — Upah masa percobaan tidak boleh di bawah upah minimum: Selama masa percobaan kerja, pengusaha dilarang membayar upah di bawah upah minimum yang berlaku di wilayah tempat pekerja bekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":false,"section":"Pasal 3","kutipan":"PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan","deskripsi":"Upah pokok Rp6.500.000 berada di atas UMP DKI Jakarta 2026 (Rp5.067.381), sehingga tidak melanggar Pasal 60 ayat (2) UU 13/2003. Kontrak juga tidak mengatur penurunan upah selama masa percobaan."}
```

## checklist:PP35-4-2

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 4 ayat (2) — PKWT tidak untuk pekerjaan bersifat tetap: PKWT didasarkan pada jangka waktu atau selesainya suatu pekerjaan tertentu, dan tidak dapat diadakan untuk pekerjaan yang bersifat tetap. Penilaian ini bergantung pada sifat pekerjaan, bukan semata-mata bunyi klausa.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":false,"section":"Pasal 1","kutipan":"Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.","deskripsi":"Pasal 1 menyatakan PKWT berdasarkan jangka waktu 12 bulan untuk jabatan Junior Software Engineer. Pasal 2 hanya menetapkan jabatan dan tempat kerja tanpa indikator pekerjaan tetap (tidak ada klausul yang secara eksplisit menugaskan pekerjaan permanen/kontinu). Tidak ditemukan klausul yang melanggar Pasal 4 ayat (2) PP 35/2021."}
```

## checklist:PP35-8

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 8 — PKWT jangka waktu paling lama 5 tahun termasuk perpanjangan: PKWT berdasarkan jangka waktu dibuat paling lama 5 tahun, dan jangka waktu keseluruhan termasuk perpanjangan tidak boleh lebih dari 5 tahun. Batas ini hanya berlaku untuk PKWT berdasarkan jangka waktu (Pasal 5 ayat 1); PKWT berdasarkan selesainya suatu pekerjaan diatur Pasal 9 dan tidak memiliki plafon 5 tahun. Ditegaskan pula dalam Putusan MK 168/PUU-XXI/2023.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":true,"section":"Pasal 8","kutipan":"Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.","deskripsi":"Pasal 8 membatasi pemberian uang kompensasi hanya jika PKWT diperpanjang. Padahal sesuai PP 35/2021 jo. UU Cipta Kerja, pekerja PKWT berhak atas uang kompensasi pada saat berakhirnya hubungan kerja (walaupun tidak diperpanjang), sehingga klausul ini melanggar hak normatif pekerja."}
```

## checklist:PP35-16

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 16 — Besaran uang kompensasi PKWT: Besaran uang kompensasi: PKWT selama 12 bulan terus-menerus sebesar 1 bulan upah; PKWT 1 bulan atau lebih tetapi kurang dari 12 bulan dihitung secara proporsional dengan rumus masa kerja dibagi 12 dikali 1 bulan upah; PKWT lebih dari 12 bulan juga dihitung proporsional. Pada usaha mikro dan usaha kecil, besaran uang kompensasi ditentukan berdasarkan kesepakatan pengusaha dan pekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":true,"section":"Pasal 8","kutipan":"Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.","deskripsi":"Klausul ini melanggar PP 35/2021 Pasal 16 karena mensyaratkan uang kompensasi 1 bulan upah hanya diberikan jika PKWT diperpanjang. Padahal pekerja PKWT 12 bulan berhak atas uang kompensasi 1 bulan upah meskipun perjanjian berakhir tanpa perpanjangan."}
```

## checklist:PP35-17

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 17 — Kompensasi tetap wajib bila hubungan kerja diakhiri lebih awal: Apabila salah satu pihak mengakhiri hubungan kerja sebelum berakhirnya jangka waktu PKWT, pengusaha tetap wajib memberikan uang kompensasi yang besarannya dihitung berdasarkan jangka waktu PKWT yang telah dijalani oleh pekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":true,"section":"Pasal 8","kutipan":"Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.","deskripsi":"Klausul ini melanggar PP 35/2021 Pasal 17 karena menggantungkan hak kompensasi pada perpanjangan PKWT, padahal pekerja yang telah menjalani PKWT berhak atas uang kompensasi tanpa syarat apakah perjanjian diperpanjang atau tidak. Ketentuan Pasal 8 kontrak ini企图 menghilangkan kewajiban pengusaha membayar kompensasi yang secara hukum tetap melekat."}
```

## checklist:PP35-26-31

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 26 dan Pasal 31 — Batas dan upah kerja lembur: Waktu kerja lembur paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu. Upah kerja lembur untuk jam pertama sebesar 1,5 kali upah sejam, dan untuk setiap jam berikutnya sebesar 2 kali upah sejam. Upah sejam dihitung sebesar 1/173 dari upah sebulan.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausul mengenai kerja lembur, sehingga tidak ada ketentuan yang melanggar batas waktu lembur (4 jam/hari, 18 jam/minggu) maupun perhitungan upah lembur menurut PP 35/2021 Pasal 26 dan 31."}
```

## checklist:MK168-79-2b

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
Putusan MK 168/PUU-XXI/2023 atas Pasal 79 ayat (2) huruf b UU 13/2003 dalam Pasal 81 angka 25 UU 6/2023 — Lima hari kerja wajib memberi 2 hari istirahat mingguan: Mahkamah Konstitusi menyatakan norma tersebut inkonstitusional sepanjang tidak dimaknai mencakup frasa istirahat mingguan 2 hari untuk 5 hari kerja dalam 1 minggu. Konsekuensinya, kontrak dengan pola 5 hari kerja dalam 1 minggu wajib memberikan 2 hari istirahat mingguan.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.","deskripsi":"Pasal 5 secara eksplisit mengatur pola 5 hari kerja dengan 2 hari istirahat mingguan (Sabtu dan Minggu), sehingga telah sesuai dengan putusan MK 168/PUU-XXI/2023."}
```

## checklist:SE-M5-2025

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
Surat Edaran Menaker M/5/HK.04.00/V/2025 tertanggal 20 Mei 2025 — Larangan menahan ijazah dan dokumen pribadi sebagai jaminan: Pemberi kerja dilarang mensyaratkan atau menahan ijazah dan dokumen pribadi pekerja sebagai jaminan untuk bekerja, mencakup antara lain sertifikat kompetensi, paspor, akta kelahiran, buku nikah, dan BPKB. Pengecualian hanya berlaku jika ijazah atau sertifikat kompetensi diperoleh melalui pendidikan dan pelatihan yang dibiayai pemberi kerja berdasarkan perjanjian kerja tertulis, dengan kewajiban pemberi kerja menjamin keamanan dokumen dan memberi ganti rugi bila dokumen rusak atau hilang. Surat Edaran bukan peraturan perundang-undangan yang mengikat dan tidak memiliki sanksi, sehingga daya paksanya terbatas; memo harus menyatakan hal ini secara eksplisit dan tidak menyebut klausa penahanan ijazah sebagai ilegal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TERTENTU (PKWT)

Nomor: 014/HR-PKWT/III/2026

Pada hari ini, Senin, tanggal 2 Maret 2026, bertempat di Jakarta Selatan, yang bertanda tangan di bawah ini:

1. Nama: Andika Pratama; Jabatan: Direktur Sumber Daya Manusia; dalam hal ini bertindak untuk dan atas nama **PT Nusantara Digital Sejahtera**, berkedudukan di Jalan Gatot Subroto Kavling 21, Jakarta Selatan, yang bergerak dalam bidang usaha pengembangan perangkat lunak. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Rani Kusuma; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Melati Nomor 7, Depok, Jawa Barat. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat untuk mengikatkan diri dalam Perjanjian Kerja Waktu Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis dan Jangka Waktu Perjanjian
Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Junior Software Engineer dan ditempatkan di kantor PIHAK PERTAMA di Jalan Gatot Subroto Kavling 21, Jakarta Selatan.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp6.500.000,00 (enam juta lima ratus ribu rupiah) per bulan, dibayarkan setiap tanggal 25 melalui transfer ke rekening bank atas nama PIHAK KEDUA.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan menyediakan lingkungan kerja yang layak. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab dan menaati peraturan perusahaan.

## Pasal 7 — Cuti
PIHAK KEDUA berhak atas cuti tahunan sesuai ketentuan peraturan perusahaan setelah bekerja sekurang-kurangnya 12 bulan.

## Pasal 8 — Uang Kompensasi
Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.

## Pasal 9 — Kerahasiaan
PIHAK KEDUA wajib menjaga kerahasiaan informasi perusahaan selama dan setelah berakhirnya hubungan kerja.

## Pasal 10 — Pengakhiran Hubungan Kerja
Hubungan kerja berakhir pada saat jangka waktu Perjanjian berakhir. Pengakhiran sebelum waktunya dilakukan sesuai ketentuan peraturan perundang-undangan.

## Pasal 11 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 12 — Penutup
Perjanjian ini dibuat dan ditandatangani di Jakarta Selatan pada tanggal 2 Maret 2026, dalam rangkap 2 (dua) bermeterai cukup, masing-masing mempunyai kekuatan hukum yang sama.

PIHAK PERTAMA,                         PIHAK KEDUA,

Andika Pratama                         Rani Kusuma
```

_Response:_

```
{"melanggar":false,"section":"","kutipan":"","deskripsi":"Seluruh 12 pasal kontrak tidak memuat klausa penahanan ijazah atau dokumen pribadi pekerja sebagai jaminan. Tidak ditemukan pelanggaran terhadap SE Menaker M/5/HK.04.00/V/2025."}
```

## Findings kept (4)

- PP35-12 @ Pasal 4
- PP35-8 @ Pasal 8
- PP35-16 @ Pasal 8
- PP35-17 @ Pasal 8

## Findings rejected by citation gate (0)


# Trajectory — s5_gated — c01

## checklist:PP35-12

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 12 — PKWT dilarang mensyaratkan masa percobaan: PKWT tidak dapat mensyaratkan adanya masa percobaan kerja. Apabila masa percobaan tetap disyaratkan dalam PKWT, masa percobaan tersebut batal demi hukum dan masa kerja pekerja tetap dihitung sejak awal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

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
{"melanggar":true,"section":"Pasal 4","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.","deskripsi":"Klausul masa percobaan 3 bulan dalam PKWT melanggar PP 35/2021 Pasal 12 karena PKWT dilarang mensyaratkan masa percobaan. Klausul ini batal demi hukum dan masa kerja dihitung sejak awal."}
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
{"melanggar":false,"section":"Pasal 4","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja.","deskripsi":"Masa percobaan ditetapkan 3 bulan, sesuai batas maksimal Pasal 60 ayat (1) UU 13/2003. Tidak ada pelanggaran."}
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
{"melanggar":false,"section":"Pasal 4","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan terhitung sejak tanggal mulai bekerja. Selama masa percobaan, PIHAK PERTAMA dapat mengakhiri hubungan kerja sewaktu-waktu tanpa kewajiban membayar ganti rugi apa pun.","deskripsi":"Pasal 4 tidak mengatur besaran upah selama masa percobaan, melainkan hanya mengatur jangka waktu masa percobaan dan hak PHK. Upah di Pasal 3 ditetapkan Rp6.500.000 per bulan tanpa membedakan masa percobaan, sehingga tidak ada klausa spesifik yang menetapkan upah masa percobaan di bawah upah minimum."}
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
{"melanggar":false,"section":"Pasal 1","kutipan":"Perjanjian ini merupakan PKWT berdasarkan jangka waktu, berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027.","deskripsi":"Kontrak PKWT berdasarkan jangka waktu 12 bulan. Jabatan Junior Software Engineer di perusahaan pengembangan perangkat lunak bukan merupakan pekerjaan yang secara tegas bersifat tetap menurut batasan Pasal 4 ayat (2) PP 35/2021, sehingga tidak ada klausa yang jelas melanggar ketentuan."}
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
{"melanggar":false,"section":"Pasal 1","kutipan":"berlaku selama 12 (dua belas) bulan, terhitung sejak 2 Maret 2026 sampai dengan 1 Maret 2027","deskripsi":"Kontrak PKWT berdasarkan jangka waktu selama 12 bulan, tidak melanggar Pasal 8 PP 35/2021 karena masih di bawah batas 5 tahun. Tidak ada klausul perpanjangan yang melebihi plafon."}
```

## checklist:PP35-16

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 16 — Besaran uang kompensasi PKWT: Besaran uang kompensasi: PKWT selama 12 bulan terus-menerus sebesar 1 bulan upah; PKWT 1 bulan atau lebih tetapi kurang dari 12 bulan dihitung secara proporsional dengan rumus masa kerja dibagi 12 dikali 1 bulan upah; PKWT lebih dari 12 bulan juga dihitung proporsional. Pada usaha mikro dan usaha kecil, besaran uang kompensasi ditentukan berdasarkan kesepakatan pengusaha dan pekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

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
{"melanggar":true,"section":"Pasal 8","kutipan":"Uang kompensasi... hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.","deskripsi":"Klausul ini melanggar PP 35/2021 Pasal 16 karena mensyaratkan perpanjangan perjanjian sebagai syarat pemberian uang kompensasi, padahal pekerja PKWT yang telah bekerja 12 bulan berhak atas uang kompensasi 1 bulan upah regardless of perpanjangan."}
```

## checklist:PP35-17

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 17 — Kompensasi tetap wajib bila hubungan kerja diakhiri lebih awal: Apabila salah satu pihak mengakhiri hubungan kerja sebelum berakhirnya jangka waktu PKWT, pengusaha tetap wajib memberikan uang kompensasi yang besarannya dihitung berdasarkan jangka waktu PKWT yang telah dijalani oleh pekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

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
{"melanggar":true,"section":"Pasal 8","kutipan":"Uang kompensasi sebagaimana diatur dalam peraturan perundang-undangan hanya akan diberikan kepada PIHAK KEDUA apabila Perjanjian ini diperpanjang untuk periode berikutnya. Apabila Perjanjian berakhir tanpa perpanjangan, PIHAK KEDUA tidak berhak atas uang kompensasi.","deskripsi":"Pasal 8 membatasi pembayaran uang kompensasi hanya jika PKWT diperpanjang, sehingga menghilangkan hak pekerja atas kompensasi ketika perjanjian berakhir tanpa perpanjangan maupun ketika diakhiri lebih awal, yang melanggar PP 35/2021 Pasal 17."}
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
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausa mengenai kerja lembur, sehingga tidak ada ketentuan yang secara spesifik melanggar PP 35/2021 Pasal 26 dan Pasal 31."}
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
{"melanggar":false,"section":"Pasal 5","kutipan":"Waktu kerja adalah 5 (lima) hari kerja dalam seminggu, Senin sampai Jumat, pukul 09.00 sampai 17.00, dengan 2 (dua) hari istirahat mingguan pada Sabtu dan Minggu.","deskripsi":"Pasal 5 sudah patuh: pola 5 hari kerja diberikan 2 hari istirahat mingguan (Sabtu dan Minggu), sesuai Putusan MK 168/PUU-XXI/2023."}
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
{"melanggar":false,"section":"","kutipan":"","deskripsi":"Kontrak tidak memuat klausa penahanan ijazah atau dokumen pribadi pekerja sebagai jaminan, sehingga tidak melanggar ketentuan Surat Edaran Menaker M/5/HK.04.00/V/2025."}
```

## absence:PP35-13

**Outcome:** 1 dari 9 item wajib tidak ada

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
{"missing":[6]}
```

## absence:PP35-15

**Outcome:** TEMUAN: ketentuan tidak dipenuhi

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 15 — Uang kompensasi wajib pada berakhirnya PKWT: Pengusaha wajib memberikan uang kompensasi kepada pekerja pada saat berakhirnya PKWT, bagi pekerja yang telah mempunyai masa kerja paling sedikit 1 bulan secara terus-menerus. Kewajiban ini tidak berlaku bagi tenaga kerja asing.

Apakah kontrak berikut MEMENUHI ketentuan di atas? Klausa yang membuat kewajiban menjadi bersyarat atau menghapusnya dihitung sebagai TIDAK memenuhi. Balas HANYA JSON: {"memenuhi":true|false,"section":"Pasal N atau ABSENT","deskripsi":"penjelasan singkat"}.

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
{"memenuhi":false,"section":"Pasal 15","deskripsi":"Pasal 8 kontrak mensyaratkan uang kompensasi hanya diberikan jika PKWT diperpanjang, sehingga menghapus kewajiban pengusaha memberikan uang kompensasi pada berakhirnya PKWT bagi pekerja dengan masa kerja ≥1 bulan."}
```

## absence:PP35-27-5

**Outcome:** tidak ada pengecualian lembur / tidak ada temuan

_Prompt:_

```
Periksa ketentuan lembur dalam kontrak berikut. Apakah kontrak MEMUAT klausa yang menyatakan suatu jabatan dikecualikan atau TIDAK BERHAK atas upah lembur? Jika kontrak TIDAK memuat pernyataan pengecualian lembur seperti itu, jawab {"melanggar":false}. Jika kontrak MEMUAT pernyataan bahwa jabatan tidak berhak lembur TETAPI tidak mendefinisikan golongan jabatan tertentu (pemikir, perencana, pelaksana, atau pengendali jalannya perusahaan), jawab {"melanggar":true,"section":"Pasal N","kutipan":"kutipan harfiah"}.

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
{"melanggar":false}
```

## Findings kept (4)

- PP35-12 @ Pasal 4
- PP35-17 @ Pasal 8
- PP35-13 @ ABSENT
- PP35-15 @ ABSENT

## Findings rejected by citation gate (1)

- PP35-16 @ Pasal 8 — kutipan tidak ditemukan dalam kontrak

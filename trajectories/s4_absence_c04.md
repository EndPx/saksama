# Trajectory — s4_absence — c04

## checklist:UU13-60-1

**Outcome:** TEMUAN: melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (1) — PKWTT masa percobaan paling lama 3 bulan: PKWTT dapat mensyaratkan masa percobaan kerja paling lama 3 bulan. Ketentuan yang menetapkan masa percobaan lebih dari 3 bulan tidak berkekuatan hukum untuk kelebihannya, dan kelebihan waktu itu dihitung sebagai masa kerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TIDAK TERTENTU (PKWTT)

Nomor: 031/HR-PKWTT/V/2026

Pada hari ini, Senin, tanggal 4 Mei 2026, bertempat di Semarang, yang bertanda tangan di bawah ini:

1. Nama: Gunawan Saputra; Jabatan: Kepala Divisi SDM; dalam hal ini bertindak untuk dan atas nama **PT Sinar Logistik Nusantara**, berkedudukan di Jalan Pemuda Nomor 55, Semarang, yang bergerak dalam bidang usaha jasa logistik dan pergudangan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Putri Handayani; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Pandanaran Nomor 12, Semarang, Jawa Tengah. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tidak Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis Perjanjian
Perjanjian ini merupakan PKWTT (hubungan kerja tetap) yang berlaku sejak 4 Mei 2026 untuk waktu yang tidak ditentukan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Operasional Gudang dan ditempatkan di gudang PIHAK PERTAMA di Semarang.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 30 melalui transfer bank.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja. Pengangkatan sebagai pekerja tetap dilakukan setelah masa percobaan tersebut dinilai berhasil.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 8 — Penutup
Perjanjian ini dibuat dan ditandatangani di Semarang pada tanggal 4 Mei 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Gunawan Saputra                       Putri Handayani
```

_Response:_

```
{"melanggar":true,"section":"Pasal 4","kutipan":"PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja.","deskripsi":"Masa percobaan 6 bulan melanggar Pasal 60 ayat (1) UU 13/2003 karena melebihi batas maksimal 3 bulan. Kelebihannya tidak berkekuatan hukum dan dihitung sebagai masa kerja."}
```

## checklist:UU13-60-2

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
UU 13/2003 Pasal 60 ayat (2) — Upah masa percobaan tidak boleh di bawah upah minimum: Selama masa percobaan kerja, pengusaha dilarang membayar upah di bawah upah minimum yang berlaku di wilayah tempat pekerja bekerja.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TIDAK TERTENTU (PKWTT)

Nomor: 031/HR-PKWTT/V/2026

Pada hari ini, Senin, tanggal 4 Mei 2026, bertempat di Semarang, yang bertanda tangan di bawah ini:

1. Nama: Gunawan Saputra; Jabatan: Kepala Divisi SDM; dalam hal ini bertindak untuk dan atas nama **PT Sinar Logistik Nusantara**, berkedudukan di Jalan Pemuda Nomor 55, Semarang, yang bergerak dalam bidang usaha jasa logistik dan pergudangan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Putri Handayani; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Pandanaran Nomor 12, Semarang, Jawa Tengah. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tidak Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis Perjanjian
Perjanjian ini merupakan PKWTT (hubungan kerja tetap) yang berlaku sejak 4 Mei 2026 untuk waktu yang tidak ditentukan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Operasional Gudang dan ditempatkan di gudang PIHAK PERTAMA di Semarang.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 30 melalui transfer bank.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja. Pengangkatan sebagai pekerja tetap dilakukan setelah masa percobaan tersebut dinilai berhasil.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 8 — Penutup
Perjanjian ini dibuat dan ditandatangani di Semarang pada tanggal 4 Mei 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Gunawan Saputra                       Putri Handayani
```

_Response:_

```
{"melanggar":false,"section":"Pasal 3","kutipan":"PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku","deskripsi":"Klausal upah secara eksplisit menyatakan bahwa gaji Rp5.500.000 berada di atas upah minimum yang berlaku, sehingga tidak melanggar Pasal 60 ayat (2) UU 13/2003."}
```

## checklist:PP35-26-31

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
PP 35/2021 Pasal 26 dan Pasal 31 — Batas dan upah kerja lembur: Waktu kerja lembur paling lama 4 jam dalam 1 hari dan 18 jam dalam 1 minggu. Upah kerja lembur untuk jam pertama sebesar 1,5 kali upah sejam, dan untuk setiap jam berikutnya sebesar 2 kali upah sejam. Upah sejam dihitung sebesar 1/173 dari upah sebulan.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TIDAK TERTENTU (PKWTT)

Nomor: 031/HR-PKWTT/V/2026

Pada hari ini, Senin, tanggal 4 Mei 2026, bertempat di Semarang, yang bertanda tangan di bawah ini:

1. Nama: Gunawan Saputra; Jabatan: Kepala Divisi SDM; dalam hal ini bertindak untuk dan atas nama **PT Sinar Logistik Nusantara**, berkedudukan di Jalan Pemuda Nomor 55, Semarang, yang bergerak dalam bidang usaha jasa logistik dan pergudangan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Putri Handayani; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Pandanaran Nomor 12, Semarang, Jawa Tengah. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tidak Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis Perjanjian
Perjanjian ini merupakan PKWTT (hubungan kerja tetap) yang berlaku sejak 4 Mei 2026 untuk waktu yang tidak ditentukan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Operasional Gudang dan ditempatkan di gudang PIHAK PERTAMA di Semarang.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 30 melalui transfer bank.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja. Pengangkatan sebagai pekerja tetap dilakukan setelah masa percobaan tersebut dinilai berhasil.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 8 — Penutup
Perjanjian ini dibuat dan ditandatangani di Semarang pada tanggal 4 Mei 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Gunawan Saputra                       Putri Handayani
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausa spesifik tentang kerja lembur, batas waktu lembur, maupun perhitungan upah lembur. Pasal 5 hanya mengatur waktu kerja reguler (5 hari kerja, 2 hari istirahat) tanpa mengatur lembur, sehingga tidak ada pelanggaran langsung terhadap PP 35/2021 Pasal 26 dan 31."}
```

## checklist:MK168-79-2b

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
Putusan MK 168/PUU-XXI/2023 atas Pasal 79 ayat (2) huruf b UU 13/2003 dalam Pasal 81 angka 25 UU 6/2023 — Lima hari kerja wajib memberi 2 hari istirahat mingguan: Mahkamah Konstitusi menyatakan norma tersebut inkonstitusional sepanjang tidak dimaknai mencakup frasa istirahat mingguan 2 hari untuk 5 hari kerja dalam 1 minggu. Konsekuensinya, kontrak dengan pola 5 hari kerja dalam 1 minggu wajib memberikan 2 hari istirahat mingguan.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TIDAK TERTENTU (PKWTT)

Nomor: 031/HR-PKWTT/V/2026

Pada hari ini, Senin, tanggal 4 Mei 2026, bertempat di Semarang, yang bertanda tangan di bawah ini:

1. Nama: Gunawan Saputra; Jabatan: Kepala Divisi SDM; dalam hal ini bertindak untuk dan atas nama **PT Sinar Logistik Nusantara**, berkedudukan di Jalan Pemuda Nomor 55, Semarang, yang bergerak dalam bidang usaha jasa logistik dan pergudangan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Putri Handayani; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Pandanaran Nomor 12, Semarang, Jawa Tengah. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tidak Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis Perjanjian
Perjanjian ini merupakan PKWTT (hubungan kerja tetap) yang berlaku sejak 4 Mei 2026 untuk waktu yang tidak ditentukan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Operasional Gudang dan ditempatkan di gudang PIHAK PERTAMA di Semarang.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 30 melalui transfer bank.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja. Pengangkatan sebagai pekerja tetap dilakukan setelah masa percobaan tersebut dinilai berhasil.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 8 — Penutup
Perjanjian ini dibuat dan ditandatangani di Semarang pada tanggal 4 Mei 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Gunawan Saputra                       Putri Handayani
```

_Response:_

```
{"melanggar":false,"section":"Pasal 5","kutipan":"Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.","deskripsi":"Klausul waktu kerja secara eksplisit menyatakan 5 hari kerja dalam seminggu dengan 2 hari istirahat mingguan, sehingga telah sesuai dengan putusan MK 168/PUU-XXI/2023."}
```

## checklist:SE-M5-2025

**Outcome:** tidak melanggar

_Prompt:_

```
Ketentuan hukum:
Surat Edaran Menaker M/5/HK.04.00/V/2025 tertanggal 20 Mei 2025 — Larangan menahan ijazah dan dokumen pribadi sebagai jaminan: Pemberi kerja dilarang mensyaratkan atau menahan ijazah dan dokumen pribadi pekerja sebagai jaminan untuk bekerja, mencakup antara lain sertifikat kompetensi, paspor, akta kelahiran, buku nikah, dan BPKB. Pengecualian hanya berlaku jika ijazah atau sertifikat kompetensi diperoleh melalui pendidikan dan pelatihan yang dibiayai pemberi kerja berdasarkan perjanjian kerja tertulis, dengan kewajiban pemberi kerja menjamin keamanan dokumen dan memberi ganti rugi bila dokumen rusak atau hilang. Surat Edaran bukan peraturan perundang-undangan yang mengikat dan tidak memiliki sanksi, sehingga daya paksanya terbatas; memo harus menyatakan hal ini secara eksplisit dan tidak menyebut klausa penahanan ijazah sebagai ilegal.

Apakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. Balas HANYA JSON: {"melanggar":true|false,"section":"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)","kutipan":"kutipan harfiah maks 200 karakter","deskripsi":"penjelasan singkat"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TIDAK TERTENTU (PKWTT)

Nomor: 031/HR-PKWTT/V/2026

Pada hari ini, Senin, tanggal 4 Mei 2026, bertempat di Semarang, yang bertanda tangan di bawah ini:

1. Nama: Gunawan Saputra; Jabatan: Kepala Divisi SDM; dalam hal ini bertindak untuk dan atas nama **PT Sinar Logistik Nusantara**, berkedudukan di Jalan Pemuda Nomor 55, Semarang, yang bergerak dalam bidang usaha jasa logistik dan pergudangan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Putri Handayani; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Pandanaran Nomor 12, Semarang, Jawa Tengah. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tidak Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis Perjanjian
Perjanjian ini merupakan PKWTT (hubungan kerja tetap) yang berlaku sejak 4 Mei 2026 untuk waktu yang tidak ditentukan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Operasional Gudang dan ditempatkan di gudang PIHAK PERTAMA di Semarang.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 30 melalui transfer bank.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja. Pengangkatan sebagai pekerja tetap dilakukan setelah masa percobaan tersebut dinilai berhasil.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 8 — Penutup
Perjanjian ini dibuat dan ditandatangani di Semarang pada tanggal 4 Mei 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Gunawan Saputra                       Putri Handayani
```

_Response:_

```
{"melanggar":false,"section":"-","kutipan":"-","deskripsi":"Kontrak tidak memuat klausa tentang penahanan ijazah atau dokumen pribadi pekerja sebagai jaminan kerja. Tidak ada pelanggaran terhadap Surat Edaran Menaker M/5/HK.04.00/V/2025."}
```

## absence:PP35-27-5

**Outcome:** tidak ada pengecualian lembur / tidak ada temuan

_Prompt:_

```
Periksa ketentuan lembur dalam kontrak berikut. Apakah kontrak MEMUAT klausa yang menyatakan suatu jabatan dikecualikan atau TIDAK BERHAK atas upah lembur? Jika kontrak TIDAK memuat pernyataan pengecualian lembur seperti itu, jawab {"melanggar":false}. Jika kontrak MEMUAT pernyataan bahwa jabatan tidak berhak lembur TETAPI tidak mendefinisikan golongan jabatan tertentu (pemikir, perencana, pelaksana, atau pengendali jalannya perusahaan), jawab {"melanggar":true,"section":"Pasal N","kutipan":"kutipan harfiah"}.

KONTRAK:
# PERJANJIAN KERJA WAKTU TIDAK TERTENTU (PKWTT)

Nomor: 031/HR-PKWTT/V/2026

Pada hari ini, Senin, tanggal 4 Mei 2026, bertempat di Semarang, yang bertanda tangan di bawah ini:

1. Nama: Gunawan Saputra; Jabatan: Kepala Divisi SDM; dalam hal ini bertindak untuk dan atas nama **PT Sinar Logistik Nusantara**, berkedudukan di Jalan Pemuda Nomor 55, Semarang, yang bergerak dalam bidang usaha jasa logistik dan pergudangan. Selanjutnya disebut **PIHAK PERTAMA**.

2. Nama: Putri Handayani; Jenis Kelamin: Perempuan; Umur: 23 tahun; Alamat: Jalan Pandanaran Nomor 12, Semarang, Jawa Tengah. Selanjutnya disebut **PIHAK KEDUA**.

Kedua belah pihak sepakat mengikatkan diri dalam Perjanjian Kerja Waktu Tidak Tertentu dengan ketentuan sebagai berikut.

## Pasal 1 — Jenis Perjanjian
Perjanjian ini merupakan PKWTT (hubungan kerja tetap) yang berlaku sejak 4 Mei 2026 untuk waktu yang tidak ditentukan.

## Pasal 2 — Jabatan dan Tempat Pekerjaan
PIHAK KEDUA diterima bekerja pada jabatan Staf Operasional Gudang dan ditempatkan di gudang PIHAK PERTAMA di Semarang.

## Pasal 3 — Upah dan Cara Pembayaran
PIHAK KEDUA berhak atas upah sebesar Rp5.500.000,00 (lima juta lima ratus ribu rupiah) per bulan, di atas upah minimum yang berlaku, dibayarkan setiap tanggal 30 melalui transfer bank.

## Pasal 4 — Masa Percobaan
PIHAK KEDUA wajib menjalani masa percobaan kerja selama 6 (enam) bulan terhitung sejak tanggal mulai bekerja. Pengangkatan sebagai pekerja tetap dilakukan setelah masa percobaan tersebut dinilai berhasil.

## Pasal 5 — Waktu Kerja
Waktu kerja adalah 5 (lima) hari kerja dalam seminggu dengan 2 (dua) hari istirahat mingguan.

## Pasal 6 — Hak dan Kewajiban Para Pihak
PIHAK PERTAMA wajib membayar upah tepat waktu dan mengikutsertakan PIHAK KEDUA dalam jaminan sosial. PIHAK KEDUA wajib melaksanakan tugas dengan penuh tanggung jawab.

## Pasal 7 — Penyelesaian Perselisihan
Setiap perselisihan diselesaikan secara musyawarah; apabila tidak tercapai, diselesaikan melalui mekanisme penyelesaian perselisihan hubungan industrial yang berlaku.

## Pasal 8 — Penutup
Perjanjian ini dibuat dan ditandatangani di Semarang pada tanggal 4 Mei 2026, dalam rangkap 2 (dua) bermeterai cukup.

PIHAK PERTAMA,                         PIHAK KEDUA,

Gunawan Saputra                       Putri Handayani
```

_Response:_

```
{"melanggar":false}
```

## Findings kept (1)

- UU13-60-1 @ Pasal 4

## Findings rejected by citation gate (0)


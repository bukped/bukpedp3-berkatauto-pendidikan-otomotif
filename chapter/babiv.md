# BAB IV - Penginstalan Aplikasi

## A.	APLIKASI PEMBANGUN YANG DIBUTUHKAN
Pada kegiatan pembuatan aplikasi ini, kita membutuhkan aplikasi yang akan digunakan sebagai pembuatan dalam aplikasi Pendidikan otomotif ini. Berikut beberapa aplikasi yang dibutuhkan.
### a.	Visual Studio Code
Visual Studio Code merupakan Aplikasi pengeditan kode yang umum digunakan oleh developer-developer sebagai aplikasi perancangan sistem. Untuk pemasangan aplikasi ini, silahkan download instalasi nya pada situs [visualstudio.com](code.visualstudio.com).

![tampilan halaman vsc](./img/tutorial/sub_vsc%20(2).png)

Klik pada “Download for Windows”. Lalu secara otomatis akan mendownload instalasinya dan kalian langsung diarahkan pada halaman dokumentasi. Silahkan di ikuti tata cara instalasi nya selagi kalian menunggu mengunduh aplikasi tersebut.

Aplikasi yang telah terinstal terlihat seperti berikut:

![tampilan ide vsc](./img/tutorial/sub_vsc%20(1).png)

Selanjutnya, kalian dapat menyambungkan akun github dengan aplikasi ini. Dengan begitu, proses commit dapat lebih mudah dan tidak memakan waktu banyak untuk memasukkan command untuk commiting.

Selain itu, kalian dapat menyambungkan akun github kalian untuk menjalankan Github Copilot yang merupakan Chat AI yang dapat memudahkan kegiatan ngoding kalian disini. Github Copilot membutuhkan biaya untuk dapat berlangganan menggunakannya, namun jika akun github kalian memiliki student membership didalamnya, Github Copilot bisa didapatkan secara gratis.


### b. MongoDB Compass

Mongo DB Compass merupakan aplikasi dari Mongo yang dapat mempermudah kita dalam menjangkau database noSQL pada Mongo. Untuk instalasi MongoDB Compass, silahkan masuk kedalam situs berikut www.mongodb.com/try/download/compass lalu cari pada bagian “MongoDB Compass Download (GUI)”.

Disitu, kalian dapat memilih versi Compass yang ingin kalian instal. Sebagai rekomendasi, silahkan gunakan 1.41.0 yang terbaru dan sudah stable.

![Versi MongoDB Compass](/img/tutorial/sub_mongo%20(1).png)

Silahkan unduh aplikasi tersebut, setelahnya, jalankan dan ikuti tahap penginstalannya.

Aplikasi yang dijalankan untuk pertama kali, terlihat seperti berikut:

![Aplikasi MongoDB Compass](./img/tutorial/sub_mongo%20(3).png)

### c.	Postman
Postman digunakan sebagai pengujian pada API yang dibuatkan dalam Google Cloud sehingga didapatkan hasil yang dapat digunakan sebagai tolak ukur pada API yang dibuat.

Untuk menginstal Postman, silahkan masuk kedalam situs berikut www.postman.com/downloads/, dan silahkan unduh langsung dengan menekan tombol "Windows 64-Bit", silahkan tunggu beberapa saat, jalankan dan silahkan ikuti tahap penginstalannya.

![Website Postman](./img/tutorial/sub_postman%20(1).png)

Setelah penginstalan, silahkan jalankan aplikasi. kalian akan diminta untuk login terlebih dahulu untuk mengakses aplikasi Postman ini. Kalian dapat membuat akun Postman menggunakan akun Google atau GitHub kalian untuk di sambungkan dengan aplikasi ini, kemudian login kan pada aplikasi.

Tampilan aplikasi kurang lebih terlihat seperti berikut:

![Aplikasi Postman](./img/tutorial/sub_postman%20(4).png)

![Aplikasi Postman 2](./img/tutorial/sub_postman%20(3).png)

### d.	Golang
Golang sebagai Bahasa pemrograman utama yang akan kita gunakan sebagai backend pada aplikasi ini. Untuk penginstalan aplikasi ini, silahkan masuk kedalam situs go.dev, kemudian pilih download untuk mengunduh aplikasi golang.

![Website Golang](./img/tutorial/sub_golang%20(4).png)

Pada tampilan berikut, silahkan pilih instalasi untuk windows, yang kemudian akan langsung di download oleh browser. Jika sudah, silahkan jalankan aplikasi instalasi golang nya, dan ikuti tahap penginstalannya. Untuk penginstalanya, pastikan kalian menginstalnya pada satu directory saja dan tidak ditempatkan dalam Program Files.

Contoh :
```
"C:\Go"
```
Selanjutnya, kalian perlu menambah environment untuk Golang. Silahkan ikuti tahap berikut:
1.	Buka Search dan Cari “Edit the System Environment Variables”
2.	Lalu buka “Environment Variables…”

![Enviroment Variables...](./img/tutorial/sub_golang%20(1).png)

3.	Pada bagian “User variables for (namauser)”. Silahkan pilih “New…”.

![New Variables](./img/tutorial/sub_golang%20(2).png)

4.	Akan muncul jendela baru untuk mengisi variables yang ingin dibuat. Isikan Nama variable dengan “GOPATH”, dan masukkan Value dengan directory.

![GOPATH Variables](./img/tutorial/sub_golang%20(3).png)

5. Save dan Restart perangkatnya.

Untuk mengetes apakah Golang terinstal dengan tepat, kita dapat mengecek instalasinya menggunakan Command Prompt (cmd). Silahkan buka aplikasi Command Promt dengan menggunakan “Search” pada windows dan tulis “cmd”. Setelah aplikasi terbuka, silahkan masukkan command “go version” dan eksekusikan. Tampilan hasilnya kurang lebih terlihat seperti berikut:

![Golang Console](./img/tutorial/golang_console.png)

### e. Git Bash

Git Bash merupakan command prompt alternatif pengganti cmd Windows yang lebih . Git Bash dilengkapi dengan beberapa command yang dapat langsung berinteraksi dengan GitHub sehjngga menjadi aplikasi yang dibutuhkan programmer untuk mengatur repository nya.

Untuk menginstalnya, kalian dapat mengunduh penginstal nya di link situs berikut www.git-scm.com/downloads, kemudian klik Download, dan pilih Windows 64-bit pada versi Standalone.

Pada saat penginstalan, pastikan kalian menyetel beberapa opsi penginstalan berikut: 

1. Centang beberapa opsi berikut:

![Git Bash Options](../img/tutorial/sub_gitbash%20(1).png)

2. Pilih Vim sebagai default editor. Jangan Khawatir, setelah ini kita masih bisa pakai VS code sebagai code Editor.

![Vim as main editor](../img/tutorial/sub_gitbash%20(2).png)

3.	Untuk line-ending conversion, pilihlah Windows Style.

![line-ending conversion](../img/tutorial/sub_gitbash%20(3).png)

4.	Opsi lain selain yang dijelaskan sebelumnya, biarkan pada opsi default nya.

Untuk menjalankan Git Bash, silahkan pilih folder yang ingin kalian gunakan sebagai source code, kemudian klik kanan untuk memunculkan context menu, dan pilih Open Git Bash.

![Git Bash in Context Menu](../img/tutorial/sub_gitbash%20(5).png)

Jika tidak ada opsi Open Git Bash, kalian perlu menjalankan regedit untuk mengutak atik registry agar memunculkan opsi nya. Kalian dapat menjalankan registry mudahnya menggunakan codingan berikut:

```
RunGitBash.reg
Copy these Code Below, save it as .reg file. Run it and accept everything.

Windows Registry Editor Version 5.00
; Open files
; Default Git-Bash Location C:\Program Files\Git\git-bash.exe

[HKEY_CLASSES_ROOT\*\shell\Open Git Bash]
@="Open Git Bash"
"Icon"="C:\\Program Files\\Git\\git-bash.exe"

[HKEY_CLASSES_ROOT\*\shell\Open Git Bash\command]
@="\"C:\\Program Files\\Git\\git-bash.exe\" \"--cd=%1\""

; This will make it appear when you right click ON a folder
; The "Icon" line can be removed if you don't want the icon to appear

[HKEY_CLASSES_ROOT\Directory\shell\bash]
@="Open Git Bash"
"Icon"="C:\\Program Files\\Git\\git-bash.exe"


[HKEY_CLASSES_ROOT\Directory\shell\bash\command]
@="\"C:\\Program Files\\Git\\git-bash.exe\" \"--cd=%1\""

; This will make it appear when you right click INSIDE a folder
; The "Icon" line can be removed if you don't want the icon to appear

[HKEY_CLASSES_ROOT\Directory\Background\shell\bash]
@="Open Git Bash"
"Icon"="C:\\Program Files\\Git\\git-bash.exe"

[HKEY_CLASSES_ROOT\Directory\Background\shell\bash\command]
@="\"C:\\Program Files\\Git\\git-bash.exe\" \"--cd=%v.\""
```

Masukkan codingan kedalam notepad. Save codingan tersebut sebagai file .reg lalu jalankan. Cek Kembali apakah Open Git Bash muncul pada opsi context menu.

Tampilan aplikasi akan terlihat seperti berikut:

![Git Bash](../img/tutorial/sub_gitbash%20(4).png)

## B.	REKOMENDASI APLIKASI PEMBANGUN LAINNYA
Selain beberapa aplikasi yang dijelaskan tadi, adapun beberapa aplikasi yang dapat kalian instal untuk digunakan dalam pembuatan aplikasi website ini:

1.	[XAMPP](https://www.apachefriends.org/download.html), digunakan untuk menjalankan website secara lokal.

2.	[GoLand](https://www.jetbrains.com/go/), aplikasi IDE seperti VSCode namun difokuskan untuk Bahasa pemrograma Golang.

3.	[Git Bash](https://www.git-scm.com/downloads), sebagai command prompt pengganti cmd Windows. Git Bash dilengkapi dengan beberapa command yang dapat langsung berinteraksi dengan GitHub sehjngga menjadi aplikasi yang dibutuhkan programmer.


Lanjut ke... [BAB V](babv.md)


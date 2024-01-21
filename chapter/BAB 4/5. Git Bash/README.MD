# e. Git Bash

Git Bash merupakan command prompt alternatif pengganti cmd Windows yang lebih . Git Bash dilengkapi dengan beberapa command yang dapat langsung berinteraksi dengan GitHub sehjngga menjadi aplikasi yang dibutuhkan programmer untuk mengatur repository nya.

Untuk menginstalnya, kalian dapat mengunduh penginstal nya di link situs berikut www.git-scm.com/downloads, kemudian klik Download, dan pilih Windows 64-bit pada versi Standalone.

Pada saat penginstalan, pastikan kalian menyetel beberapa opsi penginstalan berikut: 

1. Centang beberapa opsi berikut:

![Git Bash Options](../../../img/tutorial/sub_gitbash%20(1).png)

2. Pilih Vim sebagai default editor. Jangan Khawatir, setelah ini kita masih bisa pakai VS code sebagai code Editor.

![Vim as main editor](../../../img/tutorial/sub_gitbash%20(2).png)

3.	Untuk line-ending conversion, pilihlah Windows Style.

![line-ending conversion](../../../img/tutorial/sub_gitbash%20(3).png)

4.	Opsi lain selain yang dijelaskan sebelumnya, biarkan pada opsi default nya.

Untuk menjalankan Git Bash, silahkan pilih folder yang ingin kalian gunakan sebagai source code, kemudian klik kanan untuk memunculkan context menu, dan pilih Open Git Bash.

![Git Bash in Context Menu](../../../img/tutorial/sub_gitbash%20(5).png)

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

![Git Bash](../../../img/tutorial/sub_gitbash%20(4).png)

## B.	REKOMENDASI APLIKASI PEMBANGUN LAINNYA
Selain beberapa aplikasi yang dijelaskan tadi, adapun beberapa aplikasi yang dapat kalian instal untuk digunakan dalam pembuatan aplikasi website ini:

1.	[XAMPP](https://www.apachefriends.org/download.html), digunakan untuk menjalankan website secara lokal.

2.	[GoLand](https://www.jetbrains.com/go/), aplikasi IDE seperti VSCode namun difokuskan untuk Bahasa pemrograma Golang.
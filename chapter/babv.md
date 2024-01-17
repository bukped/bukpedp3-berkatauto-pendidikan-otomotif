# BAB V - Pembuatan Aplikasi Website

## A.	PEMBUATAN REPOSITORY GITHUB
Github merupakan situs yang menyediakan source code, kalian dapat membuat repository nya disini. Dengan GitHub, kalian dapat menyimpan berbagai repository yang akan kalian gunakan dalam tutorial ini, kalian juga dapat menjalankan aplikasi web nya disini dengan menggunakan plugin pages yang disediakan github tanpa mengeluarkan banyak biaya.

Berikut adalah tahap pembuatan Repository GitHub.

1.	Buat akun terlebih dahulu yang khusus digunakan sebagai akun aplikasi web. Kalian dapat membuat akun sendiri, atau menjadikannya sebagai akun Organization. Dalam Tutorial ini, kita hanya menggunakannya sebagai akun sendiri saja.

2.	Pada Tampilan halaman depan, lihat pada navigation bar nya, dan klik pada tombol “+”, akan muncul dropdown opsi nya dan pilih “New Repository”.

![GitHub Repository maker](../img/tutorial/sub_github%20(1).png)

3.	Tampilan berikut adalah form untuk membuat Repository, silahkan beri nama Repository yang kalian buat dan isikan deskripsinya.

![Github Repository form](../img/tutorial/sub_github%20(2).png)

Dalam hal ini, beberapa repository yang kita buat pada tutorial diperlukan License agar dapat di publikasi kan pada beberapa penyedia library pihak ketiga. Namun untuk kali ini kita tidak perlu membutuhkannya.

Untuk repository pertama, silahkan kalian beri namanya dengan (namaakun).github.io. Maka secara otomatis repository tersebut menjadi web pages. Silahkan kalian tes link nya dengan memasukkan link dengan nama yang sama seperti pada repository. Lihat hasilnya apakah dapat masuk kedalam halaman atau belum.

## B.	PEMBUATAN FRONTEND
Ada berbagai cara untuk membuat Front-end. Namun pada tutorial ini, kita akan menggunakan template sebagai Frontend. Pada tutorial ini, kita menggunakan Template sebagai Frontendnya. Kalian dapat menggunakan berbagai template yang ingin kalian gunakan, namun untuk tutorial ini, kita akan menggunakan template dari tailwind yang dapat unduh pada link berikut:

Halaman Utama:

[Download Here](uideck.com/templates/base-tailwind/)

Dashboard:

[Download Here](www.tailwindawesome.com/resources/tailwind-admin-template)

Setelah itu, kita dapat langsung memasukkannya kedalam repository yang sudah kita buat. Ikuti tahap berikut:

### a.	Memasukkan Template Kedalam repository
Ikuti tahap berikut untuk memasukkan template kedalam repository.
1.	Silahkan buat folder, beri nama sesuai keinginan untuk frontend.
2.	Silahkan buatkan folder didalamnya seperti berikut:

![Frontend Foldering](../img/tutorial/sub_frontend%20(2).png)

3.	Ekstrak kedua template frontend yang di unduh tadi, silahkan masukkan seperti pada gambar berikut.

![Frontend Foldering](../img/tutorial/sub_frontend%20(1).png)

Pastikan sesuaikan beberapa template berikut:

a.	Index.html untuk halaman utama ditempatkan pada folder utama.

b.	Untuk signin, signup, dan beberapa template halaman utama lainnya, silahkan masukkan kedalam folder pages, ganti Namanya untuk beberapa template menggunakan nama yang tertera pada gambar di atas.

c.	Folder dsb adalah tempat dimana template dashboard ditempatkan, silahkan masukkan index.html dashboard kedalam folder tersebut. Ganti namanya menjadi dashboard.html.

d.	Ada blank.html didalam template, silahkan masukkan juga untuk dibuatkan CRUD Create.

e.	Dan untuk beberapa template dashboard lainnya, silahkan masukkan dan sesuaikan namanya.

4.	Assets (CSS, Gambar dan beberapa JS) yang masuk termasuk sebagai pembangun template-template tersebut, masukkan kedalam folder-folder yang sudah dibuatkan tadi.

## C.	PEMBUATAN BACKEND
### a.	Mengenali Logika Pada Proses Backend
Sebelum memulai pembuatan Backend, ada baiknya kita mengenali terlebih dahulu alur logika pada proses backend. Dengan ini, kita mendapatkan gambaran bagaimana proses-proses pada backend yang kita buat berjalan.

1. Sign Up



2. Login



3. Tambah Artikel



4. Hapus Artikel



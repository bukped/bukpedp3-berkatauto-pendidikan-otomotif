![Cover](./img/cover.png)

# Tutorial Membuat Aplikasi Website Pendidikan Otomotif

Penulis : 

Adam Ghafara
\
Rachma Nurhaliza Parindra
\
Roni Habibi

## PRAKARTA

Otomotif merupakan salah satu topik menarik bagi para pecinta mobil, motor, dan kendaraan bermesin lainnya. Selain itu, Otomotif menjadi bahan topik bagi mereka yang ingin mengenal lebih tentangnya, baik dalam perawatan kendaraan, tips dan trik, dan juga informasi terkini. Namun, informasi pada media internet yang membahas tentang otomotif terbilang sedikit dan sangat terbatas. Hal ini membuat beberapa dari mereka kesulitan mencari informasi tepat dan ter aktual dalam otomotif.

Buku ini bertujuan untuk memberikan pemahaman tentang pembuatan website berteknologi Golang dan JavaScript untuk mendukung pemahaman yang lebih mendalam tentang teknologi tersebut. Dengan menggunakan Golang sebagai bahasa pemrograman backend, kalian dapat menciptakan aplikasi dasar dengan kinerja yang tinggi, kehandalan, dan kemampuan untuk mengelola beban kerja dengan efisien. Meskipun demikian, penggunaan JavaScript pada sisi frontend memungkinkan pembuatan antarmuka pengguna yang dinamis dan interaktif. Kombinasi kedua teknologi ini menjadi landasan untuk membangun pengalaman belajar yang kuat, yang menggabungkan komponen fungsionalitas backend yang canggih dengan responsivitas antarmuka pengguna yang menarik.

Oleh karena itu, Tutorial ini bertujuan untuk memberikan wawasan mendalam tentang pembuatan website dengan mengambil tema Otomotif sebagai bahan tutorial nya. Selain itu, Tutorial ini juga bertujuan untuk mendorong inovasi dalam penggunaan teknologi dalam pendidikan.

Buku ini juga memiliki source code sebagai bahan ajar yang dapat kalian akses menggunakan link github berikut:
https://github.com/berkatauto

## Daftar Isi

### [PRAKATA](#prakarta)
### [BAB I - Pendahuluan](#bab-i---pendahuluan)
### [BAB II - Pengenalan Otomotif](#bab-ii---pengenalan-otomotif)
### [BAB III - Dasar Pemrograman](#bab-iii---dasar-pemrograman)
### [BAB IV - Penginstalan Aplikasi](./chapter/babiv.md)
### [BAB V - Pembuatan Aplikasi Website](./chapter/babv.md)

## BAB II - PENGENALAN OTOMOTIF

### A.	KONSEP DASAR OTOMOTIF

Otomotif mencakup serangkaian konsep dasar yang penting untuk memahami cara kerja kendaraan bermotor. Salah satunya adalah mesin, jantung dari mobil. Mesin menggunakan energi dari bahan bakar untuk menciptakan gerakan mekanis yang membantu kendaraan bergerak. Selain itu, terdapat sistem transmisi yang mentransfer tenaga dari mesin ke roda kendaraan dengan mengatur kecepatan dan torsi. Sistem pembakaran internal mesin bekerja dengan membakar campuran bahan bakar dan udara untuk menghasilkan energi. Komponen lainnya antara lain transmisi, suspensi, rem, bahan bakar, pendingin, kelistrikan, kemudi, intake, dan sistem transmisi yang semuanya bekerja sama untuk membantu kendaraan tampil maksimal. Memahami konsep-konsep ini membantu merawat, memperbaiki, dan mengoperasikan kendaraan dengan lebih efektif.

### B.	PERKEMBANGAN OTOMOTIF SAAT KINI

Saat ini, industri otomotif sedang mengalami transformasi yang signifikan. Perkembangan teknologi telah melahirkan inovasi-inovasi besar di bidang kendaraan bermotor. Mobilitas  listrik semakin menarik perhatian masyarakat seiring dengan semakin banyaknya jumlah mobil listrik yang diproduksi oleh banyak pabrikan besar. Selain itu, konsep kendaraan otonom atau self-driving juga semakin mendapat perhatian, dengan dikembangkannya sistem kecerdasan buatan yang memungkinkan kendaraan mengemudi sendiri melalui sensor canggih dan teknologi pemrosesan data. Aspek keselamatan juga ditekankan dengan penerapan teknologi pencegahan kecelakaan, seperti sistem pengereman darurat otomatis dan sistem bantuan pengemudi yang semakin canggih. Selain itu, konsep berbagi mobil dan transportasi berbasis layanan (seperti persewaan mobil jangka pendek melalui aplikasi) semakin populer, sehingga mengubah cara masyarakat menggunakan dan memiliki kendaraan. Hal ini merupakan bagian dari perubahan besar dalam industri otomotif yang kini bertujuan untuk meningkatkan kenyamanan, efisiensi, keselamatan, dan keberlanjutan.

### C.	KEUNGGULAN DARI WEBSITE PENDIDIKAN OTOMOTIF

Website pendidikan otomotif memiliki manfaat yang signifikan dalam menyediakan akses universal terhadap pengetahuan otomotif. Dengan kemudahan akses melalui Internet, website ini tidak hanya menyajikan materi pembelajaran terstruktur tetapi juga memungkinkan pembaruan konten secara cepat, mencakup inovasi dan perkembangan terkini di industri mobil. Interaktivitas konten, seperti video instruksional dan simulasi, memungkinkan pemahaman yang lebih baik, sementara platform online juga mendukung forum diskusi, kolaborasi, dan pengujian yang memperkaya pengalaman belajar. Fleksibilitas untuk belajar kapan saja, di mana saja, serta penyesuaiannya, menjadikan situs web pendidikan otomotif sebagai sumber daya yang sangat berharga bagi pelajar, profesional, dan penggemar otomotif.

1.	Aksesibilitas di seluruh dunia
Melalui situs web, siapa pun dari berbagai belahan dunia dapat mengakses informasi otomotif dan materi pendidikan, selama mereka memiliki koneksi Internet. Hal ini memungkinkan penyebaran pengetahuan otomotif secara luas.

2.	Konten interaktif
Situs web dapat menyediakan konten interaktif seperti video instruksional, simulasi, atau animasi yang membantu siswa atau pengguna lebih memahami konsep otomotif.

3.	Pembaruan yang mudah dan ulasan
Situs ini dapat diperbarui secara berkala untuk memuat informasi terkini, teknologi terkini, atau perkembangan industri otomotif. Hal ini memastikan bahwa materi yang disampaikan tetap relevan dan terkini.

## BAB III - DASAR PEMROGRAMAN

### A. BAHASA PEMROGRAMAN

Bahasa pemrograman adalah seperangkat aturan dan instruksi yang digunakan untuk berkomunikasi dengan komputer. Ini adalah cara manusia merancang algoritme, mengatur data, dan menginstruksikan komputer untuk melakukan serangkaian tindakan tertentu. Setiap bahasa pemrograman memiliki sintaks dan aturannya sendiri yang memungkinkan pemrogram mengekspresikan ide dan tujuan mereka dalam format yang dapat dipahami mesin. 

Dari bahasa tingkat rendah seperti bahasa mesin dan perakitan hingga bahasa tingkat tinggi seperti Python, JavaScript, dan C++, masing-masing bahasa memiliki kegunaannya masing-masing dalam mengembangkan berbagai aplikasi, mulai dari perangkat lunak hingga pengembangan web dan buatan. Intelijen. Bahasa pemrograman adalah landasan inovasi teknologi modern, memungkinkan kita  menciptakan solusi yang lebih kompleks dan efektif di dunia digital. Berikut ini merupakan bahasa pemrograman yang digunakan :

#### a. HTML

Menurut Bimo Sunarfrihantono (2002), HTML (Hypertext Markup Language) adalah format yang digunakan untuk membuat dokumen dan aplikasi yang berjalan di halaman web. Oleh karena itu, sebelum Anda dapat membuat program aplikasi yang berjalan di halaman web, Anda harus mengenal dan menguasai HTML. “Hypertext Merkup Languange (HTML) merupakan bahasa standar yang digunakan untuk menampilkan halaman web,” kata Hidayatullah dan Kawistara (2015) dalam jurnal Fitri Ayu dan Nia Permata Sari (ISSN:2549-0222).

Kelebihan yang dimiliki HTML, diantaranya :

1.	Bahasa yang dipakai secara dan mempunyai banyak sumber serta komunitas yang besar.
2.	Dapat dijalankan setiap web browser.
3.	Open-source yang digunakan secara gratis.
4.	Bahasa markup yang tertata rapi dan konsisten.
5.	Dapat diintegrasikan dengan bahasa pemrograman yang dapat berjalan di server side seperti PHP, Asp, Java dan lain-lain.

Adapun kekurangan yang dimiliki HTML :

1.	HTML dipakai untuk membuat halaman website statis dan sederhana.
2.	HTML tidak dapat menjalankan logic.

#### b. CSS

Cascading Style Sheets (CSS) merupakan bahasa pemrograman yang berfungsi sebagai web design. Penggunaan CSS untuk membuat tampilan web yang bagus dan menarik.

Berikut Ini merupakan kelebihan dari CSS yaitu :
1.	Memisahkan desain dengan konten situs web.
2.	Mengelola desain dengan cara efisien.
3.	Lebih mudah untuk mengunduh karena lebih ringan dalam ukuran file.
4.	CSS dapat digunakan di banyak halaman web.

Berikut kekurangan dari CSS yaitu :
1.	Beberapa browser tidak mendukung css (browser lama).
2.	Butuh waktu lebih lama untuk membuat codingan css.
3.	Masih banyaknya bug atau error dalam css.

#### c. Javascript

Bahasa pemrograman JavaScript diresmikan pada tahun 1996 dan menjadi bahasa pemrograman ECMAScript. Dua tahun kemudian tepatnya tahun 1998, bahasa ECMAScript menjadi ECMAScript 2. Dan pada tahun 1999, ECMAScript 2 menjadi ECMAScript 3. ECMAScript 3 kembali dikembangkan menjadi bahasa pemrograman JavaScript yang kita kenal selama ini. Dari tahun ke tahun, bahasa JavaScript mengalami perkembangan yang pesat. Pada tahun 2016, 92% situs web di seluruh dunia menggunakan bahasa pemrograman JavaScript. Dalam 20 tahun terakhir bahasa ini memiliki banyak fungsi, apalagi bagi para pengembang website sering kita mendengar tentang bahasa pemrograman JavaScript.
JavaScript merupakan bahasa pemrograman tingkat tinggi yang kini menjadi bahasa pemrograman utama bagi pengembang web bersama dengan HTML (Hypertext Markup Language) dan CSS (Cascading Style Sheets). JavaScript adalah bahasa pemrograman yang digunakan dalam mengembangkan situs web agar lebih dinamis dan interaktif. 

Berikut Kelebihan javascript :
1.	Fleksibilitas tinggi
JavaScript adalah bahasa pemrograman yang sangat fleksibel. Hal ini memungkinkan pengembang menggunakan metode berbeda untuk  menyelesaikan tugas yang sama.
2.	Cocok untuk pengembangan web
Karena JavaScript terintegrasi dengan baik ke dalam pengembangan web, JavaScript memungkinkan interaksi langsung dengan elemen HTML, menjadikannya pilihan yang baik untuk membuat halaman web interaktif.
3.	Berfungsi lintas platform
JavaScript tidak terbatas pada platform tertentu dan dapat berjalan di berbagai sistem operasi dan perangkat.

Berikut Kekurangan javascript :
1.	Performa
Meskipun JavaScript terus mengalami peningkatan, performanya terkadang  dapat menjadi masalah dalam aplikasi yang sangat besar atau kompleks.
2.	Ketergantungan pada koneksi Internet
Beberapa pustaka JavaScript dan fitur web modern memerlukan koneksi Internet yang stabil, yang dapat membatasi penggunaan dalam situasi di mana konektivitas Internet terbatas.
3.	Keterbatasan Akses ke Fungsi Perangkat Keras
JavaScript di browser memiliki keterbatasan dalam mengakses fungsi perangkat keras langsung (seperti membaca file dari sistem), yang mengharuskan penggunaan API yang disediakan browser.

#### d. Golang (Go)

Go merupakan bahasa pemrograman yang dikembangkan oleh Google pada tahun 2009. Bahasa ini dirancang untuk fokus pada kesederhanaan, kejelasan, dan efisiensi dalam pengembangan perangkat lunak. Go menonjol karena konkurensi dan goroutinenya yang memungkinkan eksekusi paralel tugas-tugas ringan, mendukung pengembangan perangkat lunak secara bersamaan, terukur, dan efisien. 
Keuntungan lainnya termasuk kompilasi yang cepat, dukungan untuk pemrograman berorientasi objek tanpa terlalu banyak kerumitan, dan kemampuan untuk membuat perangkat lunak yang dapat berjalan di berbagai platform dengan sedikit modifikasi. Go telah banyak digunakan dalam berbagai aplikasi termasuk pengembangan aplikasi web, server backend, pemrosesan data, dan banyak proyek teknologi tinggi lainnya.
Kelebihan dari Go yaitu :
1.	Kinerja tinggi
Go memberikan kinerja yang cepat dan efisien, terutama saat mengompilasi kode.
2.	Konkurensi tinggi
Dukungan goroutine bawaan memungkinkan eksekusi tugas paralel (paralel) dengan mudah, mendukung aplikasi yang memerlukan pemrosesan bersamaan.
3.	Manajemen memori yang baik
Go memiliki manajemen memori yang baik, mengurangi risiko kebocoran memori ingat dan membuat aplikasi lebih stabil.
Kekurangan dari Go yaitu :
1.	Kurangnya keragaman perpustakaan
Meskipun perpustakaan standarnya bagus, Go mungkin kekurangan keragaman perpustakaan dan alat dibandingkan dengan bahasa lain yang lebih mapan.
2.	Sebuah bahasa yang masih dalam pengembangan 
Sebagai bahasa yang relatif baru, ada adalah ekspektasi rendah terhadap fitur atau konvensi bahasa pemrograman yang mungkin belum tersedia atau didefinisikan sepenuhnya di Go. 
3.	Komunitas Lebih Kecil
Meskipun komunitas Go terus berkembang, ukurannya masih lebih kecil dibandingkan komunitas yang lebih mapan bahasa pemrograman. Sehingga mendapatkan bantuan atau dukungan mungkin memerlukan lebih banyak usaha.
4.	Kurangnya  fitur yang diharapkan
Beberapa fitur digunakan oleh pengembang yang akrab dengan bahasa tersebut Harapkan pemrograman lain kurang tersedia atau bahkan sedang dikembangkan di Go.

Lanjutan menuju Bab Selanjutnya... [BAB IV](./chapter/babiv.md)
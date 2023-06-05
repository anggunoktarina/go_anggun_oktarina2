# (18) System Design

*******************
Diagram Design Tools: smartdraw, lucidchart, draw.io, whimsical, visio.
desain yang umum digunakan : flowchart, Use Case, ERD, HLA

key Characteristics of Distributed Systems:
1. scalability/skalabilitas
2. reliability/keandalan
3. availability/ketersediaan
4. efficiency/efisien
5. serviceability or manageability/ kemampuan melayani dan mengelola

Horizontal Scaling vc Vertical Scaling:
Pilihan antara horizontal scaling dan vertical scaling bergantung pada kebutuhan dan karakteristik sistem yang akan diukur. Horizontal scaling cenderung memberikan skala yang lebih baik secara keseluruhan, sementara vertical scaling bisa menjadi pilihan yang lebih sederhana dan terbatas.
*******************
Load Balancing/ Penyeimbangan beban kerja

Monolithic vs Microservices
antara Monolithic dan Microservices tergantung pada kompleksitas dan skala proyek perangkat lunak. Monolithic cocok untuk aplikasi kecil atau sederhana, sementara Microservices lebih cocok untuk aplikasi yang kompleks dan memerlukan skalabilitas, fleksibilitas, dan evolusi yang lebih baik.

SQL vs NoSQL
Basis data relasional (SQL) dirancang untuk segala keperluan, memiliki standar yang jelas, memiliki banyak tool (administrasi, reporting, framework)

NoSQL adalah DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan dengan model RDBMS, Menghindari: effort pada sifat transaksi ACID, kompleksitas SQL, Design schema di depan, transactions (ditangani oleh aplikasi). Kelebihan: schema less, fast development, etc.

Kapan digunakan: membutuhkan skema fleksibel, ACID tidak diperlukan, data logging (json), data sementara (chace), etc. Kapan tidak direkomendasikan : data finansial, data transaksi, business critical, ACID - compliant.


*******************
Caching digunakan pada data yang baru saja diminta kemungkinan besar akan diminta kembali. 

Database Indexing
"Indexing adalah cara untuk mengoptimalkan kinerja sebuah basis data dengan meminimalkan jumlah akses disk yang diperlukan saat sebuah query diproses. Ini adalah teknik struktur data yang digunakan untuk dengan cepat menemukan dan mengakses data dalam sebuah basis data."
Geeksforgeeks - Indexing Database
Sebagian besar basis data menggunakan B-Tree sebagai struktur data mereka untuk indexing. Dan B-Tree memiliki kompleksitas O(log n) untuk operasi pencarian, penghapusan, dan penyisipan.

Database Replication
Redundansi adalah duplikasi komponen atau fungsi kritis dari suatu sistem dengan tujuan meningkatkan kehandalan sistem, biasanya dalam bentuk cadangan atau keamanan, atau untuk meningkatkan kinerja sistem sebenarnya

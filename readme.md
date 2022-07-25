# Get Started

# Install Go Lang di windows

# Install Golang di Linux

1.  Unduh arsip installer dari https://golang.org/dl/, pilih installer untuk Linux yang sesuai dengan jenis bit komputer anda. Proses download bisa dilakukan lewat CLI, menggunakan wget atau curl.
    $ wget https://storage.googleapis.com/golang/go1...
2.  Buka terminal, extract arsip tersebut ke /usr/local.
    $ tar -C /usr/local -xzf go1...
3.  Tambahkan path binary Go ke PATH environment variable.
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
    $ source ~/.bashrc

4.  Selanjutnya, eksekusi perintah berikut untuk mengetes apakah Go sudah terinstal dengan benar.
    $ go version
5.  Jika output adalah sama dengan versi Go yang ter-install, menandakan proses instalasi berhasil.

# Inisiasi Project

$ go mod init
digunakan untuk inisiasi project baru
contoh
$ mkdir project-golang-pertama
$ cd project-golang-pertama
$ go mod init

atau bisa menggunakan dengan scema
$ go mod init <nama-project>

# Go Command

## Menjalankan Project

$ go run main.go
$ go run <nama-file.ext> hanya di gunakan pada file yang nama package nya adalah main.

$ go run main.go library.go

## Command Test

$ go test main_test.go

## Command Build

Menghasilkan
Proses Kompilasi, menghasilkan exeucte program disipan pada folder yang sedang aktif
$ go build

File Kompilasi akan disimpkan di folder temporary
$ go run

## Command Go Get

Command go get digunakan untuk men-download package. Sebagai contoh saya ingin men-download package Kafka driver untuk Go pada projek project-pertama

$ cd project-pertama
$ go get github.com/segmentio/kafka-go
$ dir

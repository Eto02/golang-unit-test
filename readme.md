# Belajar Unit Test di Golang

Repository ini berisi contoh-contoh unit test di Golang. Unit test adalah cara untuk memverifikasi bahwa kode berfungsi sesuai dengan yang diharapkan. Dengan menulis dan menjalankan unit test, Anda dapat memastikan bahwa perubahan pada kode tidak merusak fungsionalitas yang ada.

## Struktur Proyek

- `go.mod`: Berkas ini mendefinisikan modul Go dan ketergantungannya.

## Menjalankan Unit Test

Untuk menjalankan unit test di Golang, gunakan perintah berikut di terminal:

```sh
go test -v
```

## Menjalankan Unit Test Benchmark

Untuk menjalankan unit test di Golang, gunakan perintah berikut di terminal:

```sh
go test -benchmem -run=^$ -bench=Bench
```

Section 1: Golang Dasar

1. Pengenalan Golang
2. Program Hello Word `helloworld.go`
3. Tipe Data Number `number.go` [1]
   ada banyak tipedata number di golang, dan case sensitive. !make it summarize, important
4. Tipe Data Boolean `boolean.go` [1]
5. Tipe Data String inlude function string `string.go`
6. Variable `variable.go`
   Strict 1 variable unique, dan harus dipakai tidak seperti javascript, dan php
7. Constant `constant.go`
   Nilai tidak bisa diubah, setelah diberi nilai
8. Konversi Tipe Data `conversion.go`
9. Type Declaration `type-declaration.go`
   Kemampuan membuat alias type data sesuai yang kita inginkan
10. Operasi Matematika `math-opeation.go`
    +,-,'\*',/,%
11. Operasi Perbandingan `comparation-operation.go`
12. Operasi Boolean `boolean-operation.go` !deep-understanding table && and !!
13. Tipe Data Array `array.go`
    Tipe data yang berisikan kumpulan data dengan tipe data yang sama, di go perlu menentukan jml data yang di tampung oleh array tsb.
14. Tipe Data Slice `slice.go` hanya ada di go !!deep understand.
15. Tipe Data Map seperti object `map.go`
16. If Expression `if.go`
17. Switch Expression !!deep understand `switch.go`
    // Manarik bisa menggunakan expression
18. For Loop `for.go`
    // hanya ada for di golang, tidak ada while / do while
19. Break dan Continue `break.go`
20. Function `function.go`
21. Function Parameter `function-as-parameter.go`
22. Funciton Return Value `function-return-value.go` harus di declare tipe data kembalian nya
23. Returning Multiple Value `function-return-multipe-value.go`
24. Named Return Value `function-return-named-value.go`
25. Variadic Function sample (numbers ...int) `function-variadic.go`
26. Function Value `function-as-value`
27. Function as Parameter (Fitur PowerFull) `function-as-parameter`
28. Anonymous Function `function-as-anonymous.go`
29. Recursive Function `function-recursive.go`
30. Closure `closure.go`
31. Defer, Panic, & Recover
    - Defer : function yang dapat kita jadwalkan untuk dieksekusi setelah sebuah function di execute `defer.go`
    - Panic : Function yang dapat kita gunakan untuk menghentikan program, difer function akan tetap diexecute `panic.go`
    - Recover : Function yang dapat digunakan untuk menangkap data panic, recover letakan di dalam function defer `recover.go`, proces panic tidak akan menghentikan aplikasi
32. Komentar `komentar.go`
33. Stuck `struct.go`
    - Template data yang digunakan untuk menggabungkan nol atau lebih tipedata lainnya dalam satu kesatuan
    - Struck mereprentasikan
      data dalam program aplikasi yang kita buat
34. Struct Method `struct.go`
35. Interface
    - Tipe data abstract, dan tidak memiliki implementasi langsung, berisikan definisi method2, biasanya digunakan sebagai kontrak
    - Implementasi interface berbeda dengan php, java yang harus menyebutkan explisti interface `interface.go`
36. Interface Kosong
    Interface Kosing adalah interface yang tidak memiliki deklarasi method satupun, secara otomatis semua tipe data akan menjadi implementasinya `interface-kosong.go`
37. Nil (data kosong)
    - Go lang saat bikin variable sekalipun belum diinisiasi secara otomatis akan deibuat default value
    - Nill bisa di gunakan di beberapa tipe data, interface, function, map, slice, pointer dan channel `nil.go`
38. Error Interface `interface-error.go`
39. Type Assertion `type-assertions.go`
    - Mengubah type data, biasanya return di interface
40. Pointer `pointer.go`
    Kemampuan membuat reference ke lokasi data di memory yang sama tanpa menduplikasi data yg sudah ada

    - Default Pass By Value -> yg dikirim duplikasi (tidak seperti php)
    - Pass By Reference ->
    - Operator &, untuk membuat sebuah variable dengan nilai pointer ke variable yg lain
    - Operator \*

41. Pointer di Function `ponter.go`
    - Kadang kita ingin membuat function yang dapat mengubah data asli dari parameter tsb
42. Pointer di Method
43. Aktifkan GOPATH
    - mengaktifkan GOPATH $ go env -w GO111MODULE=off
44. GOPATH - JAVA HOME in windows sample - Tambahkan GOPATH
    Mac - Add to .bashrc or .profile
    export GOPATH="/Users/user/Development/GOLANG"
    jdi smua download / library masuk ke folder itu
    $ open -e .bash_profile
    di mac saya
    export GOPATH="/Users/user/Desktop/My Course/Golang"
    dan pindahkan code ke src src/belajar-golang-dasar
45. Error Package & Import
    Di Go-Lang versi terbaru, kita sudah diwajibkan menggunakan Go-Modules
    mematikan go module $ go env -w GO111MODULE=off
46. Package & Import
    Package sebenarnya hanya directory folder
47. Access Modifier
    nama variable / function Capital Public dan lower Private.
48. Package Initialization
    Ini seperti constructor di php, function diakses sescara otomatis ketika package diakses
49. Package OS (Mencoba paket bawaan go)

    - mencoba arguman $ go run os.go mohammad ricky

50. Package Flag `flag.go`
    - Untuk membuat aplikasi yang berbentu terminal
    - Berisikan funngsionalitas untuk memparsing command line agument
      $ go run flag.go -host=localhost -user=root -password=root
    - https://golang.org/pkg/flag
51. Pacakge String `package-strings.go`
    - https://pkg.go.dev/strings
    - Beberapa Contoh function package string
      .strings.Trim(string,cutset)
      .strings.ToLower(string)
52. Package strconv (String Conversion) `package-strconv.go`
53. Package Math
    - Fungsi untuk matematika
    - https://pkg.go.dev/math
54. Package Container List
55. Pakage Container Ring
56. Pacakge Sort
57. Package Time
58. Package reflec
59. Package regex
60. Materi Selanjutnya
61. Pendahuluan Module
    - Module = Library, improvement dari GOPATH
62. Pengenalan Go Module
63. Membuat Module
64. Menambahkan Dependecy
65. Upgrade Module
66. Upgrade Dependecy
67. Major Upgrade
68. Materi Selanjutnya

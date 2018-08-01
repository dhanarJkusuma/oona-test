## Cara Penggunaan Dan Asumsi

# Cara Menggunakan
- Clone repository di directory GOPATH/src
- jalankan perintah 
```
go get -v
```
- jalankan aplikasi
```
go run main.go
```

# Cara menjalankan jika menggunakan Docker
- Build image
```
docker build -t <nama_image> .
```
- Jalankan
```
docker run --publish 8080:8080 --name test --rm <nama_image>
```

# Testing
Pada terminal direktori ini jalankan 
```
go test usecase/arithmatic_test.go usecase/temperature_test.go usecase/highest_sum_test.go
```

# Penghitungan dari huruf
Untuk Penghitungan sendiri, saya menggunakan helper dr saya sendiri [idr-cardinalparser](https://github.com/dhanarJkusuma/idr-cardinalparser)
Cara mengakses adalah :
```
POST localhost:8080/api/v1/arithmatic`
```
dengan BODY
```
{
	"body" : "satu tambah satu"
}
```

# Menemukan nilai tertinggi dari total sum
Untuk menemukan total sum, ini saya agak kurang paham, untuk sekarang saya berasumsi bahwa itu menggunakan `Kadane's algorithm`
[Maximum Sub Array Problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm)
Cara mengakses adalah :

```
POST localhost:8080/api/v1/high-sum
```
dengan BODY
```
{
	"body" :  [ 3,   -9,   -1,   4,   3,   -2,   3  ] 
}
```

# Penghitungan Ringkasan di Temperature
Average := rata-rata
Median := nilai tengah
Mode := nilai yg paling sering muncul

Dari beberapa case di soal ada nilai hasil median yang sepertinya kurang sesuai, dikarenakan attributenya menggunakan median, maka saya anggap itu tetap nilai tengah, jika ganjil diambil paling tengah, jika genap adalah rata-rata 2 angka dari index [len(temperatures) / 2] & [(len(temperatures) / 2)] + 1

Cara mengakses adalah :

```
POST localhost:8080/api/v1/summaries
```
dengan BODY
```
[
	{"id": "a","timestamp": 1509493641,"temperature": 3.53},
	{"id": "b","timestamp": 1509493642,"temperature": 4.13},
	{"id": "c","timestamp": 1509493643,"temperature": 3.96},
	{"id": "a","timestamp": 1509493644,"temperature": 3.63},
	{"id": "c","timestamp": 1509493645,"temperature": 3.96},
	{"id": "a","timestamp": 1509493645,"temperature": 4.63},
	{"id": "a","timestamp": 1509493646,"temperature": 3.53},
	{"id": "b","timestamp": 1509493647,"temperature": 4.15},
	{"id": "c","timestamp": 1509493655,"temperature": 3.95},
	{"id": "a","timestamp": 1509493677,"temperature": 3.66},
	{"id": "b","timestamp": 1510113646,"temperature": 4.15},
	{"id": "c","timestamp": 1510127886,"temperature": 3.36},
	{"id": "c","timestamp": 1510127892,"temperature": 3.36},
	{"id": "a","timestamp": 1510128112,"temperature": 3.67},
	{"id": "b","timestamp": 1510128115,"temperature": 3.88}
]
```
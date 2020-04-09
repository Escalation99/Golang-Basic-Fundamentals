package main

import (
	"fmt"
	"strconv"
)

//deklarasi string
var pelajaran1 = "Go"
var pelajaran2 = "PHP"

//deklarasi variabel string kosong
var sekolah1 = ""

//deklarasi integer
var umur = 10

//atau bisa juga
var sekolah2 string
var angka int

const salary = 10000

//deklarasi banyak variabel sekaligus
var namaDepan, namaBelakang = "Ray", "Tommy"

//gabungin variabel
var namaLengkap = namaDepan + " " + namaBelakang
var pelajaran3 = pelajaran1 + pelajaran2

func printNumber() int {
	return 10
}

func multiply(a int, b int) int {
	return a * b
}

func getBiography(age int, name string, status string) string {
	return name + " adalah seorang " + status + " saat ini berumur " + strconv.Itoa(age) + " tahun"
}

//fungsi dengan 2 return type
func getBiography2(age int, name string, status string) (string, string) {
	return name + " adalah seorang " + status,
		"Umurnya sekarang " + strconv.Itoa(age)
}

func getBiography3(age int, name string, status string) (string, int) {
	return name + " adalah seorang " + status,
		age + 10
}

//function tanpa return type
func getText() {
	fmt.Println("Hello World")
}

func utangTester(uang int, utang int) (pesan string) {
	if uang < utang {
		pesan = "uang kurang"
	} else if uang == utang {
		pesan = "uang pas"
	} else {
		pesan = "uang lebih"
	}
	return
}

func main() {
	fmt.Println("Fungsi menampilkan angka: ", printNumber())
	fmt.Println("Hasil perkalian: ", multiply(12, 3))
	fmt.Println(getBiography(21, "Ray", "Programmer"))

	var basicInfo, ageInfo = getBiography2(24, "roy", "Tommy")

	fmt.Println(basicInfo, ageInfo)

	var basicInfo2, age10years = getBiography3(24, "roy", "Tommy")

	fmt.Println(basicInfo2, age10years)

	getText()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0

	for j < 5 {
		fmt.Println("Nilai J adalah: ", j)
		j++
	}

	sum := 0
	for k := 1; k <= 100; k++ {
		sum = sum + k
	}
	fmt.Println(sum)

	if sum > 1000 {
		fmt.Println("Sum lebih dari 1000")
	} else {
		fmt.Println("Sum kurang dari 1000")
	}

	color := "Hijau"

	if color == "Hijau" {
		fmt.Println("Warnanya adalah hijau")
	} else if color == "Yellow" {
		fmt.Println("Warnanya adalah yellow")
	} else {
		fmt.Println("Warnanya aneh")
	}

	message := utangTester(1000, 500)
	fmt.Println(message)

	var angkaBaru = 1
	switch {
	case angkaBaru < 1:
		fmt.Println("Angka kurang dari 1")
	case angkaBaru > 1:
		fmt.Println("Angka lebih dari 1")
	case angkaBaru == 1:
		fmt.Println("Angka sama dengan 1")
	}

	var warna = "Merah"
	switch warna {
	case "Merah":
		fmt.Println("Warna Merah")
	case "Kuning":
		fmt.Println("Warna Kuning")
	case "Hijau":
		fmt.Println("Warna Hijau")
	default:
		fmt.Println("Warna salah")
	}

	//menyimpan alamat dari suatu variabel
	var point = 100
	//pointAddress adalah sebuah variabel integer yang merujuk kepada address dari variabel point
	var pointAddress *int = &point
	fmt.Println(pointAddress)  // untuk print address dari yang dia point
	fmt.Println(*pointAddress) // untuk print value dari yang dia point

	fmt.Println("Hello"[0])       //print index tertentu pada suatu string
	fmt.Println(len("Hello"))     //print jumlah karakter pada suatu string
	fmt.Println("Hello"[0:4])     //print substring
	fmt.Println("HelloWorld"[5:]) //print substring dari awal sampai akhir

	//buat array item lalu tampilkan dengan foreach
	item2 := [3]string{
		"a", "b", "c",
	}

	for i := 0; i < len(item2); i++ {
		fmt.Println(item2[i])
	}

	//atau bisa juga satu-satu kayak gini
	var item [3]string
	item[0] = "Mangga"
	item[1] = "Apel"
	item[2] = "Pisang"

	for _, value := range item {
		fmt.Println(value)
	}

	//bisa juga slice arraynya
	item3 := [5]string{
		"mama", "papa", "adik", "kakak", "nenek",
	}
	slice1 := item3[1:3]

	fmt.Println(slice1)

	//membuat map
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Ray"
	mahasiswa["1002"] = "Tommy"
	mahasiswa["1003"] = "BINUS"

	for nim, name := range mahasiswa {
		fmt.Println("Nim" + nim + " Bernama " + name)
	}

	//inisialisasi value pada map
	elemen := map[string]string{
		"fire":  "api",
		"water": "air",
		"wind":  "udara",
		"earth": "tanah",
	}

	for elemen1, elemen2 := range elemen {
		println("Elemen " + elemen2 + " Disebut " + elemen1)
	}

	//membuat map dalam map
	pelajar := map[string]map[string]string{
		"pelajar1": map[string]string{
			"nama": "budi",
			"nim":  "1234",
		},
		"pelajar2": map[string]string{
			"nama": "aldi",
			"nim":  "12345",
		},
		"pelajar3": map[string]string{
			"nama": "charlie",
			"nim":  "123456",
		},
	}

	fmt.Println(pelajar["pelajar1"]["nama"])

	//menghapus data dari map
	delete(pelajar, "pelajar1")
	fmt.Println(pelajar)

	//menerima inputan
	var pilihan string
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanf("%s", &pilihan)
	fmt.Println(mahasiswa[pilihan])

	//menerima inputan integer
	var pilihan2 int
	fmt.Println("Masukkan nim: ")
	fmt.Scanf("%d", &pilihan2)
	fmt.Println("Pilihan anda: " + strconv.Itoa(pilihan2))

	//menerima dan menampilkan input
	var namaMahasiswa string
	var nimMahasiswa int

	fmt.Print("Masukkan nama mahasiswa: ")
	fmt.Scanf("%s", &namaMahasiswa)
	fmt.Print("Masukkan nim mahasiswa: ")
	fmt.Scanf("%d", &nimMahasiswa)
	fmt.Printf("%s %d \n", namaMahasiswa, nimMahasiswa)

	//deklarasi variabel didalam function
	//alamat := "binus square"
	//umurSaya := 20

	//fmt.Println("Alamat sekolah: " + alamat)
	//fmt.Println("Hello World, saya sedang belajar " + pelajaran3)
	//fmt.Println(" Umur saya", umurSaya, "Nama saya "+namaLengkap)

	// a := 10
	// b := 20

	// gaji := 10000
	// gaji2 := strconv.Itoa(gaji) //convert int ke string

	// bonus := "10000"
	// bonus2, _ := strconv.Atoi(bonus) //convert string ke int

	// fmt.Println("Total: ", a+b)
	// fmt.Println("Gaji: " + gaji2)
	// fmt.Println("Bonus: ", bonus2)
}

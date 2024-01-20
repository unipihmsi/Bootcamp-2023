package main

import (
	//"bootcamp-go/myInterface"
	//interfaces "bootcamp-tugas/myinterfaces"
	//"bootcamp-go/Latihan1"
	//"bootcamp-go/testcase1"
	"errors"
	"fmt"
)

// Latihan1.Coba()
// testcase1.Coba()

// ---Struct---
// type customer struct {
// 	id      uint
// 	name    string
// 	address string
// 	contact
// }

// //---Embedded Struct---
// type contact struct {
// 	phone, email string
// }

// //---Method---
// func (cust customer) getName() string {
// 	return cust.name
// }
// func (cust customer) getContact() map[string]string {
// 	return map[string]string{
// 		"Phone": cust.contact.phone,
// 		"Email": cust.contact.email,
// 	}
// }

// ---Interface---
// type customer struct {
// 	Id   uint8
// 	Name string
// }

// type customerSlice []customer

// func (v customerSlice) Len() int {
// 	return len(v)
// }

// func (v customerSlice) Swap(i, j int) {
// 	v[i], v[j] = v[j], v[i]
// }

// func (v customerSlice) Less(i, j int) bool {
// 	return v[i].Id < v[j].Id
// }

func main() {
	//---PERTEMUAN 1---
	// main - disini saya buat func untuk main
	// TO DO ! print hello world
	// fmt.Println("Hello World")
	// fmt.Println("Hello World KEDUA")

	//================================================================
	//---PERTEMUAN 2---

	// ---manifest types---
	// var _ string = "Adam"
	// var umur int = 19
	// var alamat string = "Cikupa"

	// fmt.Println("Nama Saya", firstname)
	// fmt.Println("Umur Saya", age)

	// ---inference types---
	// lastname := "Kurniawan"
	// fmt.Println("Nama Akhir Saya", lastname)

	// ---buat variabel nama, umur dan alamat---
	// ---nama saya Adam Kurniawan umur 19 alamat Cikupa---
	// ---PrintF---

	// fmt.Printf("Nama Saya %s Umur Saya %d Alamat %s \n", firstname, age, alamat)

	// ---multiple variabel---
	// var nama, umur, alamat = "Adam", 19, "Cikupa"
	// fmt.Printf("Nama Saya %s Umur Saya %d Alamat %s \n", nama, umur, alamat)

	// ---underscore variabel---
	// organisasi, _ := "HMSI", "UNIPI"
	// fmt.Println(organisasi)

	// fmt.Printf("Umur Saya %d Alamat %s \n", umur, alamat)

	// ---desimal variabel---
	// var myDesimal float64 = 36.89
	// fmt.Printf("bilangan desimal: %f\n", math.Round(myDesimal))

	// ---tipe data string---
	// var address string
	// var isBootcamp bool
	// fmt.Printf("Alamat: %s \n", address)
	// fmt.Printf("Alamat: %t \n", isBootcamp)

	//================================================================================================
	//---PERTEMUAN 3---

	//---operator aritmatika---
	//var nilai = (((2+6)%3)*4 - 2) / 3
	//fmt.Println(nilai)

	//---operator logika---
	//var left = false
	//var right = true

	//var leftAndRight = left && right
	//fmt.Printf("Left && Right \t(%t) \n", leftAndRight)

	//---operator assignment---
	//var tinggibadan uint8 = 179
	//tinggibadan = tinggibadan + 11
	//tinggibadan -= 11
	//fmt.Println(tinggibadan)

	//---Latihan---
	// ---No 1---
	// var MeminjamUang float64 = 10000000
	// var Bunga float64 = 0.1
	// var Cicilan float64 = 12
	// TotalBunga := MeminjamUang * Bunga
	// TotalPinjaman := MeminjamUang + TotalBunga
	// CicilanPerbulan := TotalPinjaman / Cicilan
	// fmt.Println("Nomor 1")
	// fmt.Println("A)", int(TotalBunga))
	// fmt.Println("B)", int(TotalPinjaman))
	// fmt.Println("C)", int(CicilanPerbulan))

	// //---No 2---
	// var NilaiAbsensiDapat float32 = 16
	// var NilaiAbsensiDefault float32 = 18
	// var bobotNilaiAbsensi float32 = 0.1
	// var NilaiTugas float32 = 70
	// var bobotNilaiTugas float32 = 0.2
	// var NilaiUts float32 = 80
	// var bobotNilaiUts float32 = 0.3
	// var NilaiUas float32 = 70
	// var bobotNilaiUas float32 = 0.4
	// NilaiAbsensiResult := NilaiAbsensiDapat / NilaiAbsensiDefault
	// NilaiAkhir := (NilaiAbsensiResult * bobotNilaiAbsensi) + (NilaiTugas * bobotNilaiTugas) + (NilaiUts * bobotNilaiUts) + (NilaiUas * bobotNilaiUas)
	// fmt.Println("Nomor 2")
	// fmt.Println("Nilai Akhir =", (NilaiAkhir))

	// //---No 3 ---
	// var Jarak float32 = 63
	// var HabisBensin float32 = 17
	// var BiayaBensin float32 = 8350
	// fmt.Println("Nomor 3")
	// TotalBensin := Jarak / HabisBensin
	// fmt.Println("Total Bensin =", (TotalBensin))
	// TotalBiaya := BiayaBensin * TotalBensin
	// fmt.Println("Total Biaya =", (TotalBiaya))

	//================================================================================================
	//---PERTEMUAN 4---

	//---if else statement---
	// var point = 6
	// if point == 10 {
	// 	fmt.Println("Lulus dengan nilai sempurna")
	// } else if point >= 6 {
	// 	fmt.Println("Lulus")
	// } else if point == 5 {
	// 	fmt.Println("Hampir Lulus")
	// } else {
	// 	fmt.Printf("Tidak lulus, nilai anda %d\n", point)
	// }

	//di golang tidak mengenal tanpa kurung kurawal
	//di golang tidak mengenal ternary

	//---Variabel Temporary---
	// var point float32 = 7850.0
	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }

	//---Switch Case---
	// var point = 7
	// switch point {
	// case 8:
	// 	fmt.Println("Perfect")
	// case 7:
	// 	fmt.Println("Awesome")
	// 	fallthrough
	// default:
	// 	fmt.Println("Not Bad")
	// }
	//fallthrough untuk memaksa seleksi selanjutnya

	//---Seleksi Kondisi Bersarang---
	// var point = 4

	// if point >= 7 {
	// 	switch point {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice!")
	// 	}
	// } else {
	// 	if point >= 5 {
	// 		fmt.Println("not bad")
	// 	} else if point == 4 {
	// 		fmt.Println("you can do it")
	// 	} else {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	//---Looping---
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Urutan nomor ke-", i)
	// }

	// fmt.Println("===================")

	// var i int = 0
	// for i < 5 {
	// 	fmt.Println("Urutan nomor ke-", i)
	// 	i++
	// }

	//---Break & Continue---
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Angka Ganjil Adalah ", i)
	// }

	//---LATIHAN---
	// var nomor int = 10
	// for i := 1; i <= nomor; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d", j)
	// 	}
	// 	fmt.Println()
	// }

	//---FizzBuzz---
	// for i := 1; i <= 30; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("insan Pemmbangunan")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Pembangunan")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Insan")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//---LATIHAN---
	//---Menghitung pecahan uang---
	// var totalbelanja int = 13500
	// var uangdiberikan int = 100000
	// fmt.Println("Total Belanja:", totalbelanja)
	// fmt.Println("Uang Diberikan:", uangdiberikan)

	// kembalian := uangdiberikan - totalbelanja
	// fmt.Println("Uang Kembalian = ", kembalian)

	// u50 := kembalian % 50000
	// pecahan50k := (kembalian - u50) / 50000
	// kembalian = kembalian - (pecahan50k * 50000)
	// fmt.Println("Jumlah Pecahan 50000 = ", pecahan50k)

	// u20 := kembalian % 20000
	// pecahan20k := (kembalian - u20) / 20000
	// kembalian = kembalian - (pecahan20k * 20000)
	// fmt.Println("Jumlah Pecahan 20000 = ", pecahan20k)

	// u10 := kembalian % 10000
	// pecahan10k := (kembalian - u10) / 10000
	// kembalian = kembalian - (pecahan10k * 10000)
	// fmt.Println("Jumlah Pecahan 10000 = ", pecahan10k)

	// u5 := kembalian % 5000
	// pecahan5k := (kembalian - u5) / 5000
	// kembalian = kembalian - (pecahan5k * 5000)
	// fmt.Println("Jumlah Pecahan 5000 = ", pecahan5k)

	// u2 := kembalian % 2000
	// pecahan2k := (kembalian - u2) / 2000
	// kembalian = kembalian - (pecahan2k * 2000)
	// fmt.Println("Jumlah Pecahan 2000 = ", pecahan2k)

	// u1 := kembalian % 1000
	// pecahan1k := (kembalian - u1) / 1000
	// kembalian = kembalian - (pecahan1k * 1000)
	// fmt.Println("Jumlah Pecahan 1000 = ", pecahan1k)

	// u05 := kembalian % 500
	// pecahan05k := (kembalian - u05) / 500
	// kembalian = kembalian - (pecahan05k * 500)
	// fmt.Println("Jumlah Pecahan 500 = ", pecahan05k)

	//===============================================================================================
	//---PERTEMUAN 5---

	//---Array---
	// var names [3]string
	// names[0] = "Adam"
	// names[1] = "Kurniawan"
	// names[2] = "Aja"
	// fmt.Println(names[0], names[1], names[2])

	//---Inisialisasi nilai awal Array
	// var names = [3]string{"Adam", "Kurniawan", "Ajah"}
	// fmt.Println("Jumlah element \t\t", len(names))
	// fmt.Println("Isi semua element \t", names)

	//---Looping array menggunakan For---
	// var names = [3]string{"Adam", "Kurniawan", "Ajah"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Printf("Elemen %d : %s\n", i, names[i])
	// }

	//---Looping array menggunakan For Range---
	// var names = [4]string{"Adam", "Kurniawan", "Ajah", "Keren"}
	// for i, name := range names {
	// 	fmt.Printf("ELemen %d : %s\n", i, name)
	// }

	//---Array keyword Make---
	// var names = make([]string, 3)
	// names[0] = "Adam"
	// names[1] = "Kurniawan"
	// names[2] = "Aja"
	// fmt.Println(names)

	//---Slice---
	// var names = []string{"Adam", "Kurniawan", "Adeww"}
	// fmt.Println(names)
	// fmt.Println(names[1])

	//---Fungsi Append---
	// var names = []string{"Adam", "Kurniawan"}
	// var namesSecond = append(names, "S.Kom")
	// fmt.Println(names)
	// fmt.Println(names[1])
	// fmt.Println(namesSecond)

	//---MAP---
	// var months map[string]string
	// months = map[string]string{}
	// months["Jan"] = "Januari"
	// months["Feb"] = "Februari"
	// fmt.Println(months["Jan"])
	// fmt.Println(months["Feb"])

	//---Inisialisasi nilai Map---
	//cara horizontal
	// var months1 = map[string]int{"Januari": 1, "Februari": 2}
	// fmt.Println("months 1 ", months1)
	// //cara vertical
	// var months2 = map[string]int{
	// 	"Januari":  1,
	// 	"Februari": 2,
	// }
	// fmt.Println("months 2 ", months2)

	//---Map For Range---
	// var months1 = map[string]int{"Januari": 1, "Februari": 2, "Maret": 3}
	// for key, value := range months1 {
	// 	fmt.Println(key, " \t:", value)
	// }

	//---Delete Map---
	// var months = map[string]int{"Januari": 1, "Februari": 2, "Maret": 3}
	// delete(months, "Januari")
	// fmt.Println(months)

	//---LATIHAN---
	// var number = []int{2, 5, 6, 8, 9, 2}
	// summary := 0
	// for _, number := range number {
	// 	summary += number
	// }
	// average := float64(summary) / float64(len(number))
	// fmt.Println("Total Summary adalah", summary)
	// fmt.Printf("Average adalah %.1f", average)

	//---NOMOR 1---
	// 	fmt.Println("---NOMOR 1---")
	// 	var largestNumber, temp int
	// 	number1 := []int{11, 6, 31, 201, 99, 861, 1, 7, 14, 79}
	// 	fmt.Println("Sederet Bilangan : ", number1)
	// 	for _, element := range number1 {
	// 		if element > temp {
	// 			temp = element
	// 			largestNumber = temp
	// 		}
	// 	}
	// 	fmt.Println("Bilangan Terbesar adalah : ", largestNumber)

	// 	//---NOMOR 2---
	// 	fmt.Println("---NOMOR 2---")
	// 	number2 := []int{99, 2, 64, 8, 111, 33, 65, 11, 102, 50}
	// 	fmt.Println("Sederet Bilangan : ", number2)
	// 	fmt.Println("Bilangan Diurutkan", BubbleSort(number2))

	// 	//---NOMOR 3---
	// 	fmt.Println("---NOMOR 3---")
	// 	number3 := []int{55, 31, 20, 1, 10, 89, 90, 66, 11, 21}
	// 	fmt.Println("Sederet Bilangan : ", number3)
	// 	sort.Ints(number3)
	// 	fmt.Println("Bilangan Diurutkan : ", number3)
	// }

	// func BubbleSort(array []int) []int {
	// 	for i := 0; i < len(array)-1; i++ {
	// 		for j := 0; j < len(array)-i-1; j++ {
	// 			if array[j] > array[j+1] {
	// 				array[j], array[j+1] = array[j+1], array[j]
	// 			}
	// 		}
	// 	}
	// 	return array

	//===============================================================================================
	//---PERTEMUAN 6---

	//---Fungsi di Golang---
	//---Penulisan Fungsi---
	// var name = ""
	// PrintPesertaBootcamp(name)

	//---Return Value---
	// Hasil := getGreater(9, 3)
	// fmt.Println("Nilai Terbesar adalah  ", Hasil)

	//---Multiple Return---
	// var diameter float64 = 15
	// luas, keliling := hitung(diameter)
	// fmt.Printf("Luas lingkaran : %.2f\n", luas)
	// fmt.Printf("Keliling lingkaran : %.2f\n", keliling)

	//---Fungsi Variadic Slice---
	// numbers := []int{1, 2, 3, 4, 5, 6, 7}
	// total := getTotal(numbers...)
	// fmt.Printf("Total : %d \n", total)

	//---Fungsi Variadic---
	// total2 := getTotal(1, 2, 3, 4, 5)
	// fmt.Printf("Total : %d \n", total2)

	//---Fungsi Variadic Parameter---
	// mySports("Adam", "Futsal", "Bultang", "E-sport")

	//---Fungsi Closure---
	// var getMinMax = func(numbers []int) (int, int) {
	// 	var min, max int

	// 	for i, n := range numbers {
	// 		if i == 0 {
	// 			min, max = n, n
	// 		} else if n > max {
	// 			max = n
	// 		} else if n < min {
	// 			min = n
	// 		}
	// 	}
	// 	return min, max
	// }

	// var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// min, max := getMinMax(numbers)
	// fmt.Printf("Nilai min adalah %d dan Nilai max adalah %d \n", min, max)

	//---Fungsi Closure IIFE (Immediately-Invoked Function Expression)---
	// var getModulus = func(number int) int {
	// 	return number % 3
	// }(11)
	// fmt.Printf("Nilai sisa bagi adalah %d \n", getModulus)

	//---Fungsi dengan parameter string, return value bool
	//---Nama fungsi isPalindrome

	//---Fungsi palindrome---
	// name := "ADA"
	// hasil := isPalindrome(name)
	// fmt.Println(name)
	// if hasil == true {
	// 	fmt.Println("Nama ini Palindrome")
	// } else {
	// 	fmt.Println("Nama ini Bukan Palindrome")
	// }

	//===============================================================================================
	//---PERTEMUAN 7---

	//---Pointer---
	// var angkaPertama int = 4
	// var angkaKedua *int = &angkaPertama
	// fmt.Println("Angka pertama (value) : ", angkaPertama)
	// fmt.Println("Angka pertama (address) : ", &angkaPertama)
	// fmt.Println("Angka kedua (value) : ", *angkaKedua)
	// fmt.Println("Angka pertama (address) : ", &angkaPertama)

	//---Parameter Pointer---
	// var angka int = 4
	// ubah(&angka, 6)
	// fmt.Println(angka)

	//---Struct---
	// var mycust customer
	// mycust.id = 1
	// mycust.name = "Adam"
	// mycust.address = "Surga"
	// fmt.Println(mycust)

	//---Inisialisasi Struct---
	// var cust = customer{id: 2, name: "adam"}
	// fmt.Println(cust)

	//---Embedded Struct---
	// var mycust = customer{
	// 	id:      1,
	// 	name:    "adam",
	// 	address: "surga",
	// 	contact: contact{email: "adam@gmail.com", phone: "085711228241"},
	// }
	// println("Nama Customer :", mycust.name)
	// println("Address Customer :", mycust.address)
	// println("Contact Email Customer :", mycust.email)
	// println("Contact Phone Customer :", mycust.phone)

	//---Kombinasi slice dan struct
	// var v = []customer{
	// 	{id: 1, name: "Adam", address: "Cikupa", contact: contact{phone: "085711228241", email: "adam@gmail.com"}},
	// 	{id: 1, name: "Adew", address: "Magelang", contact: contact{phone: "081293564998", email: "adew@gmail.com"}},
	// }
	// for _, v := range v {
	// 	fmt.Printf("Nama Customer %s, Alamat %s, Kontak : Phone(%s), Email(%s)\n", v.name, v.address, v.contact.phone, v.contact.email)
	// }

	//---Method---
	// var customer = customer{id: 1, name: "Adam", address: "Cikupa", contact: contact{phone: "085711228241", email: "adam@gmail.com"}}
	// customername := customer.getName()
	// customercontacts := customer.getContact()

	// fmt.Println("Customer name is ", customername)
	// for key, v := range customercontacts {
	// 	fmt.Printf("Customer %s is %s \n", key, v)
	// }

	//===============================================================================================
	//---PERTEMUAN 8---

	//---Interface---
	//---Cara Deklasrasi Interface---
	//var infoCustomer myInterface.Information

	//---Buat Struct Customer---
	// var dataCustomer = myInterface.Customer{Id: 1, Name: "Adam", Address: "Tangerang", Contact: myInterface.Contact{
	// 	Email: "Adam@gmail.com", Phone: "085711228241"}}

	// infoCustomer = &dataCustomer

	// name := infoCustomer.GetName()
	// fmt.Println("Nama Customer ", name)

	// contacts := infoCustomer.GetContact()
	// for k, contact := range contacts {
	// 	fmt.Printf("Contact %s is %s \n", k, contact)
	// }

	//---Interface Kosong---
	// var kosong myInterface.InterfaceKosong
	// kosong = "Ini adalah string"
	// fmt.Println(kosong)

	// kosong = 1234
	// fmt.Println(kosong)

	// kosong = true
	// fmt.Println(kosong)

	//---Casting Interface---
	// var numberInterface any
	// numberInterface = 10
	// hasilperkalian := numberInterface.(int) * 2
	// fmt.Printf("hasil perkalian adalah %d \n", hasilperkalian)

	//---Casting Interface kosong ke objek Pointer---
	// var objectInterface interface{}
	// objectInterface = &myInterface.Customer{Id: 1, Name: "Adam", Address: "Tangerang", Contact: myInterface.Contact{
	// 	Email: "Adam@gmail.com", Phone: "085711228241"}}

	// name := objectInterface.(*myInterface.Customer).Name
	// fmt.Println(name)

	//---Casting Interface kosong ke Bolean---
	// var boolInterface interface{}
	// boolInterface = true
	// nilaiAsli := boolInterface.(bool)
	// fmt.Println(nilaiAsli)

	//---Interface---
	// customers := []customer{
	// 	{Id: 1, Name: "Adam"},
	// 	{Id: 2, Name: "Adai"},
	// 	{Id: 3, Name: "Tegar"},
	// 	{Id: 4, Name: "Wisnu"},
	// }
	// fmt.Println(customers)
	// sort.Sort(customerSlice(customers))

	//---TUGAS INTERFACE---
	// fmt.Println("A.) Berapa umur cinta, jika ayu berumur 10 tahun?")
	// ayu := interfaces.Anggota{Nama: "Ayu", Umur: 10}
	// cinta := interfaces.Anggota{
	// 	Nama: "Cinta",
	// 	Umur: ayu.GetUmur() + 5,
	// }
	// fmt.Printf("Jawab : Umur %s adalah %d tahun\n\n", cinta.Nama, cinta.Umur)

	// fmt.Println("B.) Berapa umur budi dan ibu, jika ayu berumur 10 tahun?")
	// budi := interfaces.Anggota{
	// 	Nama: "Budi",
	// 	Umur: cinta.GetUmur() + 2,
	// }
	// fmt.Printf("Jawab :  Umur %s adalah %d tahun\n", budi.Nama, budi.Umur)

	// ibu := interfaces.Anggota{
	// 	Nama: "Ibu",
	// 	Umur: 3*ayu.GetUmur() + 2,
	// }
	// fmt.Printf("\t Umur %s adalah %d tahun\n\n", ibu.Nama, ibu.Umur)

	// fmt.Println("C.) Berapa umur ayah,jika ayu berumur 10 tahun?")
	// ayah := interfaces.Anggota{
	// 	Nama: "Ayah",
	// 	Umur: ayu.GetUmur() + budi.GetUmur() + cinta.GetUmur(),
	// }
	// fmt.Printf("Jawab : %s adalah %d tahun\n\n", ayah.Nama, ayah.Umur)

	// fmt.Println("D.) Berdasarkan perhitungan diatas, urutkan berdasarkan anak tertua hingga termuda?")
	// totalKeluarga := interfaces.Family{
	// 	ayu,
	// 	budi,
	// 	cinta,
	// }
	// fmt.Println("Jawab : ")
	// sort.Sort(totalKeluarga)
	// for _, v := range totalKeluarga {
	// 	fmt.Printf("- Umur %s adalah %d tahun\n", v.GetNama(), v.GetUmur())
	// }

	//===============================================================================================
	//---PERTEMUAN 9---

	//---Defer---
	// pesan("ketoprak")

	//--Kombinasi defer dan IIFE---
	// pesanan := "ketoprak"
	// if pesanan == "ketoprak" {
	// 	fmt.Println("pilihan tepat, ketoprak di tempat kami paling enak")
	// 	defer fmt.Println("terimakasih sudah pesan di warung kami")
	// }
	// fmt.Println("Pesanan anda adalah ", pesanan)

	//---Exit---
	// untuk menghentikan paksa sebuah program
	// defer fmt.Println("halo")
	// os.Exit(1)
	// fmt.Println("Selamat Datang")

	//---Error---
	// var input string = "8"
	// var errConvertion error
	// var result int

	// result, errConvertion = strconv.Atoi(input)

	// if errConvertion != nil {
	// 	fmt.Printf("Error : %s\n", errConvertion)
	// } else {
	// 	fmt.Printf("Result : %d\n", result)
	// }

	//---Custom Error---
	avg, err := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	} else {
		var msg = fmt.Sprintf("Average : %.2f\n", avg)
		fmt.Println(msg)
	}
}

// ---Custom Error---
func average(numbers ...int) (float64, error) {
	var avg float64 = 0
	var total int = 0
	if len(numbers) > 0 {
		for _, v := range numbers {
			total += v
		}
		avg = float64(total) / float64(len(numbers))
	} else {
		return avg, errors.New("numbers can't be empty")
	}
	return avg, nil
}

// ---Defer----
// func pesan(pesanan string) {
// 	defer fmt.Println("terimakasih sudah pesan di warung kami")

// 	if pesanan == "ketoprak" {
// 		fmt.Println("pilihan tepat, ketoprak di tempat kami paling enak")
// 		return
// 	}
// 	fmt.Println("Pesanan anda adalah ", pesanan)
// }

// ---Parameter Pointer---
// func ubah(nomor *int, nilai int) {
// 	*nomor = nilai
//}

//---Fungsi palindrome---
// func isPalindrome(name string) bool {
// 	for i := 0; i < len(name); i++ {
// 		if name[i] != name[len(name)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// ---Fungsi Variadic Parameter---
// func mySports(name string, sports ...string) {
// 	fmt.Println("Nama : ", name)
// 	fmt.Println("Olahraga Kesukaan : ", sports)
// }

// ---Fungsi Variadic---
// func getTotal(numbers ...int) int {
// 	var total int
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// ---Multiple Return---
// func hitung(diameter float64) (float64, float64) {
// 	luas := math.Pi * math.Pow(diameter/2, 2)
// 	keliling := math.Pi * diameter
// 	return luas, keliling
// }

// ---Return Value---
// func getGreater(firstnumber, secondnumber int) int {
// 	var result int
// 	if firstnumber > secondnumber {
// 		result = firstnumber
// 	}
// 	if firstnumber < secondnumber {
// 		result = secondnumber
// 	}
// 	return result
// }

// ---Penulisan Fungsi---
// func PrintPesertaBootcamp(name string) {
// 	if name == "" {
// 		fmt.Println("Nama Peseta Bootcamp : Kosong...")
// 		return
// 	}
// 	fmt.Println("Nama Peserta Bootcamp adalah ", name)
// }

package main

import "fmt"

func main() {
	// Variabel
	var nama string = "John"
	umur := 25
	var tinggi float64 = 175.5

	// Tipe Data
	fmt.Printf("Nama saya %s, umur saya %d, tinggi saya %.2f\n", nama, umur, tinggi)

	// Perulangan Array
	var arr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

	// Perulangan Array Object dengan Index
	var siswa = [3]struct{ nama string; kelas string }{
		{"Sarah", "12A"},
		{"John", "11B"},
		{"Mike", "10C"},
	}
	for i, s := range siswa {
		fmt.Printf("\nSiswa ke-%d: Nama = %s, Kelas = %s", i+1, s.nama, s.kelas)
	}

	// Map
	var buah = map[string]int{"apel": 10, "jeruk": 20, "anggur": 30}
	for k, v := range buah {
		fmt.Printf("\n%s = %d", k, v)
	}

	// Percabangan
	var x = 10
	if x < 5 {
		fmt.Println("\nx kurang dari 5")
	} else if x == 5 {
		fmt.Println("\nx sama dengan 5")
	} else {
		fmt.Println("\nx lebih dari 5")
	}

	// Kondisional
	var y = 20
	var z = 30
	if y > 10 && z > 20 {
		fmt.Println("\ny lebih besar dari 10 dan z lebih besar dari 20")
	}

	// Aritmatika
	var a = 5
	var b = 3
	var tambah = a + b
	var kurang = a - b
	var kali = a * b
	var bagi = a / b
	var modulo = a % b
	fmt.Printf("\nHasil penjumlahan: %d\n", tambah)
	fmt.Printf("Hasil pengurangan: %d\n", kurang)
	fmt.Printf("Hasil perkalian: %d\n", kali)
	fmt.Printf("Hasil pembagian: %d\n", bagi)
	fmt.Printf("Sisa hasil bagi: %d\n", modulo)
}
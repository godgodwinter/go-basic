package main

// https://go.dev/doc/effective_go#arrays
import "fmt"

func array() {
	// array adalah kumpulan data dengan tipe data yang sama
	// tentukan jumlah data yang bisa ditampung
	// daya tampung tidak bisa bertambah setelah array dibuat
// buat manual
	var names [3] string
	names[0]="Babeng"
	names[1]="Bata"
    names[2]="Batu"

	fmt.Println(names[1],names)
	// cara lain
	var values=[3]int{
		90,
		20,
		49,
	}

	fmt.Println(values)

	// fungsi untuk array
	// len(array) = //mandapat panjang array
	// array[index] = //mendapat data pada posisi index
	// array[index]=value //mengubah data pada posisi index

	fmt.Println(len(values))
    fmt.Println(values[0])

	// apabila tidak tahu jml array berapa bisa gunakan ...
	var months=[...]string{
		"januari",
		"februari",
        "mare",
	}
	fmt.Println(months)
}

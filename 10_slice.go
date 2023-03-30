package main

import "fmt"

func slice() {
// tipe data slice 
// tipe data slice adalah potongan dari data array
// slice mirip dengan array  yang membedakan adlah ukuran slice bisa berubah
// slice dan array selalu terkoneksi , dimana slice data yang menakses sebagian atau seluruh data di array

// slice memiliki 3 data yaitu pointer, length dan capacity
// pointer menunjukan data pertama di array para sliceconst
// length adlah panjang dari slice
// capacity adalah kapasistas dari slice, dimana length tidak boleh lebih dari capacity
var months=[...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
    "Mei",
    "Juni",
	"Juli",
    "Agustus",
    "September",
	"Oktober",
    "November",
    "Desember",
}

var slice1=months[0:3]
fmt.Println(slice1,len(slice1),cap(slice1))

// jika data pada slice diubah maka array sumber akan berubah

// append(slice,data) //digunakan untuk membuat data array baru 
//copy(slice,data) //menyalin slice dari source ke destinasi baru
//make([]Typedata,length, capacity) //membuat slice baru


}

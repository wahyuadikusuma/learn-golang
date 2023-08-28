package main

import "fmt"

func main() {
	var months = [12]string{
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

	var slice = months[4:7] // mengakses index ke 4 sampai ke-6 (sebelum 7)
	fmt.Println(slice)
	fmt.Println("length : ", len(slice))
	fmt.Println("capacity : ", cap(slice)) // hasilnya 8, karena dihitung mulai dari index 4 sampai 12
	
	fmt.Println("\nSetelah diubah")
	slice[1] = "Juni ini bos"
	fmt.Println(slice)
	fmt.Println(months) // arraynya ikut berubah walaupun yang diubah adalah slicenya

	var slice2 = months[9:]
	fmt.Println("\n",slice2)


	//ketika ingin menambah data baru, maka yang terjadi adalah membuat array baru
	var slice3 = append(slice2, "!Bulan Sabit!")
	fmt.Println(slice3)
	// array months tidak terpengaruh 
	// bisa kena dampak jika yang diappend adalah slice2 = months [2:4] (ditengah" array)
	fmt.Println(months)

	fmt.Println("\n")

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Wahyu"
	newSlice[1] = "Adi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("hasilnya sama :",copySlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
package main

import "fmt"

func main() {

	isim := "dogukan"

	var yeni_isim *string = &isim //isimin bellekteki adresi

	*yeni_isim = "bicer" //bellekteki adrese isim aktarıldı

	fmt.Println(isim, yeni_isim) // çıktı:bicer 0xc00004a250
}

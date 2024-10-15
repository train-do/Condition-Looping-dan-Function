package main

import "fmt"
func totalHarga(input []string) int {
	arrSepatu := []string{"adidas", "puma", "kappa"}
	arrHarga := []int{200000, 150000, 600000}
	var total int
	for _, value := range input{
		for j := 0; j < len(arrSepatu); j++ {
			if value == arrSepatu[j] {
				total += arrHarga[j]
			}
		}
	}
	return total
}
func hitungDiskon(input []string) int {
	var isAdidas, isPuma, isKappa bool
	var diskon int
	for _, value := range input{
		switch value {
		case "adidas":
			isAdidas = true
		case "puma":
			isPuma = true
		case "kappa":
			isKappa = true
		}
	}
	switch {
	case isAdidas && isPuma:
		diskon = 50000
	case isKappa && isPuma:
		diskon = 150000
	case isAdidas && isKappa:
		diskon = 75000
	}
	return diskon
}
func hitungTotalBayar(sepatu []string)  {
	totalHarga := totalHarga(sepatu)
	diskon := hitungDiskon(sepatu)
	fmt.Println("Harga Belanja : Rp. ", totalHarga-diskon, "(",totalHarga,", diskon", diskon,")")
}
func main() {
	a := []string {"adidas", "puma"}
	b := []string {"kappa", "puma"}
	c := []string {"kappa", "adidas"}
	hitungTotalBayar(a)
	hitungTotalBayar(b)
	hitungTotalBayar(c)
}
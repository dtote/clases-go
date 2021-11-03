package main

import "fmt"

func main() {
	even, odd := evenAndOddNumbersBetween1AndUpperBound(100)
	primes := primeNumbersBetween1AndUpperBound(100)
	sum := addAllNumbersFrom1ToUpperBound(100)
	mult, div := multAndDiv(6, 3)
	fmt.Println("\n--> Números Pares =", even)
	fmt.Println("\n--> Números Impares =", odd)
	fmt.Println("\n--> Números Primos ", primes)
	fmt.Println("\n--> Suma números del 1 al 100 =", sum)
	fmt.Println("\n--> Multiplicación: 6 * 3 =", mult, " || División: 6 / 3 =", div)
}

func evenAndOddNumbersBetween1AndUpperBound(upperBound int) (sliceEven, sliceOdd []int) {
	for i := 1; i <= upperBound; i++ {
		if i%2 == 0 {
			sliceEven = append(sliceEven, i)
		} else {
			sliceOdd = append(sliceOdd, i)
		}
	}
	return
}

func addAllNumbersFrom1ToUpperBound(upperBound int) (result int) {
	for i := 1; i <= upperBound; i++ {
		result += i
	}
	return
}

func primeNumbersBetween1AndUpperBound(upperBound int) (primeNumbers []int) {
	for i := 2; i <= upperBound; i++ {
		isPrime := true
		for j := 2; j < i && isPrime; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return
}

func multAndDiv(x float64, y float64) (mult float64, div float64) {
	mult = x * y
	div = x / y
	return
}

func checkMark() (result string) {
	var mark int
	fmt.Scanf("Introduzca una nota: %v", mark)
	switch {
	case mark >= 0 && mark < 5:
		result = "Insuficiente"
	case mark >= 5 && mark < 7:
		result = "Suficiente"
	case mark >= 7 && mark < 9:
		result = "Notable"
	case mark >= 9 && mark < 11:
		result = "Sobresaliente"
	default:
		result = "Nota no válida"
	}
	return
}

package main

import (
	"fmt"
)

func main() {
	poly1 := []int{81, 95, 74, 91, 72, 90, 68, 61, 11, 26, 70, 4, 15, 98, 9, 39, 0, 48, 86, 58, 34, 93, 37, 76, 87, 96, 14, 17, 71, 22, 23, 27, 64, 66, 6, 5, 56, 69, 28, 88, 89, 73, 10, 13, 21, 82, 55, 31, 45, 2, 97, 92, 7, 60, 24, 85, 8, 29, 75, 57, 40, 62, 54, 19, 100, 18, 43, 63, 33, 65, 78, 80, 38, 20, 84, 25, 59, 94, 50, 30, 46, 32, 41, 67, 35, 47, 52, 53, 44, 49, 12, 83, 79, 42, 16, 3, 77, 51, 1, 36}
	poly2 := []int{12, 43, 65, 41, 62, 80, 19, 89, 98, 45, 31, 16, 15, 87, 27, 18, 55, 24, 39, 5, 6, 59, 99, 40, 21, 38, 95, 33, 78, 28, 58, 84, 49, 50, 53, 64, 68, 14, 85, 77, 57, 26, 93, 74, 66, 61, 88, 22, 71, 67, 70, 1, 94, 54, 29, 51, 9, 13, 79, 91, 60, 8, 52, 3, 32, 81, 42, 35, 23, 75, 11, 25, 36, 47, 96, 10, 4, 2, 92, 76, 63, 17, 48, 73, 72, 30, 37, 20, 97, 69, 46, 56, 34, 82, 83, 44, 100, 90, 0, 86}
	polyResult := multiply(poly1, poly2)
	fmt.Println(minify(polyResult))
}

func multiply(poly1 []int, poly2 []int) []int {
	result := make([]int, len(poly1)+len(poly2))
	for i, coefficient := range poly1 {
		for j, coefficient2 := range poly2 {
			result[i+j] += coefficient * coefficient2
		}
	}
	return result
}

func minify(poly []int) []int {
	i := 0
	for poly[i] == 0 {
		i += 1
	}
	gcdPoly := poly[i]
	for _, coeff := range poly {
		if coeff == 0 {
			continue
		}
		gcdPoly = gcd(gcdPoly, coeff)
	}
	if gcdPoly == 1 {
		return poly
	}
	result := make([]int, len(poly))
	for i, coeff := range poly {
		result[i] = coeff / gcdPoly
	}
	return result
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

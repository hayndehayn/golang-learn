package main

import (
	"fmt"
	"math"
)

func main() {
	var height, weight float64
	fmt.Print("Input your height in cm: ")
	fmt.Scan(&height)
	fmt.Print("Input your weight in kg: ")
	fmt.Scan(&weight)
	result := showResult(calculateIMT(height, weight))
	fmt.Print(result)
}

func calculateIMT(height, weight float64) float64 {
	const IMTPower = 2
	heightInMeters := height / 100
	IMT := weight / math.Pow(heightInMeters, IMTPower)
	return IMT
}

func showResult(IMT float64) string {
	return fmt.Sprintf("Your index is %.2f", IMT)
}

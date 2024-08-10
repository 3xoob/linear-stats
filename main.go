package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Please provide a file path as an argument.")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: Could not open file: %v\n", err)
		os.Exit(0)
	}
	defer file.Close()

	var yValues []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("ERROR: Could not parse float: %v\n", err)
			os.Exit(0)
		}
		yValues = append(yValues, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("ERROR: Could not read file: %v\n", err)
		os.Exit(0)
	}

	n := float64(len(yValues))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range yValues {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}

	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / n

	r := (n*sumXY - sumX*sumY) / math.Sqrt((n*sumX2-sumX*sumX)*(n*sumY2-sumY*sumY))

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

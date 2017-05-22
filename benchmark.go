package main

import (
	"fmt"
	"github.com/skelterjohn/go.matrix"
	"math/rand"
	"time"
	"math"
)

func createDenseMatrix(source []float64, scale int) *matrix.DenseMatrix {
	return matrix.MakeDenseMatrix(source, scale, scale)
}

type R struct {
	rows, workers int
	runtime int64
	result float64
}

func worker(m * matrix.DenseMatrix, c chan<- float64, startRow, endRow int) {
	_, cols := m.GetSize()
	var sum float64
	for i := startRow; i < endRow; i++ {
		for j := 0; j < cols; j++ {
			v := m.Get(i, j)
			sum += v * v
		}
	}
	c<- sum
} 

func infinityNormParallel(m * matrix.DenseMatrix, w int) float64 {
	gather := make(chan float64, w * 2)
	rows, _ := m.GetSize()
	step := rows / w
	for i := 0; i < w; i++ {
		start, stop := step * i, step * (i + 1)
		go worker(m, gather, start, stop)
	}
	
	var sum float64
	for i := 0; i < w; i++ {
		sum += <-gather
	}
	return math.Sqrt(sum)
}

func benchmarkMatrixNormFunction(m *matrix.DenseMatrix, w int) R {
	start := time.Now().UnixNano()
	var result float64
	switch w {
		case 1:
			result = m.TwoNorm()
		default:
			result = infinityNormParallel(m, w)
	}
	stop := time.Now().UnixNano()
	
	rows, _ := m.GetSize()
	return R{rows, w, stop - start, result}
}

func main() {
	scale := 1024 * 8
	rand.Seed(time.Now().UTC().UnixNano())
	array := make([]float64, scale * scale)
	for i, _ := range array {
		array[i] = rand.Float64()
	}
	
	rows := []int{16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 1024 * 8}
	workers := []int{1, 4, 8, 16}
	
	for _, r := range rows {
		for _, w := range workers {
			m := createDenseMatrix(array, r)
			fmt.Println(benchmarkMatrixNormFunction(m, w))
		}
	}
}

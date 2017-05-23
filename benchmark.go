package goplayround

import (
	"github.com/skelterjohn/go.matrix"
	"math"
	"math/rand"
	"time"
)

func createDenseMatrixSpawner(maxScale int) (c func(scale int) *matrix.DenseMatrix) {
	rand.Seed(time.Now().UTC().UnixNano())
	array := make([]float64, maxScale * maxScale)
	for i := range array {
		array[i] = rand.Float64()
	}

	c = func(s int) *matrix.DenseMatrix {
		return matrix.MakeDenseMatrix(array, s, s)
	}
	return
}

type R struct {
	rows, workers int
	runtime       int64
	result        float64
}

func worker(m *matrix.DenseMatrix, c chan<- float64, startRow, endRow int) {
	_, cols := m.GetSize()
	var sum float64
	for i := startRow; i < endRow; i++ {
		for j := 0; j < cols; j++ {
			v := m.Get(i, j)
			sum += v * v
		}
	}
	c <- sum
}

func TwoNormParallel(m *matrix.DenseMatrix, w int) float64 {
	gather := make(chan float64, w*2)
	rows, _ := m.GetSize()
	step := rows / w
	for i := 0; i < w; i++ {
		start, stop := step*i, step*(i+1)
		go worker(m, gather, start, stop)
	}

	var sum float64
	for i := 0; i < w; i++ {
		sum += <-gather
	}
	return math.Sqrt(sum)
}


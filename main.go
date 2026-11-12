package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует слайс целых положительных чисел заданного размера.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = rand.Intn(1_000_000) + 1
	}

	return elements
}

// maximum находит максимальное число в слайсе (однопоточный вариант).
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, val := range data {
		if val > valGreaterThan(val, maxVal) {
			maxVal = val
		}
	}
	// Упрощенный цикл для поиска максимума:
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

// Вспомогательная функция для чистой читаемости
func valGreaterThan(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maxChunks находит максимальное значение, разделяя слайс на 8 частей в горутинах.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	chunkMaxs := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data) // захватываем остаток элементов на последнем чанке
		}

		go func(idx int, chunk []int) {
			defer wg.Done()
			chunkMaxs[idx] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(chunkMaxs)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел...\n", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println("Генерация завершена.")

	fmt.Println("Ищем максимальное значение в один поток...")
	startSingle := time.Now()
	singleMax := maximum(data)
	elapsedSingle := time.Since(startSingle).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms (мкс)\n\n", singleMax, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков...\n", CHUNKS)
	startMulti := time.Now()
	multiMax := maxChunks(data)
	elapsedMulti := time.Since(startMulti).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms (мкс)\n", multiMax, elapsedMulti)
}

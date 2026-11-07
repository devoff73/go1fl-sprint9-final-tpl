package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements создаёт слайс заданного размера и заполняет его
// случайными целыми положительными числами. Для размера <= 0 возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		// rand.Intn возвращает значение из [0, n), поэтому +1 даёт [1, math.MaxInt].
		elements[i] = rand.Intn(math.MaxInt) + 1
	}

	return elements
}

// maximum возвращает максимальный элемент слайса.
// Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks возвращает максимальный элемент слайса: делит его на CHUNKS частей
// и ищет максимум каждой части в отдельной горутине.
// Для пустого слайса возвращает 0.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Если элементов меньше, чем частей, делить слайс нет смысла.
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть забирает остаток от деления.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(chunkIndex int, chunk []int) {
			defer wg.Done()
			chunkMaximums[chunkIndex] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", maxValue, elapsed)
}

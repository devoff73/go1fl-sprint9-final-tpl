package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name         string
		size         int
		expectedSize int
	}{
		{name: "negative size", size: -5, expectedSize: 0},
		{name: "zero size", size: 0, expectedSize: 0},
		{name: "one element", size: 1, expectedSize: 1},
		{name: "many elements", size: 1000, expectedSize: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elements := generateRandomElements(tt.size)

			if elements == nil {
				t.Fatalf("generateRandomElements(%d) returned nil, want non-nil slice", tt.size)
			}
			if len(elements) != tt.expectedSize {
				t.Fatalf("generateRandomElements(%d) length = %d, want %d", tt.size, len(elements), tt.expectedSize)
			}
			for i, value := range elements {
				if value <= 0 {
					t.Errorf("element %d = %d, want positive number", i, value)
				}
			}
		})
	}
}

var maximumTestCases = []struct {
	name     string
	data     []int
	expected int
}{
	{name: "nil slice", data: nil, expected: 0},
	{name: "empty slice", data: []int{}, expected: 0},
	{name: "one element", data: []int{42}, expected: 42},
	{name: "max at the beginning", data: []int{9, 1, 2, 3}, expected: 9},
	{name: "max in the middle", data: []int{1, 2, 9, 3}, expected: 9},
	{name: "max at the end", data: []int{1, 2, 3, 9}, expected: 9},
	{name: "equal elements", data: []int{5, 5, 5, 5}, expected: 5},
	{name: "negative numbers", data: []int{-7, -3, -10}, expected: -3},
	{name: "fewer elements than chunks", data: []int{3, 8, 1}, expected: 8},
	{name: "exactly chunks elements", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 8},
	{name: "max in remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100}, expected: 100},
}

func TestMaximum(t *testing.T) {
	for _, tt := range maximumTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.expected {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, tt := range maximumTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.expected {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.expected)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(100_003)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Errorf("maxChunks = %d, maximum = %d", got, want)
	}
}

package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{
			name:    "нормальный размер",
			size:    100,
			wantLen: 100,
		},
		{
			name:    "нулевой размер",
			size:    0,
			wantLen: 0,
		},
		{
			name:    "отрицательный размер",
			size:    -10,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Errorf("generateRandomElements() длина = %d, ожидалась %d", len(got), tt.wantLen)
			}
			for _, v := range got {
				if v <= 0 {
					t.Errorf("generateRandomElements() сгенерировал неположительное число: %d", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "обычный слайс",
			data: []int{3, 9, 1, 4, 7},
			want: 9,
		},
		{
			name: "слайс с одним элементом",
			data: []int{42},
			want: 42,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "слайс с отрицательными числами",
			data: []int{-5, -2, -10},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "большой слайс для разделения на чанки",
			data: []int{10, 50, 20, 100, 5, 30, 40, 80, 90, 15},
			want: 100,
		},
		{
			name: "маленький слайс (меньше CHUNKS)",
			data: []int{3, 7, 2},
			want: 7,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

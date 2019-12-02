package main

import (
	"reflect"
	"testing"
)

func Test_Compute(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"test1", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"test2", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"test3", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"test4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Compute(tt.input)
			if false == reflect.DeepEqual(tt.input, tt.output) {
				t.Errorf("compute() = %v, want %v", tt.input, tt.output)
			}
		})
	}
}

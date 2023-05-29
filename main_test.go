package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	var tests = []struct {
		name     string
		eaters   []int
		menus    []int
		expected string
	}{
		{"Test 1", []int{1, 2, 3}, []int{1, 2, 3}, "First:1 Second:2 Third:3"},
		{"Test 2", []int{1, 1, 2, 2, 3}, []int{1, 2, 2, 3, 4}, "First:2 Second:1 Third:3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.expected

			output := calculateTop3Menus(tt.eaters, tt.menus)
			if output != expected {
				t.Errorf("Expected %s but got %s", expected, output)
			}
		})
	}
}

//This code tests the main function of the program using three different test cases.
// The first test case has all 10 menu items consumed once and expects the top 3 menu items to be returned.
// The second test case has only 9 menu items consumed once and expects the top 3 menu items to be returned.
// The third test case has only 8 menu items consumed once and expects the top 3 menu items to be returned.

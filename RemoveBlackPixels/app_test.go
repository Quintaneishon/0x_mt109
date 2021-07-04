package main

import "testing"

func TestRemove1(t *testing.T) {
	a := [][]int{  
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
    x := removeBlackPixels(a)
    if x != 7 {
		t.Errorf("Deben ser 7 los pixeles en 1")
    }
}

func TestRemove2(t *testing.T) {
	a := [][]int{  
		{1, 0, 1, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
    x := removeBlackPixels(a)
    if x != 9 {
		t.Errorf("Deben ser 9 los pixeles en 1")
    }
}
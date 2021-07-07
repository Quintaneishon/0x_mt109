package main

import "fmt"

type Pixel struct {
	x int
	y int
	hold bool
}

func contains( res []Pixel, i int , j int ) bool {
	for _, p := range res {
        if p.x == i && p.y == j {
            return true
        }
    }
    return false
}

func removeBlackPixels( a[][] int ) [][]int {
	n :=  len(a) // rows
	m :=  len(a[0]) // columns
	var res []Pixel // Array with black pixels
	
	// go through the matrix to get the edges and get the black pixels
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 1 && (i == 0 || i == n-1 || j == 0 || j == m-1)  {
				move(i,j,a,&res,n,m)
			}
		}
	}
	
	// create new matrzix with black pixels removed
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if a[i][j] == 1 && !contains(res, i, j){
				a[i][j] = 0
			}
		}
	}

	printMatrix(a)
	return a
}

func move( x int, y int, a [][]int, res *[]Pixel, n int , m int) {
	// izquierda
	if x - 1 > 0 {
		if a[x-1][y] == 1 && !contains(*res, x-1, y){ // avoid overflow
			var p Pixel
			p.x = x-1
			p.y = y
			p.hold = true
			*res = append(*res, p)
			
			move( x-1, y, a, res, n, m)
		}
	}
	// derecha
	if x + 1 < n - 1 {
		if a[x+1][y] == 1 && !contains(*res, x+1, y){ // avoid overflow
			var p Pixel
			p.x = x+1
			p.y = y
			p.hold = true
			*res = append(*res, p)
			move( x+1, y, a, res, n, m)
		}
	}
	// arriba
	if y - 1 > 0 {
		if a[x][y-1] == 1 && !contains(*res, x, y-1){ // avoid overflow
			var p Pixel
			p.x = x
			p.y = y-1
			p.hold = true
			*res = append(*res, p)
			move( x, y-1, a, res, n, m)
		}
	}
	// abajo
	if y + 1 < m - 1 {
		if a[x][y+1] == 1 && !contains(*res, x, y+1){ // avoid overflow
			var p Pixel
			p.x = x
			p.y = y+1
			p.hold = true
			*res = append(*res, p)
			move( x, y+1, a, res, n, m)
		}
	}
}

// print the matrix like the doc
func printMatrix( a[][] int ) {
	n :=  len(a[0]) // columns

	for i:=0; i < n; i++{
		fmt.Println(a[i])
		} 
}

func main() {
	a := [][]int{  
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
	printMatrix(a)
	removeBlackPixels(a)
}
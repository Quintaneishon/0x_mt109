package main

import "fmt"

func removeBlackPixels( a[][] int ) int {
	n :=  len(a[0]) // columns
	m :=  len(a) // rows
	var x,y []int // coordinates of edges with value 1
	
	// go through the matrix to get the edges and append to x and y
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && a[i][j] == 1{
				x = append(x,i)
				y = append(y,j)
			} else if i == n-1 && a[i][j] == 1{
            	x = append(x,i)
				y = append(y,j)
			} else if j == 0 && a[i][j] == 1{
				x = append(x,i)
				y = append(y,j)
			} else if j == m-1 && a[i][j] == 1{
            	x = append(x,i)
				y = append(y,j)
			}
		}
	}
	// go through the matrix to get the edges
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// not the edges
			if i!= 0 && i != n-1 && j!=0 && j!=m-1{
				// if value is 1
				if a[i][j] == 1{
					// go through the coordinates
					for k := 0; k < len(x); k++ {
						// if is conected, add to the coordinates slice, and break the loop
						if ( i-1 == x[k] && j == y[k]) || ( i == x[k] && j-1 == y[k]) || ( i == x[k] && j+1 == y[k]) || ( i+1 == x[k] && j == y[k]){
							x = append(x,i)
							y = append(y,j)
							a[i][j] = 1
							break
						} else {
							// if not connected, change the value to 0 
							a[i][j] = 0
						}
					}
				}
			}	
		}
	}
	printMatrix(a)
	return len(x)
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
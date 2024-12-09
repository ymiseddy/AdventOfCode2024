package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println("Hello World!")

	a := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	fmt.Println(a.RowView(0))

	col := a.ColView(0)
	row := a.RowView(0)
	fmt.Println(row.Dims())
	fmt.Println(col.Dims())
	/*
		b := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		fmt.Println(b)
		fmt.Println(a)
	*/

}

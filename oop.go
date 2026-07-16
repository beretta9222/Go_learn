package main

import "fmt"

type Figure struct {
	h int
	w int
}

type figure interface {
	getSquare() int
	getPerimetr() int
}

func (fig figure) getSquare(f Figure) (int, error) {
	return f.h * f.w, nil
}

func (f Figure) getPerimetr() int {
	return 2 * (f.h + f.w)
}

func main() {
	var f figure
	var r = f.getPerimetr()
	fmt.Println(r)
}

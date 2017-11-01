// Implement Pic. It should return a slice of length dy, each element of which
// is a slice of dx 8-bit unsigned integers. When you run the program, it will
// display your picture, interpreting the integers as grayscale (well,
// bluescale) values.

// The choice of image is up to you. Interesting functions include (x+y)/2, x*y,
// and x^y.
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// (Use uint8(intValue) to convert between types.

package main

import (
	"golang.org/x/tour/pic"
)

// Pic generates a slice of slices of uint8
func Pic(dx, dy int) [][]uint8 {
	myImage := make([][]uint8, dy)
	for i := range myImage {
		myImage[i] = make([]uint8, dx)
		for j := range myImage[i] {
			// myImage[i][j] = uint8(i + j/2)
			// myImage[i][j] = uint8(i ^ j)
			myImage[i][j] = uint8(i * j)
		}
	}

	return myImage
}

func main() {
	pic.Show(Pic)
	// myPic := Pic(256, 256)
	// fmt.Println(myPic)
	// fmt.Printf("%T\n", myPic)
}

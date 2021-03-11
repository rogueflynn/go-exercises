package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)
	for y := 0; y < dy; y++ {
		pic := make([]uint8, 0)
		for x := 0; x < dx; x++ {
			pic = append(pic, uint8(y+1)*uint8(x+1))
		}
		picture[y] = append(picture[y], pic...)
	}
	return picture
}

func main() {
	// pic.Show(Pic)
	fmt.Println(Pic(2,2))
}
package main

func Pic(dx, dy int) [][]uint8 {
	pict := make([][]uint8, dy)
	for i := range pict {
		pict[i] = make([]uint8, dx)
		for j := range pict[i] {
			pict[i][j] = uint8((i + j) / 2)
		}
	}
	return pict
}

//func main() {
//    // This function receives a function as a parameter
//    pic.Show(Pic)
//}

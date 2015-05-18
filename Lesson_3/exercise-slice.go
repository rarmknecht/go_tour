package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    data := make([][]uint8, dy)
    for n := range data {
        data[n] = make([]uint8, dx)
    }
    
    for i := range data {
        for j := range data[i] {
            data[i][j] = uint8((i+j)/5)
        }
    }
    
    return data
}

func main() {
    pic.Show(Pic)
}

package main

import (
	"code.google.com/p/go-tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {
    rand.Seed(37)
    rtn := make([][]uint8, dx)
    for i := range rtn {
        rtn[i] = make([]uint8, dy)
    }
    for i:=1; i<dx; i++ {
        for j:=1; j<dy; j++ {
            rtn[i][j] = uint8(rand.Intn(255))
        }
    }
    return rtn
}

func main() {
    pic.Show(Pic)
}




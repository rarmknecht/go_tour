package main

import (
    "fmt"
    "math"
)

func main() {
    val := 4.0
    for ; val <= 36; val++ {
        fmt.Printf("Value: %.4g\tResult: %.4g\tMath: %.4g\n", val, mysqrt(val), math.Sqrt(val))
    }
}

func mysqrt(x float64) float64 {
    // Arbitrary z value
    z := 1.5

    for i:=1; i <= 10; i++ {
        z = calc_z(z,x)
    }
    return z
}

func calc_z(z, x float64) float64 {
    // Newton's method
    newz := z - (((z*z) - x)/(2*z))
    return newz
}

package main

import (
	"fmt"
	"math"
	//"os"
)

func Sqrt(x float64) (z float64) {
	z = 1
	//var z0 float64
	
	for {
		z = z - (z * z -x)/2*z
		/*if math.Abs(z0 - z) < 0.001 {
			return
		}*/
		if math.Abs(math.Sqrt(x) - z) < 0.0001 {
			return
		}
		
		//z0 = z
		fmt.Println(z)
	}
	return
}

func main() {
	Sqrt(2)
}

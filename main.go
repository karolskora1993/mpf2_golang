package main

import (
	"fmt"
	"math"
)

const (
	temp    = 820
	temp_K  = temp + 273
	v_nagrz = 0.5
	Q       = 140000
	R       = 8.3144
	d0      = 0.000041
	dx      = 0.1
	dt      = 0.001
	C_ALFA  = (temp - 912) / -240
	LENGTH = 100
)

var (
	D = d0 * math.Exp(-Q/(R*temp_K)) * 10000000000
	ksi = 0
)


func main() {
	var c1 [LENGTH]float64

	for i:=0 ; i<100 ; i++ {
		if i<6 {
			c1[i] = 0.67
		} else if i==6 {
			c1[i] = 0.67
			ksi = i
		} else {
			c1[i] = 0.02
		}
	}

	var c2 = copy(c1)

	steps := 1000

	for i:=0 ; i<steps ; i++ {
		for j:=0 ; j<=ksi ; j++ {
			if j==0 {
				c2[j] = (1 - 2 * (D * dt / math.Pow(dx, 2))) * c1[j] + (D * dt / math.Pow(dx, 2)) * (c1[j] + c1[j + 1])
			} else if j==ksi {
				c2[j] = (1 - 2 * (D * dt / math.Pow(dx, 2))) * c1[j] + (D * dt / math.Pow(dx, 2)) * (c1[j - 1] + c1[j])

			} else {
				c2[j] = (1 - 2 * (D * dt / math.Pow(dx, 2))) * c1[j] + (D * dt / math.Pow(dx, 2)) * (c1[j - 1] + c1[j + 1])
			}
	}
		if ksi<len(c1)-1 {
			c2[ksi + 1] = c2[ksi]
		}
		c1 = copy(c2)

		if c1[ksi]>=C_ALFA && ksi<len(c1)-1 {
			ksi+=1
		}

	}
	fmt.Println(c1)
}

func copy(a [LENGTH]float64) [LENGTH]float64 {
	var b [LENGTH]float64
	for i:=0; i<LENGTH ; i++ {
		b[i] = a[i]
	}
	return b
}

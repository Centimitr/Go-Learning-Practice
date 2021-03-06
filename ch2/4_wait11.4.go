//for the sake of implicity, this practice write in single .go file
package main

import ()

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>0*8)] +
		pc[byte(x>>1*8)] +
		pc[byte(x>>2*8)] +
		pc[byte(x>>3*8)] +
		pc[byte(x>>4*8)] +
		pc[byte(x>>5*8)] +
		pc[byte(x>>6*8)] +
		pc[byte(x>>7*8)])
}

func PopCountStep(x uint64) int {
	cnt := 0
	for x > 0 {
		if (x & 1) == 1 {
			cnt++
		}
		x >>= 1
	}
	return cnt
}

func main() {
}

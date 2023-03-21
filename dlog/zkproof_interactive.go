package main

import (
	"fmt"
	"math"
	"math/rand"
)

//func dlogProof(x *big.Int, g *big.Int, p *big.Int) big.Int {
//	fmt.Println("generating proof")
//
//	b := big.NewInt(0)
//	b.Exp(g, x, p)
//	return big.Int{}
//}
//
//func dlogVerify(y big.Int, g big.Int, p big.Int, pf big.Int) bool {
//	fmt.Println("verifying")
//	return false
//}

func dlogProof(x uint64, g uint64, p uint64, b uint64) (float64, float64, uint64) {
	// x: int
	// g: int
	// p: int
	// return: [int, int]
	// ZK proof of knowledge of x such that g^x = y mod p

	fmt.Println("generating proof")

	r := math.Floor(rand.Float64() * float64(p-1))
	t := math.Mod(math.Pow(float64(g), r), float64(p))
	s := r + float64(b*x)
	return t, s, b
}

func dlogVerify(y uint64, g uint64, p uint64, t uint64, s uint64, b uint64) bool {
	// y: int
	// g: int
	// p: int
	// t: int
	// s: int
	// Verify ZK proof (t,s) of knowledge x such that g^x = y mod p
	fmt.Println("verifying")
	lhs := math.Mod(math.Pow(float64(g), float64(s)), float64(p))
	if b != 1 {
		y = uint64(1)
	}
	rhs := math.Mod(float64(y)*float64(t), float64(p))

	if lhs != rhs {
		return false
	}
	return true
}

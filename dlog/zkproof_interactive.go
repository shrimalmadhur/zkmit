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

func main() {
	//prime := big.NewInt(0)
	//_, ok := prime.SetString("215115837579444693413581944743", 10)
	//if !ok {
	//	fmt.Println("something is wrong")
	//}
	//
	//fmt.Println(prime.String())
	//
	//generator := big.NewInt(0)
	//_, ok = generator.SetString("182319915291536251337341050091", 10)
	//if !ok {
	//	fmt.Println("something is wrong")
	//}
	//
	//fmt.Println(generator.String())
	//
	//primeMinusOne := big.NewInt(0)
	//primeMinusOne.Sub(prime, big.NewInt(1))
	//primeMinusOne.Mul(primeMinusOne, rand.Float64())
	//fl := big.NewFloat(rand.Float64())
	//big.r
	//fl.Mul(fl, primeMinusOne)
	//r := math.Floor(rand.Float64() * primeMinusOne)

	p := uint64(31)
	g := uint64(3)
	x := uint64(17)
	y := uint64(22)
	// Let's prove we know x such that g^x = y mod p

	b := uint64(1)
	if rand.Float64() <= 0.5 {
		b = 0
	}
	t, s, _ := dlogProof(x, g, p, b)

	goodResult := dlogVerify(y, g, p, uint64(t), uint64(s), b)
	fmt.Printf("this should be true: %v\n", goodResult)

	badResult := dlogVerify(uint64(23), g, p, uint64(t), uint64(s), b)
	fmt.Printf("this should be false: %v\n", badResult)
}

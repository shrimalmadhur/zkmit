package main

import (
	"fmt"
	"math/rand"
)

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

	fmt.Println("----------------------Running Interactive DLOG ZK proof----------------------")
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
	fmt.Println("----------------------End Interactive DLOG ZK proof----------------------")
	fmt.Println()
	fmt.Println("----------------------Running Non Interactive DLOG ZK proof----------------------")

	fmt.Println("----------------------End Non Interactive DLOG ZK proof----------------------")
}

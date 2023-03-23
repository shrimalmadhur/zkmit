// https://en.wikipedia.org/wiki/Fiat%E2%80%93Shamir_heuristic

package main

import "fmt"

func dlogProof1(x uint64, g uint64, p uint64) (float64, float64) {
	fmt.Println("generating proof")

	return 0, 0
}

func dlogVerify1(y uint64, g uint64, p uint64, t uint64, s uint64) bool {
	fmt.Println("verifying")
	return false
}

package challenge57

import (
	"fmt"
	"math/big"

	"github.com/davecgh/go-spew/spew"

	"github.com/MashaSamoylova/DH/pkg/cipher"
	"github.com/MashaSamoylova/DH/pkg/tools"
)

var one = big.NewInt(1)
var zero = big.NewInt(0)

// https://toadstyle.org/cryptopals/57.txt
func Attack(P,G,Q *big.Int, victim cipher.DiffieHellman) (*big.Int, *big.Int) {
	var err error

	J := new(big.Int)

	tmp := new(big.Int)
	P_1 := new(big.Int)



	P_1.Sub(P, one)
	J.Div(P_1, Q)
	rS := factor(J)
	bS := make([]*big.Int, len(rS))

	for i, r := range rS {
		H := big.NewInt(1)
		for H.Cmp(one) == 0 {
			A := cipher.GenerateWithMod(P)
			tmp.Div(P_1, r)
			H.Exp(A, tmp, P)
		}
		victim.GenerateSessionKey(H)
		newB := findB(H, P, r, victim.SessionKey)
		bS[i] = newB
	}

	fmt.Println("Rs:")
	spew.Dump(rS)
	fmt.Println("Bs:")
	spew.Dump(bS)

	n, r,  err := tools.CRT(bS, rS)
	if err != nil {
		fmt.Println(err)
		panic("CRT failed")
	}
	fmt.Println("Bob's private: ")
	spew.Dump(victim.PrivateKey)
	fmt.Println("Found: ")
	spew.Dump(n)
	return n, r
}

func findB(H, P, r, SessionKey *big.Int) *big.Int {
	result := new(big.Int)
	var i uint64
	fmt.Println(r.Uint64())
	for i = 0; i <= r.Uint64(); i++ {
		x := big.NewInt(int64(i))
		result.Exp(H, x, P)
		if SessionKey.Cmp(result) == 0 {
			return x
		}
	}
	panic("Not Found")
}

func factor(A *big.Int) []*big.Int {
	tmp := new(big.Int)
	rem := new(big.Int)
	var rS []*big.Int

	for i := int64(2); i < 0x100000; i++ {
		r := big.NewInt(i)
		if r.ProbablyPrime(10) == false {
			continue
		}
		tmp.DivMod(A, r, rem)
		if rem.Cmp(zero) == 0 {
			rS = append(rS, r)
		}
	}
	return rS
}


func main() {
	G := new(big.Int)
	P := new(big.Int)
	Q := new(big.Int)

	G.SetString("4565356397095740655436854503483826832136106141639563487732438195343690437606117828318042418238184896212352329118608100083187535033402010599512641674644143", 10)
	P.SetString("7199773997391911030609999317773941274322764333428698921736339643928346453700085358802973900485592910475480089726140708102474957429903531369589969318716771", 10)

	Q.SetString("236234353446506858198510045061214171961", 10)

	privateKey := cipher.GenerateWithMod(Q)
	Bob := cipher.NewDiffieHellman(privateKey, P, G)
	Attack(P, G, Q, Bob)
}
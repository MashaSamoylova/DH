package cipher

import (
	"crypto/rand"
	"math/big"
)

func GenerateNum(bits int64) *big.Int {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(bits), nil).Sub(max, big.NewInt(1))

	k, err := rand.Int(rand.Reader, max)
	if err != nil {
		//error handling
	}
	return k
}

func GeneratePrivateKey(p, g, private *big.Int) *big.Int {
	z := new(big.Int)
	(*z).Exp(g, private, p)
	return z
}

func GenerateSessionKey(p, a, private *big.Int) *big.Int {
	z := new(big.Int)
	(*z).Exp(a, private, p)
	return z
}

package tools

import (
	"fmt"
	"math"
	"math/big"
)

// f(y) = 2^y (mod k)
func f(y, k, p *big.Int) *big.Int {
	two := big.NewInt(2)
	mod := new(big.Int).Mod(y, k)
	return new(big.Int).Exp(two, mod, p)
}

// https://toadstyle.org/cryptopals/58.txt
func Kangaroo(g, y, a, b, p *big.Int) *big.Int {
	aSmall := a.Int64()
	bSmall := b.Int64()
	kSmall := math.Log(float64(bSmall-aSmall)) / math.Log(4)
	k := big.NewInt(int64(kSmall))
	N := 4 * ((int(math.Pow(float64(2), kSmall)) - 1) / int(kSmall))

	xT := big.NewInt(0)
	yT := new(big.Int).Exp(g, b, p)

	for i := 0; i < N; i++ {
		step := f(yT, k, p)
		xT.Add(xT, step)
		yT.Mul(yT, new(big.Int).Exp(g, step, p))
		yT.Mod(yT, p)
	}

	xW := big.NewInt(0)
	yW := new(big.Int).Set(y)

	max := new(big.Int).Sub(b, a)
	max.Add(max, xT)

	for xW.Cmp(max) == -1 {
		step := f(yW, k, p)
		xW.Add(xW, step)
		yW.Mul(yW, new(big.Int).Exp(g, step, p))
		yW.Mod(yW, p)

		if yW.Cmp(yT) == 0 {
			fmt.Println(yW)
			fmt.Println(yT)
			b.Add(b, xT)
			b.Sub(b, xW)
			if b.Cmp(big.NewInt(0)) == -1 {
				b.Add(b, p)
			}
			return b
		}
	}
	fmt.Println("Kangaroo failed")
	return big.NewInt(0)
}

package main

import (
	"math/big"

	"github.com/MashaSamoylova/DH/cmd/challenge57"
	"github.com/MashaSamoylova/DH/pkg/cipher"
	"github.com/MashaSamoylova/DH/pkg/tools"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	P, _ := new(big.Int).SetString("11470374874925275658116663507232161402086650258453896274534991676898999262641581519101074740642369848233294239851519212341844337347119899874391456329785623", 10)
	G, _ := new(big.Int).SetString("622952335333961296978159266084741085889881358738459939978290179936063635566740258555167783009058567397963466103140082647486611657350811560630587013183357", 10)
	Q, _ := new(big.Int).SetString("335062023296420808191071248367701059461", 10)

	// m = 132167685656.0
	privateKey, _ := new(big.Int).SetString("59218903646839673978326597577070215555", 10)
	Bob := cipher.NewDiffieHellman(privateKey, P, G)

	n, r := challenge57.Attack(P, G, Q, Bob)

	g1 := new(big.Int).Exp(G, r, P)

	rightB := new(big.Int).Sub(Q, big.NewInt(1))
	rightB.Div(rightB, r)
	rightB.Mod(rightB, P)

	GNReverse := new(big.Int).Exp(G, n, P)
	GNReverse.ModInverse(GNReverse, P)

	y1 := new(big.Int).Mul(Bob.PublicKey, GNReverse)
	y1.Mod(y1, P)

	m := tools.Kangaroo(g1, y1, big.NewInt(0), rightB, P)

	privateBob := new(big.Int).Mul(m, r)
	privateBob.Mod(privateBob, P)
	privateBob.Add(privateBob, n)
	privateBob.Mod(privateBob, P)
	spew.Dump(privateBob)
}

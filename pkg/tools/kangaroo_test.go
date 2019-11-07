package tools

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKangaroo(t *testing.T) {
	p, _ := new(big.Int).SetString("11470374874925275658116663507232161402086650258453896274534991676898999262641581519101074740642369848233294239851519212341844337347119899874391456329785623", 10)
	g, _ := new(big.Int).SetString("622952335333961296978159266084741085889881358738459939978290179936063635566740258555167783009058567397963466103140082647486611657350811560630587013183357", 10)

	y, _ := new(big.Int).SetString("7760073848032689505395005705677365876654629189298052775754597607446617558600394076764814236081991643094239886772481052254010323780165093955236429914607119", 10)
	fmt.Println("====================================================")
	fmt.Println("g = ", g)
	fmt.Println("y = ", y)
	x := Kangaroo(g, y, big.NewInt(0), big.NewInt(1<<20), p)
	fmt.Println("x = ", x)
	fmt.Println("check = ", new(big.Int).Exp(g, x, p))
	assert.Equal(t, y, new(big.Int).Exp(g, x, p))

	fmt.Println("====================================================")
	y, _ = new(big.Int).SetString("9388897478013399550694114614498790691034187453089355259602614074132918843899833277397448144245883225611726912025846772975325932794909655215329941809013733", 10)
	fmt.Println("g = ", g)
	fmt.Println("y = ", y)
	x = Kangaroo(g, y, big.NewInt(0), big.NewInt(1<<40), p)
	fmt.Println("x = ", x)
	fmt.Println("check = ", new(big.Int).Exp(g, x, p))
	assert.Equal(t, y, new(big.Int).Exp(g, x, p))
}

package main

import (
	"fmt"
	"math/big"

	"github.com/MashaSamoylova/DH/pkg/cipher"
)

// MITM key-fixing attack on Diffie-Hellman with parameter injection
// https://cryptopals.com/sets/5/challenges/34
func Challenge34() {
	G := new(big.Int)
	P := new(big.Int)
	G.SetString("A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5", 16)
	P.SetString("B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371", 16)

	Alice := cipher.NewDiffieHellmanWithRandom(P, G)
	Bob := cipher.NewDiffieHellmanWithRandom(P, G)
	Eva := cipher.NewDiffieHellmanWithRandom(P, G)

	// Malefactor swaps public key to P for getting zero session key.
	Alice.GenerateSessionKey(P)
	Bob.GenerateSessionKey(P)
	Eva.SessionKey = new(big.Int)
	Eva.SessionKey.SetString("0", 16)

	Alice.InitAES()
	Bob.InitAES()
	Eva.InitAES()

	encryptedMsg := Alice.EncryptMsg([]byte("Hello, Bob! :-))"))
	evaReceived := Eva.DecryptMsg(encryptedMsg)
	bobReceived := Bob.DecryptMsg(encryptedMsg)

	fmt.Println("Eva received:", string(evaReceived))
	fmt.Println("Bob received", string(bobReceived))

	fmt.Println("💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻💃🏻")
	encryptedMsg = Bob.EncryptMsg([]byte("Hello, Alice! :)"))
	evaReceived = Eva.DecryptMsg(encryptedMsg)
	aliceReceived := Alice.DecryptMsg(encryptedMsg)
	fmt.Println("Eva received:", string(evaReceived))
	fmt.Println("Alice received", string(aliceReceived))
}

func main() {
	Challenge34()
}

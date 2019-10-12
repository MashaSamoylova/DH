package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
)

type DiffieHellman struct {
	PrivateKey *big.Int
	PublicKey  *big.Int
	SessionKey *big.Int
	P          *big.Int
	G          *big.Int

	Cipher cipher.Block
}

func NewDiffieHellman(PrivateKey, P, G *big.Int) DiffieHellman {
	z := new(big.Int)
	(*z).Exp(G, PrivateKey, P)
	return DiffieHellman{
		PrivateKey: PrivateKey,
		PublicKey:  z,
		P:          P,
		G:          G,
	}
}

func NewDiffieHellmanWithRandom(P, G *big.Int) DiffieHellman {
	privateKey := generateNum(100)
	return NewDiffieHellman(privateKey, P, G)
}

func GenerateWithMod(M *big.Int) *big.Int {
	k, err := rand.Int(rand.Reader, M)
	if err != nil {
		panic("Failed to rand num -_0_0_")
	}
	return k
}

func (p *DiffieHellman) GenerateSessionKey(a *big.Int) {
	p.SessionKey = new(big.Int)
	(*p.SessionKey).Exp(a, p.PrivateKey, p.P)
}

func (p *DiffieHellman) InitAES() {
	sum := sha1.Sum(p.SessionKey.Bytes())
	c, err := aes.NewCipher(sum[:aes.BlockSize])
	if err != nil {
		fmt.Println("Failed to init AES")
		panic(err)
	}

	p.Cipher = c
}

func (p *DiffieHellman) EncryptMsg(msg []byte) (encryptedMsg []byte) {
	encryptedMsg = make([]byte, aes.BlockSize)
	p.Cipher.Encrypt(encryptedMsg, msg)
	return encryptedMsg
}

func (p *DiffieHellman) DecryptMsg(encryptedMsg []byte) (msg []byte) {
	msg = make([]byte, aes.BlockSize)
	p.Cipher.Decrypt(msg, encryptedMsg)
	return msg
}

func generateNum(bits int64) *big.Int {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(bits), nil).Sub(max, big.NewInt(1))

	k, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic("Failed to rand num -_0_0_")
	}
	return k
}

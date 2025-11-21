package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	key, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	key.Add(key, big.NewInt(2))
	return key
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	gBig := big.NewInt(g)
	pub := new(big.Int).Exp(gBig, private, p)
	return pub
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	key := PrivateKey(p)
	pub := PublicKey(key, p, g)
	return key, pub
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	secret := new(big.Int).Exp(public2, private1, p)
	return secret
}

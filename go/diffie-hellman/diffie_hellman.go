package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	max := big.NewInt(0).Set(p)             // Clone
	max.Sub(max, big.NewInt(2))             // max = p - 2
	priv, err := rand.Int(rand.Reader, max) // [0, p-2)
	if err != nil {
		panic(err)
	}
	return priv.Add(priv, big.NewInt(2)) // [0, p-2) + 2 == [2, p) == [2, p - 1]
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	pub := big.NewInt(g)
	return pub.Exp(pub, private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	pub := PublicKey(priv, p, g)
	return priv, pub
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	sec := big.NewInt(0).Set(public2)
	return sec.Exp(sec, private1, p)
}

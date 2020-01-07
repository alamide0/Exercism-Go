package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	x, err := rand.Int(rand.Reader, p)

	if err != nil {
		return big.NewInt(0)
	}

	if x.Cmp(big.NewInt(2)) < 0 {
		return PrivateKey(p)
	}

	return x
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

package diffiehellman

import (
    "math/big"
    "crypto/rand"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
    max:= new(big.Int).Sub(p, big.NewInt(2)) //shouldn't mutate p
	a, _:= rand.Int(rand.Reader, max)
    a.Add(a, big.NewInt(2))
    return a
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	gbig:= big.NewInt(g)
    z:= new(big.Int).Exp(gbig, private, p)
    return z
}
//return public-private key
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private:= PrivateKey(p)
    public:= PublicKey(private, p, g)
    return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}

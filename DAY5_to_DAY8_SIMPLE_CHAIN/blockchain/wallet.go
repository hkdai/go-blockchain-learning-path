package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubKey := append(privKey.PublicKey.X.Bytes(), privKey.Y.Bytes()...)
	address := GenerateAddress(pubKey)

	return &Wallet{
		PrivateKey: privKey,
		PublicKey:  pubKey,
		Address:    address,
	}
}

func GenerateAddress(pubKey []byte) string {
	pubHash := sha256.Sum256(pubKey)
	r := ripemd160.New()
	r.Write(pubHash[:])
	return hex.EncodeToString(r.Sum(nil))
}

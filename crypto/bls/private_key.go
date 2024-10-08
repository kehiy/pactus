package bls

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	bls12381 "github.com/kilic/bls12-381"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/util"
	"github.com/pactus-project/pactus/util/bech32m"
	"github.com/pactus-project/pactus/util/errors"
	"golang.org/x/crypto/hkdf"
)

var _ crypto.PrivateKey = &PrivateKey{}

const PrivateKeySize = 32

type PrivateKey struct {
	fr bls12381.Fr
}

// PrivateKeyFromString decodes the input string and returns the PrivateKey
// if the string is a valid bech32m encoding of a BLS public key.
func PrivateKeyFromString(text string) (*PrivateKey, error) {
	// Decode the bech32m encoded private key.
	hrp, typ, data, err := bech32m.DecodeToBase256WithTypeNoLimit(text)
	if err != nil {
		return nil, err
	}

	// Check if hrp is valid
	if hrp != crypto.PrivateKeyHRP {
		return nil, crypto.InvalidHRPError(hrp)
	}

	if typ != crypto.SignatureTypeBLS {
		return nil, errors.Errorf(errors.ErrInvalidPrivateKey,
			"invalid private key type: %v", typ)
	}

	return PrivateKeyFromBytes(data)
}

// KeyGen generates a private key deterministically from a secret octet string
// IKM and an optional octet string keyInfo.
// Based on https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-bls-signature-04#section-2.3
func KeyGen(ikm, keyInfo []byte) (*PrivateKey, error) {
	// L is `ceil((3 * ceil(log2(r))) / 16) = 48`,
	//    where `r` is the order of the BLS 12-381 curve
	//    r:  0x73eda753 299d7d48 3339d808 09a1d805 53bda402 fffe5bfe ffffffff 00000001
	// 	  https://datatracker.ietf.org/doc/html/draft-yonezawa-pairing-friendly-curves-02#section-4.2.2
	//

	if len(ikm) < 32 {
		return nil, fmt.Errorf("ikm is too short")
	}

	secret := make([]byte, 0, len(ikm)+1)
	secret = append(secret, ikm...)
	secret = append(secret, util.I2OSP(big.NewInt(0), 1)...)

	l := int64(48)
	pseudoRandomKey := make([]byte, 0, len(keyInfo)+2)
	pseudoRandomKey = append(pseudoRandomKey, keyInfo...)
	pseudoRandomKey = append(pseudoRandomKey, util.I2OSP(big.NewInt(l), 2)...)

	g1 := bls12381.NewG1()

	salt := []byte("BLS-SIG-KEYGEN-SALT-")
	x := big.NewInt(0)
	for x.Sign() == 0 {
		h := sha256.Sum256(salt)
		salt = h[:]

		okm := make([]byte, l)
		prk := hkdf.Extract(sha256.New, secret, salt)
		reader := hkdf.Expand(sha256.New, prk, pseudoRandomKey)
		_, _ = reader.Read(okm)

		r := g1.Q()
		x = new(big.Int).Mod(util.OS2IP(okm), r)
	}

	sk := make([]byte, 32)
	x.FillBytes(sk)

	return PrivateKeyFromBytes(sk)
}

// PrivateKeyFromBytes constructs a BLS private key from the raw bytes.
func PrivateKeyFromBytes(data []byte) (*PrivateKey, error) {
	if len(data) != PrivateKeySize {
		return nil, errors.Errorf(errors.ErrInvalidPrivateKey,
			"private key should be %d bytes, but it is %v bytes", PrivateKeySize, len(data))
	}

	fr := bls12381.NewFr()
	fr.FromBytes(data)
	if fr.IsZero() {
		return nil, errors.Errorf(errors.ErrInvalidPrivateKey,
			"private key is zero")
	}

	return &PrivateKey{fr: *fr}, nil
}

// String returns a human-readable string for the BLS private key.
func (prv *PrivateKey) String() string {
	str, _ := bech32m.EncodeFromBase256WithType(
		crypto.PrivateKeyHRP,
		crypto.SignatureTypeBLS,
		prv.Bytes())

	return strings.ToUpper(str)
}

// Bytes return the raw bytes of the private key.
func (prv *PrivateKey) Bytes() []byte {
	return prv.fr.ToBytes()
}

// Sign calculates the signature from the private key and given message.
// It's defined in section 2.6 of the spec: CoreSign.
func (prv *PrivateKey) Sign(msg []byte) crypto.Signature {
	return prv.SignNative(msg)
}

func (prv *PrivateKey) SignNative(msg []byte) *Signature {
	g1 := bls12381.NewG1()

	q, err := g1.HashToCurve(msg, dst)
	if err != nil {
		panic(err)
	}
	pointG1 := g1.MulScalar(g1.New(), q, &prv.fr)
	data := g1.ToCompressed(pointG1)

	return &Signature{
		data:    data,
		pointG1: pointG1,
	}
}

func (prv *PrivateKey) PublicKeyNative() *PublicKey {
	g2 := bls12381.NewG2()

	pointG2 := g2.MulScalar(g2.New(), g2.One(), &prv.fr)
	data := g2.ToCompressed(pointG2)

	return &PublicKey{
		data:    data,
		pointG2: pointG2,
	}
}

func (prv *PrivateKey) PublicKey() crypto.PublicKey {
	return prv.PublicKeyNative()
}

func (prv *PrivateKey) EqualsTo(x crypto.PrivateKey) bool {
	xBLS, ok := x.(*PrivateKey)
	if !ok {
		return false
	}

	return prv.fr.Equal(&xBLS.fr)
}

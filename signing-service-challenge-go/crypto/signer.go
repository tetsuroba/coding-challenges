package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

type RSASigner struct {
	PrivateKey    []byte        `json:"-" bson:"-"`
	PublicKey     []byte        `json:"-" bson:"-"`
	RSAMarshaller *RSAMarshaler `json:"-" bson:"-"`
}

type ECCSigner struct {
	PrivateKey   []byte        `json:"-" bson:"-"`
	PublicKey    []byte        `json:"-" bson:"-"`
	ECCMarshaler *ECCMarshaler `json:"-" bson:"-"`
}

// NewRSASigner creates a new RSASigner.
func NewRSASigner(pair *RSAKeyPair) (*RSASigner, error) {
	marshaller := NewRSAMarshaler()
	publicKey, privateKey, err := marshaller.Marshal(*pair)
	if err != nil {
		return nil, err
	}
	return &RSASigner{
		PublicKey:     publicKey,
		PrivateKey:    privateKey,
		RSAMarshaller: &marshaller,
	}, nil
}

// NewECCSigner creates a new ECCSigner.
func NewECCSigner(pair *ECCKeyPair) (*ECCSigner, error) {
	marshaller := NewECCMarshaler()
	publicKey, privateKey, err := marshaller.Encode(*pair)
	if err != nil {
		return nil, err
	}
	return &ECCSigner{
		PublicKey:    publicKey,
		PrivateKey:   privateKey,
		ECCMarshaler: &marshaller,
	}, nil
}

// Sign signs the dataToBeSigned using RSA
func (d *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	decodedKeys, err := d.RSAMarshaller.Unmarshal(d.PrivateKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(dataToBeSigned)
	signature, err := rsa.SignPKCS1v15(rand.Reader, decodedKeys.Private, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// Sign signs the dataToBeSigned using ECC
func (d *ECCSigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	decodedKeys, err := d.ECCMarshaler.Decode(d.PrivateKey)
	if err != nil {
		return nil, err
	}

	signature, err := ecdsa.SignASN1(rand.Reader, decodedKeys.Private, dataToBeSigned)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

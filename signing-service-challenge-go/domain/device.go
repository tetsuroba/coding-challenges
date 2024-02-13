package domain

import (
	"encoding/base64"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
	"strconv"
	"sync"
)

type Algorithm string

const (
	RSA   Algorithm = "RSA"
	ECDSA Algorithm = "ECDSA"
)

type Device struct {
	ID               uuid.UUID     `json:"id" bson:"ID"`
	Label            string        `json:"label,omitempty" bson:"Label,omitempty"`
	SignatureCounter int           `json:"signatureCounter" bson:"SignatureCounter"`
	LastSignature    string        `json:"lastSignature,omitempty" bson:"LastSignature,omitempty"`
	Algorithm        Algorithm     `json:"-" bson:"-"`
	Signer           crypto.Signer `json:"-" bson:"-"`
	LockDevice       sync.Mutex    `json:"-" bson:"-"`
}

func NewDevice(label string, algorithm Algorithm, signer crypto.Signer) (*Device, error) {
	deviceUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Device{
		ID:               deviceUUID,
		Label:            label,
		SignatureCounter: 0,
		Algorithm:        algorithm,
		Signer:           signer,
	}, nil
}

type DeviceRequest struct {
	ID        string    `json:"id" bson:"ID"`
	Label     string    `json:"label" bson:"Label"`
	Algorithm Algorithm `json:"algorithm" bson:"Algorithm"`
}

type SignRequest struct {
	DeviceId string `json:"deviceId" bson:"DeviceId"`
	Data     string `json:"data" bson:"Data"`
}

// SignTransaction signs the given data with the device's private key and returns the signature and the data to be signed
func (d *Device) SignTransaction(id string, signed string) (*string, *string, error) {
	d.LockDevice.Lock()
	defer d.LockDevice.Unlock()
	var lastSignature string
	if d.SignatureCounter == 0 {
		lastSignature = base64.StdEncoding.EncodeToString([]byte(id))
	} else {
		lastSignature = d.LastSignature
	}
	dataToBeSigned := strconv.Itoa(d.SignatureCounter) + "_" + signed + "_" + lastSignature
	signature, err := d.Signer.Sign([]byte(dataToBeSigned))
	if err != nil {
		d.LockDevice.Unlock()
		return nil, nil, err
	}
	d.SignatureCounter++
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	d.LastSignature = encodedSignature

	return &encodedSignature, &dataToBeSigned, nil
}

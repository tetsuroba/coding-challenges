package domain

import (
	"encoding/base64"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"testing"
)

var (
	testId       = "testId"
	base64TestId = base64.StdEncoding.EncodeToString([]byte(testId))
	rsaGenerator = crypto.RSAGenerator{}
)

func TestDevice_SignTransaction(t *testing.T) {
	rsaKeyPair, err := rsaGenerator.Generate()
	if err != nil {
		t.Fatal(err)
	}
	rsaSigner, err := crypto.NewRSASigner(rsaKeyPair)
	if err != nil {
		t.Fatal(err)
	}
	device, err := NewDevice("testDevice", RSA, rsaSigner)
	if err != nil {
		t.Fatal(err)
	}
	dataToBeSignedExpected := "0_testData_" + base64TestId
	encodedSignature, dataToBeSigned, err := device.SignTransaction("testId", "testData")
	if err != nil {
		t.Errorf("Failed to sign message %v", err)
	}
	if dataToBeSignedExpected != *dataToBeSigned {
		t.Errorf("Data to be signed is not correct was %s expected %s", *dataToBeSigned, dataToBeSignedExpected)
	}
	if encodedSignature == nil {
		t.Errorf("Signature is nil")
	}
}

func TestDevice_SignTransaction_Multiple(t *testing.T) {
	rsaKeyPair, err := rsaGenerator.Generate()
	if err != nil {
		t.Fatal(err)
	}
	rsaSigner, err := crypto.NewRSASigner(rsaKeyPair)
	if err != nil {
		t.Fatal(err)
	}
	device, err := NewDevice("testDevice", RSA, rsaSigner)
	if err != nil {
		t.Fatal(err)
	}
	dataToBeSignedExpected1 := "0_testData_" + base64TestId
	encodedSignature, dataToBeSigned, err := device.SignTransaction("testId", "testData")
	if err != nil {
		t.Errorf("Failed to sign message %v", err)
	}
	if dataToBeSignedExpected1 != *dataToBeSigned {
		t.Errorf("Data to be signed is not correct was %s expected %s", *dataToBeSigned, dataToBeSignedExpected1)
	}
	if encodedSignature == nil {
		t.Errorf("Signature is nil")
	}
	dataToBeSignedExpected2 := "1_testData_" + *encodedSignature
	encodedSignature, dataToBeSigned, err = device.SignTransaction("testId", "testData")
	if err != nil {
		t.Errorf("Failed to sign message %v", err)
	}
	if dataToBeSignedExpected2 != *dataToBeSigned {
		t.Errorf("Data to be signed is not correct was %s expected %s", *dataToBeSigned, dataToBeSignedExpected2)
	}
	if encodedSignature == nil {
		t.Errorf("Signature is nil")
	}
}

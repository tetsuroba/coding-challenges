package service

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/google/uuid"
	"testing"
)

var (
	testLabel  = "testLabel"
	testLabel2 = "testLabel2"
)

func TestDeviceService_CreateDevice_TestRSADevice(t *testing.T) {
	service := NewDeviceService(persistence.NewMockDeviceRepository())

	device, err := service.CreateDevice(&domain.DeviceRequest{
		Label:     testLabel,
		Algorithm: domain.RSA,
		ID:        "testID",
	})
	if err != nil {
		t.Errorf("Failed to create device: %v", err)
	}
	if device == nil {
		t.Errorf("Device is nil")
	}
	if device.Label != testLabel {
		t.Errorf("Device label is not correct was %s expected %s", device.Label, testLabel)
	}
	if device.Algorithm != domain.RSA {
		t.Errorf("Device algorithm is not correct was %s expected %s", device.Algorithm, domain.RSA)
	}
	if device.Signer == nil {
		t.Errorf("Device signer is nil")
	}
	if device.ID == uuid.Nil {
		t.Errorf("Device ID is nil")
	}
	if device.SignatureCounter != 0 {
		t.Errorf("Device signature counter is not correct was %d expected %d", device.SignatureCounter, 0)
	}
	if device.LastSignature != "" {
		t.Errorf("Device last signature is not correct was %s expected %s", device.LastSignature, "")
	}
}

func TestDeviceService_CreateDevice_TestECDSADevice(t *testing.T) {
	service := NewDeviceService(persistence.NewMockDeviceRepository())

	device, err := service.CreateDevice(&domain.DeviceRequest{
		Label:     testLabel,
		Algorithm: domain.ECDSA,
		ID:        "testID",
	})
	if err != nil {
		t.Errorf("Failed to create device: %v", err)
	}
	if device == nil {
		t.Errorf("Device is nil")
	}
	if device.Label != testLabel {
		t.Errorf("Device label is not correct was %s expected %s", device.Label, testLabel)
	}
	if device.Algorithm != domain.ECDSA {
		t.Errorf("Device algorithm is not correct was %s expected %s", device.Algorithm, domain.ECDSA)
	}
	if device.Signer == nil {
		t.Errorf("Device signer is nil")
	}
	if device.ID == uuid.Nil {
		t.Errorf("Device ID is nil")
	}
	if device.SignatureCounter != 0 {
		t.Errorf("Device signature counter is not correct was %d expected %d", device.SignatureCounter, 0)
	}
	if device.LastSignature != "" {
		t.Errorf("Device last signature is not correct was %s expected %s", device.LastSignature, "")
	}
}

func TestDeviceService_GetDevice(t *testing.T) {
	service := NewDeviceService(persistence.NewMockDeviceRepository())

	device, err := service.GetDevice("testID")
	if err != nil {
		t.Errorf("Failed to get device: %v", err)
	}
	if device.Label != testLabel {
		t.Errorf("Device label is not correct was %s expected %s", device.Label, testLabel)
	}
	if device.Algorithm != domain.RSA {
		t.Errorf("Device algorithm is not correct was %s expected %s", device.Algorithm, domain.RSA)
	}
	if device.SignatureCounter != 100 {
		t.Errorf("Device signature counter is not correct was %d expected %d", device.SignatureCounter, 100)
	}
}

func TestDeviceService_GetDevices(t *testing.T) {
	service := NewDeviceService(persistence.NewMockDeviceRepository())

	devices, err := service.GetDevices()
	if err != nil {
		t.Errorf("Failed to get devices: %v", err)
	}
	if len(devices) != 2 {
		t.Errorf("Device count is not correct was %d expected %d", len(devices), 2)
	}
	if devices[0].Label != testLabel {
		t.Errorf("Device label is not correct was %s expected %s", devices[0].Label, testLabel)
	}
	if devices[0].Algorithm != domain.RSA {
		t.Errorf("Device algorithm is not correct was %s expected %s", devices[0].Algorithm, domain.RSA)
	}
	if devices[0].SignatureCounter != 100 {
		t.Errorf("Device signature counter is not correct was %d expected %d", devices[0].SignatureCounter, 100)
	}
	if devices[1].Label != testLabel2 {
		t.Errorf("Device label is not correct was %s expected %s", devices[1].Label, testLabel2)
	}
	if devices[1].Algorithm != domain.ECDSA {
		t.Errorf("Device algorithm is not correct was %s expected %s", devices[1].Algorithm, domain.ECDSA)
	}
	if devices[1].SignatureCounter != 200 {
		t.Errorf("Device signature counter is not correct was %d expected %d", devices[1].SignatureCounter, 200)
	}
}

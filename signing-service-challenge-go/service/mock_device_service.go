package service

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

var (
	testDevice = &domain.Device{
		ID:               uuid.NameSpaceX500,
		Label:            "testLabel",
		SignatureCounter: 100,
		LastSignature:    "lastSignature",
		Algorithm:        domain.RSA,
	}
	testDevice2 = &domain.Device{
		ID:               uuid.NameSpaceOID,
		Label:            "testLabel2",
		SignatureCounter: 200,
		LastSignature:    "lastSignature2",
		Algorithm:        domain.ECDSA,
	}
)

type mockDeviceService struct{}

func NewMockDeviceService() IDeviceService {
	return &mockDeviceService{}
}

func (s *mockDeviceService) CreateDevice(_ *domain.DeviceRequest) (*domain.Device, error) {
	return testDevice, nil
}

func (s *mockDeviceService) GetDevice(_ string) (*domain.Device, error) {
	return testDevice, nil
}

func (s *mockDeviceService) GetDevices() ([]*domain.Device, error) {
	return []*domain.Device{
		testDevice,
		testDevice2,
	}, nil
}

func (s *mockDeviceService) SignTransaction(_, _ string) (*string, *string, error) {
	signature := "signature"
	signedData := "signedData"

	return &signature, &signedData, nil
}

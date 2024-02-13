package persistence

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type mockDeviceRepository struct{}

func NewMockDeviceRepository() IDeviceRepository {
	return &mockDeviceRepository{}
}

func (m *mockDeviceRepository) CreateDevice(device *domain.Device, _ string) (*domain.Device, error) {
	return device, nil
}

func (m *mockDeviceRepository) GetDevice(_ string) (*domain.Device, error) {
	return &domain.Device{ID: uuid.NameSpaceX500, Label: "testLabel", SignatureCounter: 100, LastSignature: "lastSignature", Algorithm: domain.RSA}, nil
}

func (m *mockDeviceRepository) GetDevices() ([]*domain.Device, error) {
	return []*domain.Device{{ID: uuid.NameSpaceX500, Label: "testLabel", SignatureCounter: 100, LastSignature: "lastSignature", Algorithm: domain.RSA},
		{ID: uuid.NameSpaceOID, Label: "testLabel2", SignatureCounter: 200, LastSignature: "lastSignature2", Algorithm: domain.ECDSA}}, nil
}

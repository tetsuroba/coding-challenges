package persistence

import "github.com/fiskaly/coding-challenges/signing-service-challenge/domain"

type IDeviceRepository interface {
	CreateDevice(device *domain.Device, ID string) (*domain.Device, error)
	GetDevice(ID string) (*domain.Device, error)
	GetDevices() ([]*domain.Device, error)
}

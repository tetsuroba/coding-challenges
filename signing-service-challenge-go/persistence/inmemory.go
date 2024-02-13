package persistence

import (
	"fmt"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"sync"
)

type DeviceRepository struct {
	sync.RWMutex
	devices map[string]*domain.Device
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{
		devices: make(map[string]*domain.Device),
	}
}

func (r *DeviceRepository) CreateDevice(device *domain.Device, ID string) (*domain.Device, error) {
	r.Lock()
	if r.devices[ID] != nil {
		r.Unlock()
		return nil, fmt.Errorf("device with ID %s already exists", ID)
	}
	r.devices[ID] = device
	r.Unlock()
	return device, nil
}

func (r *DeviceRepository) GetDevice(ID string) (*domain.Device, error) {
	r.RLock()
	device, ok := r.devices[ID]
	r.RUnlock()
	if ok {
		return device, nil
	}
	return nil, nil
}

func (r *DeviceRepository) GetDevices() ([]*domain.Device, error) {
	r.RLock()
	devices := make([]*domain.Device, 0, len(r.devices))
	for _, device := range r.devices {
		devices = append(devices, device)
	}
	r.RUnlock()
	return devices, nil
}

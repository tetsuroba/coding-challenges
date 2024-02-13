package persistence

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
	"testing"
)

func TestDeviceRepository_CreateDevice(t *testing.T) {
	repository := NewDeviceRepository()

	device := domain.Device{ID: uuid.NameSpaceX500, Label: "testLabel", SignatureCounter: 100, LastSignature: "lastSignature", Algorithm: domain.RSA}

	createdDevice, err := repository.CreateDevice(&device, "testID")
	if err != nil {
		t.Errorf("Error creating device: %v", err)
	}
	if repository.devices["testID"] != createdDevice {
		t.Errorf("Expected device to be created")
	}
}

func TestDeviceRepository_CreateDevice_Duplicates(t *testing.T) {
	repository := NewDeviceRepository()

	device := domain.Device{ID: uuid.NameSpaceX500, Label: "testLabel", SignatureCounter: 100, LastSignature: "lastSignature", Algorithm: domain.RSA}

	createdDevice, err := repository.CreateDevice(&device, "testID")
	if err != nil {
		t.Errorf("Error creating device: %v", err)
	}
	if repository.devices["testID"] != createdDevice {
		t.Errorf("Expected device to be created")
	}

	_, err = repository.CreateDevice(&device, "testID")
	if err == nil {
		t.Errorf("Expected error creating device")
	}
	if err.Error() != "device with ID testID already exists" {
		t.Errorf("Expected error message to be 'device with ID testID already exists'")
	}
}

func TestDeviceRepository_GetDevice(t *testing.T) {
	repository := NewDeviceRepository()

	device := domain.Device{ID: uuid.NameSpaceX500, Label: "testLabel", SignatureCounter: 100, LastSignature: "lastSignature", Algorithm: domain.RSA}

	createdDevice, err := repository.CreateDevice(&device, "testID")
	if err != nil {
		t.Errorf("Error creating device: %v", err)
	}
	if repository.devices["testID"] != createdDevice {
		t.Errorf("Expected device to be created")
	}

	foundDevice, err := repository.GetDevice("testID")
	if err != nil {
		t.Errorf("Error getting device: %v", err)
	}
	if foundDevice != createdDevice {
		t.Errorf("Expected device to be found")
	}
}

func TestDeviceRepository_GetDevices(t *testing.T) {
	repository := NewDeviceRepository()

	device := domain.Device{ID: uuid.NameSpaceX500, Label: "testLabel", SignatureCounter: 100, LastSignature: "lastSignature", Algorithm: domain.RSA}
	device2 := domain.Device{ID: uuid.NameSpaceDNS, Label: "testLabel2", SignatureCounter: 200, LastSignature: "lastSignature2", Algorithm: domain.ECDSA}

	createdDevice, err := repository.CreateDevice(&device, "testID")
	if err != nil {
		t.Errorf("Error creating device: %v", err)
	}
	if repository.devices["testID"] != createdDevice {
		t.Errorf("Expected device to be created")
	}

	createdDevice2, err := repository.CreateDevice(&device2, "testID2")
	if err != nil {
		t.Errorf("Error creating device: %v", err)
	}
	if repository.devices["testID2"] != createdDevice2 {
		t.Errorf("Expected device to be created")
	}

	devices, err := repository.GetDevices()
	if err != nil {
		t.Errorf("Error getting devices: %v", err)
	}
	if len(devices) != 2 {
		t.Errorf("Expected 2 device to be found")
	}
}

package service

import (
	"fmt"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"log/slog"
	"os"
)

type IDeviceService interface {
	CreateDevice(request *domain.DeviceRequest) (*domain.Device, error)
	GetDevice(ID string) (*domain.Device, error)
	GetDevices() ([]*domain.Device, error)
	SignTransaction(deviceID, dataToBeSigned string) (*string, *string, error)
}

type DeviceService struct {
	DeviceRepository persistence.IDeviceRepository
	logger           *slog.Logger
	ECCGenerator     *crypto.ECCGenerator
	RSAGenerator     *crypto.RSAGenerator
}

func NewDeviceService(repository persistence.IDeviceRepository) *DeviceService {
	return &DeviceService{
		DeviceRepository: repository,
		logger:           slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		ECCGenerator:     &crypto.ECCGenerator{},
		RSAGenerator:     &crypto.RSAGenerator{},
	}
}

func (ds *DeviceService) CreateDevice(request *domain.DeviceRequest) (*domain.Device, error) {
	var signer crypto.Signer
	switch request.Algorithm {
	case domain.RSA:
		rsaKeyPair, err := ds.RSAGenerator.Generate()
		if err != nil {
			ds.logger.Error("Failed to generate RSA key pair", "error", err)
			return nil, err
		}
		signer, err = crypto.NewRSASigner(rsaKeyPair)
		if err != nil {
			ds.logger.Error("Failed to create RSA signer", "error", err)
			return nil, err
		}
	case domain.ECDSA:
		eccKeyPair, err := ds.ECCGenerator.Generate()
		if err != nil {
			ds.logger.Error("Failed to generate ECDSA key pair", "error", err)
			return nil, err
		}
		signer, err = crypto.NewECCSigner(eccKeyPair)
		if err != nil {
			ds.logger.Error("Failed to create ECC signer", "error", err)
			return nil, err
		}
	}
	device, err := domain.NewDevice(request.Label, request.Algorithm, signer)
	if err != nil {
		ds.logger.Error("Failed to create device", "error", err)
		return nil, err
	}

	createdDevice, err := ds.DeviceRepository.CreateDevice(device, request.ID)
	if err != nil {
		ds.logger.Error("Failed to create device", "error", err)
		return nil, err
	}
	return createdDevice, nil
}

func (ds *DeviceService) GetDevice(ID string) (*domain.Device, error) {
	device, err := ds.DeviceRepository.GetDevice(ID)
	if err != nil {
		ds.logger.Error("Failed to get device", "error", err)
		return nil, err
	}
	if device == nil {
		ds.logger.Error("Device not found", "ID", ID)
		return nil, fmt.Errorf("device not found")
	}
	return device, nil
}

func (ds *DeviceService) GetDevices() ([]*domain.Device, error) {
	devices, err := ds.DeviceRepository.GetDevices()
	if err != nil {
		ds.logger.Error("Failed to get devices", "error", err)
		return nil, err
	}
	return devices, nil
}

func (ds *DeviceService) SignTransaction(deviceID, dataToBeSigned string) (*string, *string, error) {
	device, err := ds.DeviceRepository.GetDevice(deviceID)
	if err != nil {
		ds.logger.Error("Failed to get device", "error", err)
		return nil, nil, err
	}
	if device == nil {
		ds.logger.Error("Device not found", "deviceID", deviceID)
		return nil, nil, fmt.Errorf("device not found")
	}
	signature, signedData, err := device.SignTransaction(deviceID, dataToBeSigned)
	if err != nil {
		ds.logger.Error("Failed to sign transaction", "error", err)
		return nil, nil, err
	}
	return signature, signedData, nil
}

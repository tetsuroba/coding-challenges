package api

import (
	"encoding/json"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"net/http"
)

type DeviceResponse struct {
	Device *domain.Device `json:"device" bson:"Device"`
}

type DevicesResponse struct {
	Devices []*domain.Device `json:"devices" bson:"Devices"`
}

type SignatureResponse struct {
	Signature  string `json:"signature" bson:"Signature"`
	SignedData string `json:"signed_data" bson:"SignedData"`
}

// CreateDevice creates a new device and returns the created device
func (s *Server) CreateDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		s.logger.Error("Method not allowed on the endpoint /device", "method", request.Method)
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	var deviceRequest domain.DeviceRequest
	err := json.NewDecoder(request.Body).Decode(&deviceRequest)
	if err != nil {
		s.logger.Error("Failed to decode request", "error", err)
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
		})
		return
	}
	if deviceRequest.ID == "" {
		s.logger.Error("ID is missing")
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"ID is missing",
		})
		return
	}
	if deviceRequest.Algorithm == "" {
		s.logger.Error("Algorithm is missing")
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Algorithm is missing",
		})
		return
	}
	device, err := s.deviceService.CreateDevice(&deviceRequest)
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			err.Error(),
		})
		return
	}

	deviceResponse := DeviceResponse{
		Device: device,
	}

	WriteAPIResponse(response, http.StatusCreated, deviceResponse)
}

// GetDevice retrieves a device with the given ID
func (s *Server) GetDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		s.logger.Error("Method not allowed on the endpoint /device/", "method", request.Method)
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	ID := request.URL.Query().Get("id")
	if ID == "" {
		s.logger.Error("ID is missing")
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"ID is missing",
		})
		return
	}
	device, err := s.deviceService.GetDevice(ID)
	if err != nil {
		WriteInternalError(response)
		return
	}

	if device == nil {
		WriteErrorResponse(response, http.StatusNotFound, []string{
			http.StatusText(http.StatusNotFound),
		})
		return
	}

	deviceResponse := DeviceResponse{
		Device: device,
	}
	WriteAPIResponse(response, http.StatusOK, deviceResponse)
}

// GetDevices retrieves all devices
func (s *Server) GetDevices(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		s.logger.Error("Method not allowed on the endpoint /devices", "method", request.Method)
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	devices, err := s.deviceService.GetDevices()
	if err != nil {
		WriteInternalError(response)
		return
	}

	deviceResponse := DevicesResponse{
		Devices: devices,
	}
	WriteAPIResponse(response, http.StatusOK, deviceResponse)
}

// SignTransaction signs the given data with the device with the given ID
func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		s.logger.Error("Method not allowed on the endpoint /sign", "method", request.Method)
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	var signRequest domain.SignRequest
	err := json.NewDecoder(request.Body).Decode(&signRequest)
	if err != nil {
		s.logger.Error("Failed to decode request", "error", err)
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if signRequest.DeviceId == "" {
		s.logger.Error("DeviceId is missing")
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"DeviceId is missing",
		})
		return
	}

	if signRequest.Data == "" {
		s.logger.Error("Data is missing")
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Data is missing",
		})
		return
	}

	signature, signedData, err := s.deviceService.SignTransaction(signRequest.DeviceId, signRequest.Data)
	if err != nil {
		s.logger.Error("Failed to sign transaction", "error", err)
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			err.Error(),
		})
		return
	}

	signatureResponse := SignatureResponse{
		Signature:  *signature,
		SignedData: *signedData,
	}
	WriteAPIResponse(response, http.StatusOK, signatureResponse)
}

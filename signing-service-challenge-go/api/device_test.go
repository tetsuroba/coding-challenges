package api

import (
	"bytes"
	"encoding/json"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/service"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server = &Server{
		deviceService: service.NewMockDeviceService(),
		listenAddress: ":8080",
		logger:        slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
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

func TestServer_CreateDevice(t *testing.T) {
	t.Run("returns created device", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/device", bytes.NewBuffer([]byte(`{"id":"testID","label":"testLabel","algorithm":"RSA"}`)))

		response := httptest.NewRecorder()
		server.CreateDevice(response, request)

		got := response.Body.String()
		deviceResponse := DeviceResponse{
			Device: testDevice,
		}
		httpResp := Response{
			Data: deviceResponse,
		}
		marshaledExpectedResponse, err := json.MarshalIndent(httpResp, "", "  ")
		if err != nil {
			t.Errorf("failed to marshal response: %v", err)
		}
		want := string(marshaledExpectedResponse)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if method is not POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/device", nil)

		response := httptest.NewRecorder()
		server.CreateDevice(response, request)

		got := response.Body.String()
		want := `{"errors":["Method Not Allowed"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if ID is missing", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/device", bytes.NewBuffer([]byte(`{}`)))

		response := httptest.NewRecorder()
		server.CreateDevice(response, request)

		got := response.Body.String()
		want := `{"errors":["ID is missing"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if Algorithm is missing", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/device", bytes.NewBuffer([]byte(`{"id":"testID"}`)))

		response := httptest.NewRecorder()
		server.CreateDevice(response, request)

		got := response.Body.String()
		want := `{"errors":["Algorithm is missing"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestServer_GetDevice(t *testing.T) {
	t.Run("returns device with the given ID", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/device?id=testID", nil)

		response := httptest.NewRecorder()
		server.GetDevice(response, request)

		got := response.Body.String()
		deviceResponse := DeviceResponse{
			Device: testDevice,
		}
		httpResp := Response{
			Data: deviceResponse,
		}
		marshaledExpectedResponse, err := json.MarshalIndent(httpResp, "", "  ")
		if err != nil {
			t.Errorf("failed to marshal response: %v", err)
		}
		want := string(marshaledExpectedResponse)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if method is not GET", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/device", nil)

		response := httptest.NewRecorder()
		server.GetDevice(response, request)

		got := response.Body.String()
		want := `{"errors":["Method Not Allowed"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if ID is missing", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/device", nil)

		response := httptest.NewRecorder()
		server.GetDevice(response, request)

		got := response.Body.String()
		want := `{"errors":["ID is missing"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestServer_GetDevices(t *testing.T) {
	t.Run("returns all devices", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/devices", nil)

		response := httptest.NewRecorder()
		server.GetDevices(response, request)

		got := response.Body.String()
		deviceResponse := DevicesResponse{
			Devices: []*domain.Device{testDevice, testDevice2},
		}
		httpResp := Response{
			Data: deviceResponse,
		}
		marshaledExpectedResponse, err := json.MarshalIndent(httpResp, "", "  ")
		if err != nil {
			t.Errorf("failed to marshal response: %v", err)
		}
		want := string(marshaledExpectedResponse)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if method is not GET", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/devices", nil)

		response := httptest.NewRecorder()
		server.GetDevices(response, request)

		got := response.Body.String()
		want := `{"errors":["Method Not Allowed"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestServer_SignTransaction(t *testing.T) {
	t.Run("returns signature and data to be signed", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/sign", bytes.NewBuffer([]byte(`{"deviceId":"testID","data":"testData"}`)))

		response := httptest.NewRecorder()
		server.SignTransaction(response, request)

		got := response.Body.String()
		signResponse := SignatureResponse{
			Signature:  "signature",
			SignedData: "signedData",
		}
		httpResp := Response{
			Data: signResponse,
		}
		marshaledExpectedResponse, err := json.MarshalIndent(httpResp, "", "  ")
		if err != nil {
			t.Errorf("failed to marshal response: %v", err)
		}
		want := string(marshaledExpectedResponse)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if method is not POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/sign", nil)

		response := httptest.NewRecorder()
		server.SignTransaction(response, request)

		got := response.Body.String()
		want := `{"errors":["Method Not Allowed"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if DeviceId is missing", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/sign", bytes.NewBuffer([]byte(`{"data":"testData"}`)))

		response := httptest.NewRecorder()
		server.SignTransaction(response, request)

		got := response.Body.String()
		want := `{"errors":["DeviceId is missing"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns error if Data is missing", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/sign", bytes.NewBuffer([]byte(`{"deviceId":"testID"}`)))

		response := httptest.NewRecorder()
		server.SignTransaction(response, request)

		got := response.Body.String()
		want := `{"errors":["Data is missing"]}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

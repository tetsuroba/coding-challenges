package api

import (
	"encoding/json"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/service"
	"log/slog"
	"net/http"
	"os"
)

// Response is the generic API response container.
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is the generic error API response container.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Server manages HTTP requests and dispatches them to the appropriate services.
type Server struct {
	listenAddress string
	deviceService service.IDeviceService
	logger        *slog.Logger
}

// NewServer is a factory to instantiate a new Server.
func NewServer(listenAddress string) *Server {
	deviceRepository := persistence.NewDeviceRepository()
	return &Server{
		listenAddress: listenAddress,
		deviceService: service.NewDeviceService(deviceRepository),
		logger:        slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

// Run registers all HandlerFuncs for the existing HTTP routes and starts the Server.
func (s *Server) Run() error {
	mux := http.NewServeMux()

	mux.Handle("/api/v0/health", http.HandlerFunc(s.Health))
	mux.Handle("/api/v0/device", http.HandlerFunc(s.CreateDevice))
	mux.Handle("/api/v0/device/", http.HandlerFunc(s.GetDevice))
	mux.Handle("/api/v0/devices", http.HandlerFunc(s.GetDevices))
	mux.Handle("/api/v0/sign", http.HandlerFunc(s.SignTransaction))

	return http.ListenAndServe(s.listenAddress, mux)
}

// WriteInternalError writes a default internal error message as an HTTP response.
func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

// WriteErrorResponse takes an HTTP status code and a slice of errors
// and writes those as an HTTP error response in a structured format.
func WriteErrorResponse(w http.ResponseWriter, code int, errors []string) {
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Errors: errors,
	}

	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}

// WriteAPIResponse takes an HTTP status code and a generic data struct
// and writes those as an HTTP response in a structured format.
func WriteAPIResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)

	response := Response{
		Data: data,
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}

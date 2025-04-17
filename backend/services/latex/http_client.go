package latex

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/huuloc2026/generateJD-microservices.git/config"
	"github.com/huuloc2026/generateJD-microservices.git/model"
)

type HTTPClient struct {
	PythonEndpoint string
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		PythonEndpoint: config.AppConfig.PythonHost,
	}
}

func (h *HTTPClient) Generate(req model.JDRequest) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(h.PythonEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

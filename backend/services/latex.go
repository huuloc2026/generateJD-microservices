package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/huuloc2026/generateJD-microservices.git/config"
	"github.com/huuloc2026/generateJD-microservices.git/model"
)

func GeneratePDF(data model.JDRequest) ([]byte, error) {
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(config.AppConfig.PythonHost, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

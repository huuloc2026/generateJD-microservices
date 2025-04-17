package latex

import "github.com/huuloc2026/generateJD-microservices.git/model"

type Generator interface {
	Generate(req model.JDRequest) ([]byte, error)
}

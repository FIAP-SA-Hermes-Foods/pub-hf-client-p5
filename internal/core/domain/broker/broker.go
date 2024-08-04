package broker

import (
	"pub-hf-client-p5/internal/core/domain/entity/dto"
)

type ClientBroker interface {
	GetClientByCPF(input dto.ClientBroker) error
	SaveClient(input dto.ClientBroker) error
}

package useCase

import "pub-hf-client-p5/internal/core/domain/entity/dto"

type ClientUseCase interface {
	SaveClient(reqClient dto.RequestClient) error
	GetClientByCPF(uuid string) error
	
}

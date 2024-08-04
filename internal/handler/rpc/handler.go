package rpc

import (
	"context"
	cp "pub-hf-client-p5/client_pub_proto"
	l "pub-hf-client-p5/external/logger"
	"pub-hf-client-p5/internal/core/application"
	"pub-hf-client-p5/internal/core/domain/entity/dto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedClientServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) GetClientByCPF(ctx context.Context, req *cp.GetClientByCPFRequest) (*cp.GetClientByCPFResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	if err := h.app.GetClientByCPF(msgID, req.Cpf); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) CreateClient(ctx context.Context, req *cp.CreateClientRequest) (*cp.CreateClientResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestClient{
		Name:        req.Name,
		CPF:         req.Cpf,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}

	if err := h.app.SaveClient(msgID, input); err != nil {
		return nil, err
	}

	return nil, nil
}

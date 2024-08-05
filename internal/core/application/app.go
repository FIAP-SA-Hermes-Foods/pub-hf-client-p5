package application

import (
	"context"
	l "pub-hf-client-p5/external/logger"
	ps "pub-hf-client-p5/external/strings"
	"pub-hf-client-p5/internal/core/domain/broker"
	"pub-hf-client-p5/internal/core/domain/entity/dto"
)

type Application interface {
	GetClientByCPF(msgID string, uuid string) error
	SaveClient(msgID string, client dto.RequestClient) error
}

type application struct {
	ctx          context.Context
	clientBroker broker.ClientBroker
}

func NewApplication(ctx context.Context, clientBroker broker.ClientBroker) Application {
	return application{
		ctx:          ctx,
		clientBroker: clientBroker,
	}
}

func (app application) GetClientByCPF(msgID string, uuid string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetClientByCPFApp: ", " | ", uuid)

	inputBroker := dto.ClientBroker{
		UUID:      uuid,
		MessageID: msgID,
	}

	if err := app.clientBroker.GetClientByCPF(inputBroker); err != nil {
		l.Errorf(msgID, "GetClientByCPFApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "GetClientByCPFApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) SaveClient(msgID string, client dto.RequestClient) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "SaveClientApp: ", " | ", ps.MarshalString(client))

	inputBroker := dto.ClientBroker{
		UUID:        client.UUID,
		MessageID:   msgID,
		Name:        client.Name,
		CPF:         client.CPF,
		Email:       client.Email,
		PhoneNumber: client.PhoneNumber,
		Address:     client.Address,
		CreatedAt:   client.CreatedAt,
	}

	if err := app.clientBroker.SaveClient(inputBroker); err != nil {
		l.Errorf(msgID, "SaveClientApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "SaveClientApp output: ", " | ", "message sent with success!")
	return nil
}

type contextKey string

func (app *application) setMessageIDCtx(msgID string) {
	if app.ctx == nil {
		app.ctx = context.WithValue(context.Background(), contextKey("messageID"), msgID)
		return
	}
	app.ctx = context.WithValue(app.ctx, contextKey("messageID"), msgID)
}

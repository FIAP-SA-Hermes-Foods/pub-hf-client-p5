package broker

import (
	l "pub-hf-client-p5/external/logger"
	ps "pub-hf-client-p5/external/strings"
	sqsBroker "pub-hf-client-p5/internal/core/broker"
	cBroker "pub-hf-client-p5/internal/core/domain/broker"
	"pub-hf-client-p5/internal/core/domain/entity/dto"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var _ cBroker.ClientBroker = (*clientBroker)(nil)

type clientBroker struct {
	queueURL string
	broker   sqsBroker.SQSBroker
}

func NewClientBroker(broker sqsBroker.SQSBroker, queueURL string) *clientBroker {
	return &clientBroker{broker: broker, queueURL: queueURL}
}

func (p *clientBroker) GetClientByCPF(input dto.ClientBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *clientBroker) SaveClient(input dto.ClientBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

package streaming

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"nats"
	"nats/pkg/service"
)

type Stream struct {
	s  *service.Service
	SC stan.Conn
}

func NewStream(s *service.Service) *Stream {
	return &Stream{s: s}
}

const (
	clusterID = "test-cluster"
	clientID  = "order-subscriber"
	channel   = "order-notification"
)

func (str *Stream) handleOrder(orderMSG *stan.Msg) {
	order := nats.Order{}

	err := json.Unmarshal(orderMSG.Data, &order)
	if err != nil {
		logrus.Errorf("Cannot unmarshal data from nats-streaming-server")
		return
	}

	err = str.s.Create(order)
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
}

func (str *Stream) natsStreamingSubscribe() {
	//_, err := str.SC.Subscribe(channel, str.handleOrder, stan.StartWithLastReceived())
	_, err := str.SC.Subscribe(channel, str.handleOrder)
	if err != nil {
		logrus.Fatalf("Cannot subscribe to nats-streaming-server chanel")
	}
}

func (str *Stream) NatsStreamingSetup() error {
	var err error
	str.SC, err = stan.Connect(clusterID, clientID)
	if err != nil {
		return err
	}
	str.natsStreamingSubscribe()
	return nil
}

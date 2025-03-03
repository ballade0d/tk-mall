package server

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	pb "mall/api/mall/service/v1"
	"mall/app/payment/internal/config"
	"mall/app/payment/internal/service"
	"mall/ent"
)

func NewRabbitMQServer(conf *config.Config, paymentService *service.PaymentService) {
	conn, err := amqp.Dial(conf.RabbitMQ.Addr)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	orderQueue, err := ch.Consume("order_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for msg := range orderQueue {
			var o ent.Order
			err := json.Unmarshal(msg.Body, &o)
			if err != nil {
				log.Println(err)
				continue
			}
			_, err = paymentService.PayOrder(context.Background(), &pb.PayOrderRequest{OrderId: int64(o.ID)})
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}()
}

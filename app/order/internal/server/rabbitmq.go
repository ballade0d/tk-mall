package server

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"mall/app/order/internal/config"
	"mall/app/order/internal/service"
	"mall/ent"
	"mall/ent/order"
	"mall/ent/payment"
)

func NewRabbitMQServer(conf *config.Config, orderService *service.OrderService) {
	conn, err := amqp.Dial(conf.RabbitMQ.Addr)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := ch.Consume("pay_result_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for msg := range msgs {
			var p ent.Payment
			err := json.Unmarshal(msg.Body, &p)
			if err != nil {
				log.Println(err)
				continue
			}
			if p.Status == payment.StatusPaid {
				err = orderService.UpdateOrderStatus(context.Background(), p.Edges.Order.ID, order.StatusPaid)
			}
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}()
	orderDeadLetterQueue, err := ch.Consume("order_dead_letter_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for msg := range orderDeadLetterQueue {
			var o ent.Order
			err := json.Unmarshal(msg.Body, &o)
			if err != nil {
				log.Println(err)
				continue
			}
			err = orderService.UpdateOrderStatus(context.Background(), o.ID, order.StatusCancelled)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}()
}

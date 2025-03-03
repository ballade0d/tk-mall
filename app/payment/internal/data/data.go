package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"log"
	"mall/app/payment/internal/config"
	"mall/ent"
)

var ProviderSet = wire.NewSet(NewData, NewPaymentRepo)

type Data struct {
	conf *config.Config
	db   *ent.Client
	rdb  *redis.Client
	es   *elasticsearch.TypedClient
	mq   *amqp.Channel
}

func NewData(conf *config.Config) (*Data, error) {
	// Open the database connection
	drv, err := sql.Open(
		conf.Database.Driver, conf.Database.Source,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db := ent.NewClient(ent.Driver(drv))
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Open the redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.Database,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Open the elasticsearch connection
	cfg := elasticsearch.Config{
		Addresses: conf.ElasticSearch.Addresses,
		APIKey:    conf.ElasticSearch.APIKey,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	exists, err := es.Indices.Exists(conf.ElasticSearch.Indices).Do(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if !exists {
		_, err := es.Indices.Create(conf.ElasticSearch.Indices).
			Request(&create.Request{
				Mappings: &types.TypeMapping{
					Properties: map[string]types.Property{
						"name":        types.NewTextProperty(),
						"description": types.NewTextProperty(),
						"price":       types.NewFloatNumberProperty(),
					},
				},
			}).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}

	// Open the rabbitmq connection
	mq, err := amqp.Dial(conf.RabbitMQ.Addr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ch, err := mq.Channel()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Declare the dead letter queue
	dlxQueue := "order_dead_letter_queue"
	_, err = ch.QueueDeclare(dlxQueue, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Declare the order queue
	queueName := "order_queue"
	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "",
			"x-dead-letter-routing-key": dlxQueue,
			"x-message-ttl":             900000, // 15分钟超时（毫秒）
		},
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Declare the pay result queue
	payResultQueue := "pay_result_queue"
	_, err = ch.QueueDeclare(payResultQueue, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &Data{
		conf: conf,
		db:   db,
		rdb:  rdb,
		es:   es,
		mq:   ch,
	}, nil
}

package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"log"
	"mall/app/auth/internal/config"
	"mall/ent"
)

var ProviderSet = wire.NewSet(NewData, NewPasswordRepo, NewUserRepo)

type Data struct {
	conf *config.Config
	db   *ent.Client
	es   *elasticsearch.TypedClient
}

func NewData(conf *config.Config) (*Data, error) {
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
	return &Data{
		conf: conf,
		db:   db,
		es:   es,
	}, nil
}

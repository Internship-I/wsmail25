package database

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewMySQLConnection(dbConfig string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB.Statement.RaiseErrorOnNotFound = true

	if err != nil {
		panic(err)
	}
	return DB
}

func NewMongoDBConnection(uri string, dbName string) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	return client, client.Database(dbName), nil
}

type NullableInt struct {
	Valid bool
	Value int
}

func (n *NullableInt) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		return nil
	}
	var temp int
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	n.Value = temp
	n.Valid = true
	return nil
}

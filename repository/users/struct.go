package users

import "go.mongodb.org/mongo-driver/mongo"

func NewPenggunaTable(db *mongo.Client) *MUsers {
	return &MUsers{
		db: db.Database("intership1"),
	}
}

type MUsers struct {
	db *mongo.Database
}

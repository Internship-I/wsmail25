package trans

import "go.mongodb.org/mongo-driver/mongo"

func NewTransaksiTable(db *mongo.Client) *MTrans {
	return &MTrans{
		db: db.Database("Internship1"),
	}
}

type MTrans struct {
	db *mongo.Database
}
package trans

import "go.mongodb.org/mongo-driver/mongo"

func NewTransaksiTable(db *mongo.Client) *MTransaction {
	return &MTransaction{
		db: db.Database("intership1"),
	}
}

type MTransaction struct {
	db *mongo.Database
}

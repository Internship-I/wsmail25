package repository

import (
	"context"
	"wsmail25/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context) ([]model.Users, error)
	InsertUser	(ctx context.Context, user model.Users) (model.Users, error)
	GetAllTransaction(ctx context.Context) ([]model.Transaction, error)
	GetByDeliveryStatus(ctx context.Context, status string) ([]model.Transaction, error)
	InsertTransaction(ctx context.Context, trans model.Transaction) (model.Transaction, error)
	GetByConnote(ctx context.Context, connote string) (model.Transaction, error)
	UpdateDeliveryStatus(ctx context.Context, id primitive.ObjectID, status string) (model.Transaction, error)	
	SendWAOnDelivery(ctx context.Context, transactionID primitive.ObjectID) (model.Transaction, error)
	SendWADelivered(ctx context.Context, transactionID primitive.ObjectID) (model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string) (model.Transaction, error)
}

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context) ([]model.Transaction, error)
	InsertTransaction(ctx context.Context, trans model.Transaction) (model.Transaction, error)
	GetByConnote(ctx context.Context, connote string) (model.Transaction, error)
	GetByDeliveryStatus(ctx context.Context, status string) ([]model.Transaction, error)
	// UpdateDeliveryStatus(ctx context.Context, connote string, status string) (model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string) (model.Transaction, error)
}

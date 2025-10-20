package repository

import (
	"context"
	"wsmail25/model"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context) ([]model.Users, error)
	InsertUser	(ctx context.Context, user model.Users) (model.Users, error)
	GetAllTransaction(ctx context.Context) ([]model.Transaction, error)
	InsertTransaction(ctx context.Context, trans model.Transaction) (model.Transaction, error)
	
}

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context) ([]model.Transaction, error)
	InsertTransaction(ctx context.Context, trans model.Transaction) (model.Transaction, error)
}

package repository

import "wsmail25/model"

type UsersRepository interface {
	InsertUser(users model.User) (err error)
	GetAllUsers() (users model.User, err error)
}

type TransactionRepository interface{
	InsertTransaction(transaction model.Transaction) (err error)
	GetAllTransaction() (transaction model.Transaction, err error)
}
package repository

import "wsmail25/model"

type UsersRepository interface {
	InsertUser(users model.Users) (err error)
	GetAllUsers() (users model.Users, err error)
}

type TransactionRepository interface{
	InsertTransaction(transaction model.Transaction) (err error)
	GetAllTransaction() (transaction model.Transaction, err error)
}
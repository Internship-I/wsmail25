package user

import (
	"go.mongodb.org/mongo-driver/mongo"
	"wsmail25/model"
)

func NewUsersTable(db *mongo.Client) *MUsers {
	return &MUsers{
		db: db.Database("internship1"),
	}
}

type MUsers struct {
	db *mongo.Database
}

// GetAllUsers implements repository.UsersRepository.
func (r *MUsers) GetAllUsers() (users model.User, err error) {
	panic("unimplemented")
}

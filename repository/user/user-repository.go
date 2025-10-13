package user

import (
	"context"
	"wsmail25/model"
)

func (r *MUsers) InsertUser(users model.User) (err error) {
	collection := r.db.Collection("users")
	_, err = collection.InsertOne(nil, users)
	if err != nil {
		return err
	}
	return nil
}

func (r *MUsers) GetAllUser() (users model.User, err error) {
	collection := r.db.Collection("users")
	_, err = collection.Find(context.Background(), &users)
	if err != nil {
		return users, err
	}
	return users, nil
}
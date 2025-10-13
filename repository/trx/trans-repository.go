package trans

import (
	"context"
	"wsmail25/model"
)

func (r *MTransaction) InsertTransaction(trans model.Transaction) (err error){
	collection := r.db.Collection("MailApp")
	_, err = collection.InsertOne(nil, trans)
	if err != nil {
		return err
	}
	return nil
}

func (r *MTransaction) GetAllTransaction() (trans model.Transaction, err error){
	collection := r.db.Collection("MailApp")
	_, err = collection.Find(context.Background(), &trans)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

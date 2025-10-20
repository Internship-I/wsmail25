package users

import (
	"context"
	"fmt"
	"log"
	"time"
	"wsmail25/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *MUsers) GetAllUsers(ctx context.Context) ([]model.Users, error){
	var users []model.Users
	collection := u.db.Collection("User")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.Users
		if err := cursor.Decode(&user); err != nil {
			log.Println("failed to decode user:", err)
			continue
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengambil data user: %w", err)
	}

	return users, nil
}

func (u *MUsers) InsertUser	(ctx context.Context, user model.Users) (model.Users, error){
	collection := u.db.Collection("User")
	userData := bson.M{
		"fullname": user.FullName,
		"phone_number": user.PhoneNumber,
		"username": user.Username,
		"password": user.Password,
		"role":		user.Role,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	result, err := collection.InsertOne(ctx, userData)
	if err != nil {
		log.Println("Gagal menyimpan data ke database", err)
		return model.Users{}, fmt.Errorf("gagal menyimpan data user: %w", err)
	}

	insertedID :=fmt.Sprintf("%v", result.InsertedID)
	user.ID = insertedID

	log.Println("Data user berhasil disimpan dengan ID:", insertedID)
	return user, nil
}

func (u *MUsers) GetUserByID(ctx context.Context, userID string) (model.Users, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return model.Users{}, fmt.Errorf("invalid user ID: %w", err)
	}
	var user model.Users
	collection := u.db.Collection("User")
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

func (u *MUsers) DeleteUser(ctx context.Context, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	filter := bson.M{"_id": objectID}
	log.Println("[INFO] Menghapus data user dengan filter:", filter)

	var deletedUser model.Users
	collection := u.db.Collection("User")
	err = collection.FindOneAndDelete(ctx, filter).Decode(&deletedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user dengan id %s tidak ditemukan", userID)
		}
		return fmt.Errorf("failed to delete user: %w", err)
	}

	log.Println("[INFO] User berhasil dihapus dengan ID:", userID)
	return nil
} 
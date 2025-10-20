package trans

import (
	"context"
	"fmt"
	"log"
	"time"
	"wsmail25/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MTrans) GetAllTransaction(ctx context.Context) ([]model.Transaction, error) {
	var trans []model.Transaction
	collection := r.db.Collection("MailApp")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var transaction model.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %v", err)
			continue
		}
		trans = append(trans, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("gagal mengambil data transaksi: %w", err)
	}

	return trans, nil
}

//GetByConnote
func (r *MTrans) GetByConnote(ctx context.Context, connote string) (model.Transaction, error) {
	var tr model.Transaction
	collection := r.db.Collection("MailApp")

	log.Println("[DEBUG] Searching transaction with connote:", connote)

	// Handle kemungkinan perbedaan nama field di MongoDB
	filter := bson.M{
		"$or": []bson.M{
			{"consignment_note": connote},
			{"ConsigmentNote": connote}, // fallback jika field lama masih pakai typo
		},
	}

	err := collection.FindOne(ctx, filter).Decode(&tr)
	if err != nil {
		log.Println("[ERROR] Transaction not found for connote:", connote, "| Error:", err)
		return model.Transaction{}, fmt.Errorf("transaction not found: %w", err)
	}

	log.Println("[INFO] Transaction found:", tr.ConsigmentNote)
	return tr, nil
}

//Generate Consignment Note
func (r *MTrans) GenerateConnote() string {
	// Format tanggal: ddMMyy (contoh: 201025)
	datePart := time.Now().Format("020106")

	// Buat 6 digit angka acak
	randomPart := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)

	// Gabungkan jadi format akhir: P2507210080398
	return fmt.Sprintf("P%s%s", datePart, randomPart)
}

func (r *MTrans) InsertTransaction(ctx context.Context, trans model.Transaction) (model.Transaction, error){
	collection := r.db.Collection("MailApp")

	// Generate Connote otomatis jika kosong
	if trans.ConsigmentNote == "" {
		trans.ConsigmentNote = r.GenerateConnote()
	}
	now := time.Now()

	transData := bson.M{
		"consignment_note": trans.ConsigmentNote,
		"sender_name":      trans.SenderName,
		"sender_phone":     trans.SenderPhone,
		"receiver_name":    trans.ReceiverName,
		"address_receiver": trans.AddressReceiver,
		"receiver_phone":   trans.ReceiverPhone,
		"item_content":     trans.ItemContent,
		"service_type":     trans.ServiceType,
		"cod_value":        trans.CODValue,
		"wa_sent":          trans.WASent,
		"wa_sent_at":       trans.WASentAt,
		"created_at":       now,
		"updated_at":       now,
	}

	result, err := collection.InsertOne(ctx, transData)
	if err != nil {
		log.Println("[ERROR] Failed to save transaction data:", err)
		return model.Transaction{}, fmt.Errorf("failed to save transaction data: %w", err)
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		trans.ID = insertedID
	}
	trans.CreatedAt = now
	trans.UpdatedAt = now

	return trans, nil
}

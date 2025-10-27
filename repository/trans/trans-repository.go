package trans

import (
	"context"
	"fmt"
	"log"
	"time"
	"errors"
	"wsmail25/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	log.Println("[INFO] Transaction found:", tr.ConsignmentNote)
	return tr, nil
}
// GetByDeliveryStatus
func (r *MTrans) GetByDeliveryStatus(ctx context.Context, status string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	collection := r.db.Collection("MailApp")
	filter := bson.M{"delivery_status": status}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions by status: %v", err)	
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var transaction model.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through transactions: %w", err)
	}
	return transactions, nil
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

//InsertTransaction
func (r *MTrans) InsertTransaction(ctx context.Context, trans model.Transaction) (model.Transaction, error){
	collection := r.db.Collection("MailApp")

	// Generate Connote otomatis jika kosong
	if trans.ConsignmentNote == "" {
		trans.ConsignmentNote = r.GenerateConnote()
	}
	now := time.Now()

	transData := bson.M{
		"consignment_note": trans.ConsignmentNote,
		"sender_name":      trans.SenderName,
		"sender_phone":     trans.SenderPhone,
		"receiver_name":    trans.ReceiverName,
		"address_receiver": trans.AddressReceiver,
		"receiver_phone":   trans.ReceiverPhone,
		"item_content":     trans.ItemContent,
		"service_type":     trans.ServiceType,
		"cod_value":        trans.CODValue,
		// Assignment (optional saat create)
		"courier_id":         trans.CourierID,
		"courier_name":       trans.CourierName,
		// WhatsApp tracking - default false
		"wa_on_delivery_sent":    false, // PERBAIKAN: field baru
		"wa_on_delivery_sent_at": nil,
		"wa_delivered_sent":      false, // PERBAIKAN: field baru
		"wa_delivered_sent_at":   nil,
		// Delivery status
		"delivery_status": trans.DeliveryStatus,
		// Timestamps
		"created_at": now,
		"updated_at": now,
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

// UpdateDeliveryStatus by ObjectID
func (r *MTrans) UpdateDeliveryStatus(ctx context.Context, id primitive.ObjectID, status string) (model.Transaction, error) {
	col := r.db.Collection("MailApp")

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"delivery_status": status,
			"updated_at":      time.Now(),
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated model.Transaction
	if err := col.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated); err != nil {
		return model.Transaction{}, fmt.Errorf("gagal memperbarui status pengiriman: %w", err)
	}
	return updated, nil
}

// SendWAOnDelivery
func (r *MTrans) SendWAOnDelivery(ctx context.Context, transactionID primitive.ObjectID) (model.Transaction, error) {
	col := r.db.Collection("MailApp")

	var trx model.Transaction
	if err := col.FindOne(ctx, bson.M{"_id": transactionID}).Decode(&trx); err != nil {
		return model.Transaction{}, fmt.Errorf("transaksi tidak ditemukan: %w", err)
	}

	if trx.WAOnDeliverySent {
		return model.Transaction{}, fmt.Errorf("pesan WA 'on delivery' sudah pernah dikirim")
	}

	now := time.Now()
	update := bson.M{
		"$set": bson.M{
			"wa_on_delivery_sent":    true,
			"wa_on_delivery_sent_at": now,
			"delivery_status":        "on_delivery",
			"updated_at":             now,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated model.Transaction
	if err := col.FindOneAndUpdate(ctx, bson.M{"_id": transactionID}, update, opts).Decode(&updated); err != nil {
		return model.Transaction{}, fmt.Errorf("gagal update transaksi: %w", err)
	}

	log.Printf("[WA] Notifikasi pengiriman dikirim untuk transaksi %s\n", trx.ConsignmentNote)
	return updated, nil
}

// SendWADelivered
func (r *MTrans) SendWADelivered(ctx context.Context, transactionID primitive.ObjectID) (model.Transaction, error) {
	col := r.db.Collection("MailApp")

	var trx model.Transaction
	if err := col.FindOne(ctx, bson.M{"_id": transactionID}).Decode(&trx); err != nil {
		return model.Transaction{}, fmt.Errorf("transaksi tidak ditemukan: %w", err)
	}

	if trx.WADeliveredSent {
		return model.Transaction{}, fmt.Errorf("pesan WA 'delivered' sudah pernah dikirim")
	}

	now := time.Now()
	update := bson.M{
		"$set": bson.M{
			"wa_delivered_sent":    true,
			"wa_delivered_sent_at": now,
			"delivery_status":      "delivered",
			"updated_at":           now,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated model.Transaction
	if err := col.FindOneAndUpdate(ctx, bson.M{"_id": transactionID}, update, opts).Decode(&updated); err != nil {
		return model.Transaction{}, fmt.Errorf("gagal update transaksi: %w", err)
	}

	log.Printf("[WA] Notifikasi paket tiba dikirim untuk transaksi %s\n", trx.ConsignmentNote)
	return updated, nil
}

func (r *MTrans) DeleteTransaction(ctx context.Context, id string) (model.Transaction, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("id tidak valid: %w", err)
	}

	filter := bson.M{"_id": objectID}
	log.Println("[INFO] Menghapus transaksi dengan filter:", filter)

	var deleted model.Transaction
	col := r.db.Collection("MailApp")
	err = col.FindOneAndDelete(ctx, filter).Decode(&deleted)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Transaction{}, fmt.Errorf("transaksi dengan id %s tidak ditemukan", id)
		}
		return model.Transaction{}, fmt.Errorf("gagal menghapus transaksi: %w", err)
	}

	log.Printf("[INFO] Transaksi berhasil dihapus: %+v\n", deleted)
	return deleted, nil
}

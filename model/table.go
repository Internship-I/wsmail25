package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction represents a delivery transaction
type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConsignmentNote  string             `bson:"consignment_note,omitempty" json:"consignment_note,omitempty"`
	// Sender
	SenderName      string             `bson:"sender_name,omitempty" json:"sender_name,omitempty"`
	SenderPhone     string             `bson:"sender_phone,omitempty" json:"sender_phone,omitempty"`
	// Receiver
	ReceiverName    string             `bson:"receiver_name,omitempty" json:"receiver_name,omitempty"`
	AddressReceiver string             `bson:"address_receiver,omitempty" json:"address_receiver,omitempty"`
	ReceiverPhone   string             `bson:"receiver_phone,omitempty" json:"receiver_phone,omitempty"`
	// Package
	ItemContent     string             `bson:"item_content,omitempty" json:"item_content,omitempty"`
	ServiceType     string             `bson:"service_type,omitempty" json:"service_type,omitempty"`
	// Payment
	CODValue 		float64 			`bson:"cod_value,omitempty" json:"cod_value,omitempty"`  
	// Assignment
	CourierID       primitive.ObjectID  `bson:"courier_id,omitempty" json:"courier_id,omitempty"` 
	CourierName     string              `bson:"courier_name,omitempty" json:"courier_name,omitempty"` 
	// WhatsApp - Pesan 1: Saat akan dikirim (on delivery)
	WAOnDeliverySent   bool       		`bson:"wa_on_delivery_sent" json:"wa_on_delivery_sent"` // default false
	WAOnDeliverySentAt *time.Time 		`bson:"wa_on_delivery_sent_at,omitempty" json:"wa_on_delivery_sent_at,omitempty"`
	// WhatsApp - Pesan 2: Saat paket tiba (delivered)
	WADeliveredSent   bool      		`bson:"wa_delivered_sent" json:"wa_delivered_sent"` // default false
	WADeliveredSentAt *time.Time 		`bson:"wa_delivered_sent_at,omitempty" json:"wa_delivered_sent_at,omitempty"`
	// Delivery Status
    DeliveryStatus  string             `bson:"delivery_status" json:"delivery_status"`
	// Timestamps
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

// WALog for tracking WhatsApp message history
type WALog struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TransactionID primitive.ObjectID `bson:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	PhoneNumber   string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	MessageSent   string             `bson:"message_sent,omitempty" json:"message_sent,omitempty"`
	Status        string             `bson:"status,omitempty" json:"status,omitempty"` 
	ErrorMessage  string             `bson:"error_message,omitempty" json:"error_message,omitempty"`
	SentAt        time.Time          `bson:"sent_at,omitempty" json:"sent_at,omitempty"`
}	

// User represents a courier or admin
type Users struct {
	ID          string 				`bson:"_id,omitempty" json:"_id,omitempty"`
	Username    string             `bson:"username,omitempty" json:"username,omitempty"`
	PhoneNumber string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
	Role        string             `bson:"role,omitempty" json:"role,omitempty"` // "admin" or "kurir"
	CreatedAt 	time.Time 		   `bson:"created_at, omitempty" json:"created_at,omitempty"`
	UpdatedAt	time.Time		   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`	
}

// Token for authentication sessions
type Token struct {
	ID        string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Token     string    `bson:"token,omitempty" json:"token,omitempty"`
	AdminID   string    `bson:"admin_id,omitempty" json:"admin_id,omitempty"`
}
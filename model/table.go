package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction represents a delivery transaction
type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConsigmentNote  string             `bson:"consignment_note,omitempty" json:"consignment_note,omitempty"`

	// Sender
	SenderName      string             `bson:"sender_name,omitempty" json:"sender_name,omitempty"`
	SenderPhone     string             `bson:"sender_phone,omitempty" json:"sender_phone,omitempty"`

	// Receiver
	ReceiverName    string             `bson:"receiver_name,omitempty" json:"receiver_name,omitempty"`
	AddressReceiver string             `bson:"address_receiver,omitempty" json:"address_receiver,omitempty"`
	ReceiverPhone   string             `bson:"receiver_phone,omitempty" json:"receiver_phone,omitempty"`

	// Package
	ItemContent     string             `bson:"item_content,omitempty" json:"item_content,omitempty"`
	Weight          string             `bson:"weight,omitempty" json:"weight,omitempty"`
	ServiceType     string             `bson:"service_type,omitempty" json:"service_type,omitempty"`

	// Payment
	CODValue        float64            `bson:"cod_value,omitempty" json:"cod_value,omitempty"`
	IsPaid          bool               `bson:"is_paid,omitempty" json:"is_paid,omitempty"`

	// WhatsApp
	WASent          bool               `bson:"wa_sent,omitempty" json:"wa_sent,omitempty"`
	WASentAt        *time.Time         `bson:"wa_sent_at,omitempty" json:"wa_sent_at,omitempty"`

	// Timestamps
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

// MessageTemplate for WhatsApp notifications
type MessageTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TemplateName string             `bson:"template_name,omitempty" json:"template_name,omitempty"`
	TemplateCode string             `bson:"template_code,omitempty" json:"template_code,omitempty"` // akan_dikirim/sedang_dikirim/sampai
	TemplateText string             `bson:"template_text,omitempty" json:"template_text,omitempty"`
	Variables    []string           `bson:"variables,omitempty" json:"variables,omitempty"`
	IsActive     bool               `bson:"is_active,omitempty" json:"is_active,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

// WALog for tracking WhatsApp message history
type WALog struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TransactionID primitive.ObjectID `bson:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	PhoneNumber   string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	MessageSent   string             `bson:"message_sent,omitempty" json:"message_sent,omitempty"`
	Status        string             `bson:"status,omitempty" json:"status,omitempty"` // success/failed
	ErrorMessage  string             `bson:"error_message,omitempty" json:"error_message,omitempty"`
	SentAt        time.Time          `bson:"sent_at,omitempty" json:"sent_at,omitempty"`
}

// User represents a courier or admin
type Users struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName    string             `bson:"name,omitempty" json:"name,omitempty"`
	PhoneNumber string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Username    string             `bson:"username,omitempty" json:"username,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
	Role        string             `bson:"role,omitempty" json:"role,omitempty"` // "admin" or "kurir"
}

// Token for authentication sessions
type Token struct {
	ID        string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Token     string    `bson:"token,omitempty" json:"token,omitempty"`
	AdminID   string    `bson:"admin_id,omitempty" json:"admin_id,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
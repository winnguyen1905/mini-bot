package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents user table structure in a database.
type Voucher struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TelegramID int64              `bson:"telegram_id,omitempty"`
	Entry      string             `bson:"entry"`
	TakeProfit string             `bson:"take_profit"`
	StopLoss   string             `bson:"stop_loss"`
	Leverage   string             `bson:"leverage"`
	Volume     string             `bson:"volume"`
}

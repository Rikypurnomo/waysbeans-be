package models

import "time"

type Transaction struct {
	ID                 int                     `json:"id" gorm:"primary_key:auto_increment"`
	UserID             int                     `json:"-"`
	User               UserTransactionResponse `json:"user"`
	ProductTransaction []ProductTransaction    `json:"products" gorm:"foreignKey:TransactionID"`
	TotalQuantity      int                     `json:"total_quantity" gorm:"type: int"`
	TotalPrice         int                     `json:"total_price" gorm:"type: int"`
	CreatedAt          time.Time               `json:"-"`
	UpdatedAt          time.Time               `json:"-"`
}

type TransactionUSerResponse struct {
	ID            int `json:"id"`
	UserID        int `json:"-"`
	TotalQuantity int `json:"total_quantity"`
	TotalPrice    int `json:"total_price"`
}

func (TransactionUSerResponse) TableName() string {
	return "transactions"
}

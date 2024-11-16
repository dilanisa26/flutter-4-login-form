package models

type TransactionCategory struct {
	TransactionCategoryID int64  `json:"transaction_category_id" gorm:"primaryKey;autoIncrement;<-:false"`
	Name                  string `json:"name"`
}

// Untuk memastikan ORM menggunakan name tabel yang benar
func (TransactionCategory) TableName() string {
	return "transaction_categories"
}

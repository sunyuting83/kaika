package database

// Img List
type Card struct {
	ID          int64  `json:"id" gorm:"primary_key, column:id"`
	Card        string `json:"card" gorm:"varchar(128);index:idx_card_id;column:card"`
	Other       string `json:"other" gorm:"column:other;index:idx_other_id;column:other"`
	Status      int    `json:"status" gorm:"column:status;index:idx_status_id;column:status"`
	UpdateTime  int64  `json:"updatetime" gorm:"column:updatetime"`
	CreatedTime int64  `json:"createdtime" gorm:"column:createdtime"`
}

// TableName change table name
func (Card) TableName() string {
	return "card"
}

package database

// Img List
type Card struct {
	ID          int64  `json:"id" gorm:"primary_key, column:id"`
	Card        string `json:"card" gorm:"varchar(128);column:card"`
	UpdateTime  int64  `json:"updatetime" gorm:"column:updatetime"`
	CreatedTime int64  `json:"createdtime" gorm:"column:createdtime"`
}

// TableName change table name
func (Card) TableName() string {
	return "card"
}

// add admin
func (card *Card) Insert() (id int64, err error) {
	c := Eloquent.Create(&card)
	id = card.ID
	if c.Error != nil {
		err = c.Error
		return
	}
	return
}

// Delete Card
func (card *Card) DeleteCard(id string) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	Eloquent.Where("id = ?", id).Delete(&card)
}

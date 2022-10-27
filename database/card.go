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

// Get Count
func (card *Card) GetCount() (count int64, err error) {
	if err = Eloquent.Model(&card).Count(&count).Error; err != nil {
		return
	}
	return
}

// Card List
func (card *Card) GetCardList(page int64) (cards []Card, err error) {
	p := makePage(page)
	if err = Eloquent.
		Select("id, card, updatetime, createdtime").
		Order("id desc").
		Limit(100).Offset(p).
		Find(&cards).Error; err != nil {
		return
	}
	return
}

// Check ID
func (card *Card) CheckID(id int64) (cards Card, err error) {
	if err = Eloquent.First(&cards, "id = ?", id).Error; err != nil {
		return
	}
	return
}

// Updata to UpDateTime
func (card *Card) UpDateTime(id int64) (cards Card, err error) {
	if err = Eloquent.First(&cards, "id = ?", id).Error; err != nil {
		return
	}
	if err = Eloquent.Model(&cards).Updates(&card).Error; err != nil {
		return
	}
	return
}

// Search
func (card *Card) Search(c string) (cards Card, err error) {
	if err = Eloquent.First(&cards, "card = ?", c).Error; err != nil {
		return
	}
	return
}

// makePage make page
func makePage(p int64) int64 {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * 100
	return page
}

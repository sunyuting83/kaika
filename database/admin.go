package database

// Img List
type Admin struct {
	ID          int64  `json:"id" gorm:"primary_key, column:id"`
	Username    string `json:"username" gorm:"varchar(128);index:idx_username_id;column:username"`
	Password    string `json:"password" gorm:"column:password"`
	Status      int    `json:"status" gorm:"column:status"`
	UpdateTime  int64  `json:"updatetime" gorm:"column:updatetime"`
	CreatedTime int64  `json:"createdtime" gorm:"column:createdtime"`
}

// TableName change table name
func (Admin) TableName() string {
	return "admin"
}
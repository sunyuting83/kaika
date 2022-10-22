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

// add admin
func (admin *Admin) Insert() error {
	Eloquent.Create(&admin)
	return nil
}

// login
func (admin *Admin) CheckAdminLogin(username, password string) (admins Admin, err error) {
	if err = Eloquent.First(&admins, "username = ? AND status = ? AND password = ?", username, 0, password).Error; err != nil {
		return
	}
	return
}

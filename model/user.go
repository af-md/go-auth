package model

type User struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Gender   string `gorm:"column:gender"`
	Age      uint   `gorm:"column:age"`
}

package model

type User struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Gender   string `gorm:"column:gender"`
	Age      uint   `gorm:"column:age"`
	Location int    `gorm:"column:location"`
}

type Swipe struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	UserID       uint   `gorm:"column:UserId"`
	SwipedUserId uint   `gorm:"column:SwipedUserId"`
	Preference   string `gorm:"column:preference"`
}

type Match struct {
	ID      uint `gorm:"column:id;primaryKey"`
	UserId1 uint `gorm:"column:UserId1"`
	UserId2 uint `gorm:"column:UserId2"`
}

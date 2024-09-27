package user

type User struct{
	Id int `gorm:"column:id"  json:"id"`
	UserName string `gorm:"column:username" json:"username" `
	Password string	`gorm:"column:password_hash" json:"password" `
}
type UserSignIn struct {
	UserName string `gorm:"column:username" json:"username" `
	Password string	`gorm:"column:password_hash" json:"password" `
}
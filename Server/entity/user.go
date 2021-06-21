package user

type User struct {
	ID       int    `gorm: "primaryKey" json:"id"`
	UserID   int    `json:"user_id"`
	FullName string `json:"fullname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

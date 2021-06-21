package migration

type PassMan struct {
	ID int gorm "primaryKey"
	Website string
	Email string 
	Password string
	CreatedAt time.time 
	UpdatedAt time.time
}

type User struct {
	ID int gorm "primaryKey"
	FullName string
	Address string
	Email string
	Password string
}


package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Posts    []Post `json:"posts"` // One-to-many relationship
}

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	UserID uint   `json:"user_id"` // Foreign key referencing User.ID
	User   User
}

var Db *gorm.DB

func Connect() {
	var err error
	Db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	Db.AutoMigrate(&User{}, &Post{})
	fmt.Println("success to connectDb and migrateDb")
}

package models

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	DB_NAME string = "db.db"
)

type BaseModel struct {
	Id        int       `gorm:"AUTO_INCREMENT;primary_key"` // 自增
	CreatedAt time.Time // 列名为 `created_at`
	UpdatedAt time.Time // 列名为 `updated_at`
}

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file!")
	} else {

		DB_NAME = os.Getenv("DB_NAME")
	}
	DB, err = gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connsect database")
	}
	// db.AutoMigrate(.User{})
	return DB
}

//数据库迁移函数
func Migrate(tables Tables) {
	for _, table := range tables {
		log.Println(table.TableName())
		//t :=reflect.ValueOf(tables).Type()
		//tableDB := reflect.New(t).Elem()
		DB.AutoMigrate(table)
		if DB.Error != nil {
			log.Print(DB.Error)
		}
		log.Println("success!")
	}

}

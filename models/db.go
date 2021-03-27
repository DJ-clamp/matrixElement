package models

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	DB_NAME     string = "db.db"
	DB_TYPE     string = "sqlite"
	DB_USERNAME string = "root"
	DB_PASSWORD string = "root"
	DB_DATABASE string = "user"
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
		DB_TYPE = os.Getenv("DB_TYPE")
		DB_USERNAME = os.Getenv("DB_USERNAME")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_DATABASE = os.Getenv("DB_DATABASE")
	}
	if DB_TYPE == "mysql" {
		dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + DB_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	}
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

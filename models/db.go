package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type ORM struct {
	*gorm.DB
}

type BaseModel struct {
	Id        int       `gorm:"AUTO_INCREMENT;primary_key"` // 自增
	CreatedAt time.Time // 列名为 `created_at`
	UpdatedAt time.Time // 列名为 `updated_at`
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

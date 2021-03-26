package models

//数据库迁移模型索引
type Tables []interface{ TableName() string }

var TableList = Tables{
	User{},
}

func AddTables() {
	// TableList = append(TableList)
}

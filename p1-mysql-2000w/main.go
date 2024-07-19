package main

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Person 模型
type Person struct {
	ID          int       `gorm:"primaryKey;autoIncrement;comment:'主键'"`
	PersonID    uint8     `gorm:"not null;comment:'用户id'"`
	PersonName  string    `gorm:"size:200;comment:'用户名称'"`
	GmtCreate   time.Time `gorm:"comment:'创建时间'"`
	GmtModified time.Time `gorm:"comment:'修改时间'"`
}

// TableName 设置表名
func (Person) TableName() string {
	return "person"
}

func main() {
	// 连接数据库
	dsn := "root:@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 单条 sql 查询超大数据量测试，如何不设置超时时间，会持续读数据，mysql 服务端内存不会暴涨，cpu会涨。客户端内存会暴涨，cpu 也会暴涨
	var persons []Person
	result := db.WithContext(ctx).Find(&persons).Where("person_id = ?", 1)
	if result.Error != nil {
		log.Fatalf("failed to query data: %v", result.Error)
	}
	println(len(persons))
}

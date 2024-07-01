package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Member struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	TeamID   uint   `gorm:"not null"`
	No       uint   `gorm:"not null;unique"`
	Name     string `gorm:"type:varchar(64);default:'';not null"`
	JoinTime uint   `gorm:"not null"`
}

func (m *Member) TableName() string {
	return "member"
}

func main2() {
	// 模拟并发操作
	go func() {
		// 启动一个事务
		tx := db.Begin()
		defer tx.Commit()
		// 锁定 member 表
		tx.Exec("LOCK TABLES member WRITE")
		log.Println("第一个事务锁定了 member 表")
		// 插入操作
		newMember := Member{TeamID: 1, No: 101, Name: "John Doe", JoinTime: 1623157200}
		if err := tx.Create(&newMember).Error; err != nil {
			log.Fatalf("插入成员失败: %v", err)
		}
		log.Println("第一个事务插入了一条记录")
		// 模拟长时间操作
		time.Sleep(10 * time.Second)
		// 解锁表
		tx.Exec("UNLOCK TABLES")
		log.Println("第一个事务解锁了 member 表")
	}()

	time.Sleep(1 * time.Second)

	go func() {
		// 启动一个事务
		tx := db.Begin()
		defer tx.Commit()
		log.Println("第二个事务尝试锁定 member 表")
		// 尝试锁定 member 表
		tx.Exec("LOCK TABLES member WRITE")
		log.Println("第二个事务锁定了 member 表")
		// 插入操作
		newMember := Member{TeamID: 2, No: 102, Name: "Jane Doe", JoinTime: 1623157300}
		if err := tx.Create(&newMember).Error; err != nil {
			log.Fatalf("插入成员失败: %v", err)
		}
		log.Println("第二个事务插入了一条记录")

		// 解锁表
		tx.Exec("UNLOCK TABLES")
		log.Println("第二个事务解锁了 member 表")
	}()

	// 防止主 goroutine 提前退出
	time.Sleep(20 * time.Second)
}

func main() {
	// DSN (Data Source Name) 格式：用户名:密码@tcp(主机:端口)/数据库名?参数
	dsn := "root:@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 进行一些操作以确认连接成功
	var result int64
	db.Raw("SELECT 1").Scan(&result)
	if result == 1 {
		log.Println("连接成功！")
	} else {
		log.Println("连接失败！")
	}

	main2()
}

func print(res any) {
	s, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(s))
}

// 获取所有成员
func GetAllMembers(db *gorm.DB) ([]Member, error) {
	var members []Member
	result := db.Find(&members)
	print(result)
	return members, result.Error
}

// 根据ID获取成员
func GetMemberByID(db *gorm.DB, id uint) (*Member, error) {
	var member Member
	result := db.First(&member, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &member, nil
}

// 插入新成员
func InsertMember(db *gorm.DB, member *Member) error {
	result := db.Create(member)
	return result.Error
}

// 批量插入成员
func InsertMembers(db *gorm.DB, members []Member) error {
	result := db.Create(&members)
	return result.Error
}

// 根据ID删除成员
func DeleteMemberByID(db *gorm.DB, id uint) error {
	result := db.Delete(&Member{}, id)
	return result.Error
}

// 批量删除成员
func DeleteMembersByIDs(db *gorm.DB, ids []uint) error {
	result := db.Delete(&Member{}, ids)
	return result.Error
}

// 更新成员信息
func UpdateMember(db *gorm.DB, member *Member) error {
	result := db.Save(member)
	return result.Error
}

// 批量更新成员信息
func UpdateMembers(db *gorm.DB, members []Member) error {
	for _, member := range members {
		result := db.Save(&member)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

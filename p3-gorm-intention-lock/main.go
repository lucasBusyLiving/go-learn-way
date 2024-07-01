package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Member struct {
	ID        uint `gorm:"primaryKey"`
	TeamID    uint
	No        uint   `gorm:"unique"`
	Name      string `gorm:"size:64"`
	JoinTime  uint
	CreatedAt time.Time
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 自动迁移
	db.AutoMigrate(&Member{})

	// 开启两个事务模拟意向锁和行锁
	go transaction1(db)
	time.Sleep(2 * time.Second) // 确保第一个事务先运行
	go transaction2(db)

	// 等待一段时间以便查看结果
	time.Sleep(10 * time.Second)
}

func transaction1(db *gorm.DB) {
	tx := db.Begin()
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	// 获取意向独占锁（IX）并对某一行加独占锁（X）
	fmt.Println("Transaction 1: Acquiring IX lock on table and X lock on row")
	// 先表上加上意向独占锁，然后对读取的记录加独占锁
	// select ... for update;
	tx.Exec("SELECT * FROM member WHERE id = 1 FOR UPDATE")

	time.Sleep(5 * time.Second) // 模拟长时间操作

	fmt.Println("Transaction 1: Committing")
	tx.Commit()
}

func transaction2(db *gorm.DB) {
	tx := db.Begin()
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	// 尝试获取共享锁（S）
	fmt.Println("Transaction 2: Acquiring S lock on row")
	var member Member
	// 先在表上加上意向共享锁，然后对读取的记录加共享锁
	// select ... lock in share mode;
	tx.Set("gorm:query_option", "LOCK IN SHARE MODE").First(&member, 1)
	fmt.Printf("Transaction 2: Acquired S lock on row: %+v\n", member)

	fmt.Println("Transaction 2: Committing")
	tx.Commit()
}

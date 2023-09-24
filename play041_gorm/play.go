package play041_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func Play() {
	// 连接到MySQL数据库
	dsn := "root:gzn%zkTJ8x!gGZO6@tcp(mysql:3306)/biz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("gorm.Open: ", err)
	}

	// 自动迁移模型结构到数据库表
	db.AutoMigrate(&User{})

	// 创建新用户
	newUser := User{Name: "John", Age: 30}
	db.Create(&newUser)

	// 查询用户
	var foundUser User
	db.First(&foundUser, 1) // 通过ID查询第一个用户
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", foundUser.ID, foundUser.Name, foundUser.Age)

	// 更新用户信息
	db.Model(&foundUser).Update("Age", 31)

	// 删除用户
	//db.Delete(&foundUser)
}

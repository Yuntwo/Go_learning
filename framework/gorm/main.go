package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
	Age   int
}

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 从环境变量中获取配置信息
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 构造 DSN
	dsn := dbUser + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有 SQL 和调试信息
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected!")

	// 自动迁移，表名由 GORM 自动生成，默认是结构体名称的小写复数形式，这里是users
	// AutoMigrate 是 GORM 提供的一种自动迁移工具，用于检查数据库中是否存在与模型（User 结构体）对应的表结构。
	// 如果表不存在，AutoMigrate 会自动创建表。
	// 如果表已经存在，AutoMigrate 会检查表结构是否与模型一致，并根据模型更新表的字段（例如新增字段或更改字段类型，但是不支持删除字段）。
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println("AutoMigrate Failed!")
	}

	// 创建记录
	db.Create(&User{Name: "Alice", Email: "alice@example.com", Age: 25})

	// 读取记录
	var user User
	// 读取结果会存在user中
	result := db.First(&user, 1) // 根据主键查询
	if result.Error != nil {
		log.Println("Error retrieving user:", result.Error)
	} else {
		log.Printf("User retrieved: %+v\n", user)
	}

	result = db.First(&user, "email = ?", "alice@example.com")
	if result.Error != nil {
		log.Println("Error retrieving user by email:", result.Error)
	} else {
		log.Printf("User retrieved by email: %+v\n", user)
	}

	// 更新记录
	db.Model(&user).Update("Age", 30)

	// 删除记录
	db.Delete(&user)

}

package initialize

import (
	"log"

	"github.com/yhcui/web3study/task4/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GROM() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/person_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.User{}, &model.Posts{})
	db.AutoMigrate(&model.User{}, &model.Posts{}, &model.Comments{})
	return db
}

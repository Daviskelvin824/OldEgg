package model

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrate(){
	DB.AutoMigrate(&model.User{})
}

func connect(){
	psqlInfo := "host=localhost user=postgres password=postgres dbname=oldegg port=5432 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	migrate()
}

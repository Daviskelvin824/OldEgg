package config

import(
	"github.com/Daviskelvin824/OldEgg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrate(){
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.ProductCategory{})
	DB.AutoMigrate(&models.SearchQuery{})
	DB.AutoMigrate(&models.Shop{})
	DB.AutoMigrate(&models.OneTimeCode{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Wishlist{})
	DB.AutoMigrate(&models.WishlistDetail{})
	DB.AutoMigrate(&models.SavedForLaterItems{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.DeliveryProvider{})
	DB.AutoMigrate(&models.PaymentMethod{})
}

func Connect(){
	psqlInfo := "host=localhost user=postgres password=postgres dbname=oldegg port=5432 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	migrate()
}

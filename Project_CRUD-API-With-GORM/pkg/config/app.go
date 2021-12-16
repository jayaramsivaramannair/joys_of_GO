package config

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	db *gorm.DB
)


func Connect() {
	dsn := "host=localhost user=jayaramnair password=gorm dbname=simplerest port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	d, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic (err)
	}

	db = d
}


func GetDB() *gorm.DB {
	return db
}
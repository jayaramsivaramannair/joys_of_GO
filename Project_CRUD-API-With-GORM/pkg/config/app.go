package config

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	db *gorm.DB
)


func Connect() {
	dsn := "host=localhost user=jayaramnair password=jayaramnair dbname=jayaramnair port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic (err)
	}

	db = d
}


func GetDB() *gorm.DB {
	return db
}
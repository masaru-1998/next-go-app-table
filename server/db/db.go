package db

import (
				"gorm.io/driver/postgres"
				"gorm.io/gorm"
)

/*
DBへの接続を行う
*/
func Init() *gorm.DB{
				dsn := "host=db port=5432 user=postgres database=postgres password=password"
				db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
				if err != nil {
					panic(err)
				}

				return db
}

/*
* DBを閉じる
*/
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err = sqlDB.Close(); err != nil {
		panic(err)
	}
}
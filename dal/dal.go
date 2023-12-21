package dal

import (
	"fmt"
	"sync"

	"gorm.io/gorm"

	"known-anchors/config"

	"gorm.io/driver/mysql"

	"known-anchors/dal/db/model"
)

var DB *gorm.DB
var once sync.Once

func InitDB() {
	once.Do(func() {
		DB = ConnectDB(config.Conf.Mysql.DataSources).Debug()
		_ = DB.AutoMigrate(&model.User{}, &model.Card{}, &model.Comment{}, &model.Deck{}, &model.Post{})
	})
}

func ConnectDB(DataSources string) (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open(DataSources), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}

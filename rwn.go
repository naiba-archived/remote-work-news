package rwn

import (
	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB 数据库对象
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "data/rwn.db")
	if err != nil {
		panic(err)
	}
	DB = DB.Debug()
	DB.AutoMigrate(&News{})
}

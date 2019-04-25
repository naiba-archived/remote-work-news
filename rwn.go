package rwn

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Config 网站配置
type Config struct {
	BuildVersion string `mapstructure:"-"`
	ServerChan   string `mapstructure:"server_chan"`
	Debug        bool
}

// DB 数据库对象
var DB *gorm.DB

//C 全站配置
var C Config

//BuildVersion 构建版本
var BuildVersion = "_BuildVersion_"

func init() {
	var err error
	viper.SetConfigFile("data/rwn.yml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
	C.BuildVersion = BuildVersion[:8]

	DB, err = gorm.Open("sqlite3", "data/rwn.db")
	if err != nil {
		panic(err)
	}
	if C.Debug {
		DB = DB.Debug()
	}
	DB.AutoMigrate(&News{})
}

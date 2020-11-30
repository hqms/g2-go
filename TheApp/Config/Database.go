package Config

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type DBConfig struct {
	file string
}

func BuildDBConfig() *DBConfig{
	return &DBConfig{file: "thedb.db"}
}
func DbUrl(dbConfig *DBConfig) string{
	return fmt.Sprintf("%s/%s", os.Getwd(), dbConfig.file)
}
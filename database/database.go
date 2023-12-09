package database

import (
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var singleInstance *gorm.DB

func GetInstance() *gorm.DB {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			db, err := gorm.Open(mysql.New(mysql.Config{
				DSN: os.Getenv("DB_DNS"), // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
			}), &gorm.Config{})

			if err != nil {
				panic("Falha na conexão com o banco de dados!")
			}

			singleInstance = db
		}
	}

	return singleInstance
}

package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gorm.io/gorm"

	// MySQL接続Driver
	"gorm.io/driver/mysql"
)

var connection *gorm.DB
var once sync.Once

// NewDBConnection はDBのコネクションを返します。
func NewDBConnection() *gorm.DB {
	once.Do(func() {
		connection = createDBConnection()
	})

	return connection
}

// GetDBConnection はDBのコネクションを返します。
func GetDBConnection() *gorm.DB {
	// 最初にこの関数を使いDBコネクションを初期化した場合、
	// 以後上記の関数を使用してもDBコネクションが返る。
	// この関数も同様、最初に上記の関数を使用してDBコネクションを初期化すると、
	// この関数もプロダクションのDBコネクションを返す。
	once.Do(func() {
		connection = createDBConnection()
	})
	return connection
}

func createDBConnection() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")
	loc := "Asia%2FTokyo" // "/"はURLエンコーディングする
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s", user, password, host, port, dbname, loc)

	var conn *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		conn, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if err == nil {
			db, _ := conn.DB()
			if pingErr := db.Ping(); pingErr == nil {
				return conn
			}
		}
		time.Sleep(1 * time.Second)
	}

	panic(err)
}

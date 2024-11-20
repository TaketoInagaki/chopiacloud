package main

import (
	"app/environment"
	"app/environment/db"
	"os"
)

func main() {
	// DBコネクションを生成
	db := db.NewDBConnection()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	// echoインスタンスを生成
	app := environment.NewApp(db)

	// ポート番号を取得
	port := os.Getenv("PORT")
	if len(port) == 0 {
		panic("環境変数:PORTが未設定")
	}

	// サーバ起動
	app.Start(port)
}

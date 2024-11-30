package main

import (
	"database/sql"
	"fmt"
	"go_db/config"
	"go_db/database"
	"go_db/utils"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// データベースとの接続を管理する*sql.DB型のポインタを定義
var db *sql.DB

func init() {

	// iniのload
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// DBとの接続
	db, err = database.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	// 接続後にテーブル作成を実施
	if err := database.SetupTables(db); err != nil {
		log.Fatal(err)
	}

	log.Println("テーブルの準備が完了しました")

}

// inert 処理


func selectUser(db *sql.DB) error {
	query := `
	SELECT id, age, name, role
	FROM users ORDER BY id
	`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("クエリ実行エラー: %v", err)
	}
	defer rows.Close()

	type UserResult struct {
		ID   int
		Age  int
		Name string
		Role string
	}

	fmt.Println("-----ユーザ一覧-------")
	for rows.Next() {
		var user UserResult
		err := rows.Scan(&user.ID, &user.Age, &user.Name, &user.Role)
		if err != nil {
			return fmt.Errorf("読み取りエラー: %v", err)
		}
		fmt.Printf("id: %d, age: %d, name: %s, role: %s", user.ID, user.Age, user.Name, user.Role)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("結果セット処理エラー: %v", err)
	}

	return nil

}

func main() {

	defer db.Close()
	
	// コマンドライン引数から値を取得
	filename, err := utils.GetFilePath()
	if err != nil {
		log.Fatal(err)
	}
	
	// ファイル取得
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	// 取得したファイルからデータを読み込み、insert
   if err := database.LogImport(file, db); err != nil {
		log.Fatal(err)
	 }
	
	selectUser(db)

}

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/xerrors"
)

// DB は、データベースを表す。
type DB struct {
	Conn *gorm.DB
}

// NewDB は、DB を生成し、返す。
func NewDB() (*DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(nvgdb:3306)/nuxt_vue_go_chat?charset=utf8mb4&parseTime=True") // TODO 後でちゃんとした形にする
	if err != nil {
		return nil, xerrors.Errorf("failed to connect database: %w", err)
	}

	// TODO ここに SET や Table 作成を行う

	db.LogMode(true)

	return &DB{
		Conn: db,
	}, nil
}

// CloseDB は、DB を close する。
func (db *DB) CloseDB() error {
	return db.Conn.Close()
}

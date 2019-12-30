package datastore

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sekky0905/modern-chat/domain/repository"
	"golang.org/x/xerrors"
)

// db は、データベースを表す。
type db struct {
	conn *gorm.DB
}

// NewDB は、DB を生成し、返す。
func NewDB() (repository.DB, error) {
	conn, err := gorm.Open("mysql", "root:@tcp(nvgdb:3306)/nuxt_vue_go_chat?charset=utf8mb4&parseTime=True") // TODO 後でちゃんとした形にする
	if err != nil {
		return nil, xerrors.Errorf("failed to connect database: %w", err)
	}

	// TODO ここに SET や Table 作成を行う

	conn.LogMode(true)

	return &db{
		conn: conn,
	}, nil
}

// Close
func (db *db) Close() error {
	return db.conn.Close()
}

// DB
func (db *db) DB() *sql.DB {
	return db.conn.DB()
}

// Where
func (db *db) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.conn.Where(query, args)
}

// Or
func (db *db) Or(query interface{}, args ...interface{}) *gorm.DB {
	return db.conn.Or(query, args)
}

// Not
func (db *db) Not(query interface{}, args ...interface{}) *gorm.DB {
	return db.conn.Not(query, args)
}

// Limit
func (db *db) Limit(limit interface{}) *gorm.DB {
	return db.conn.Limit(limit)
}

// Offset
func (db *db) Offset(offset interface{}) *gorm.DB {
	return db.conn.Offset(offset)
}

// Order
func (db *db) Order(value interface{}, reorder ...bool) *gorm.DB {
	return db.conn.Order(value, reorder...)
}

// Select
func (db *db) Select(query interface{}, args ...interface{}) *gorm.DB {
	return db.conn.Select(query, args)
}

// Omit
func (db *db) Omit(columns ...string) *gorm.DB {
	return db.conn.Omit(columns...)
}

// Group
func (db *db) Group(query string) *gorm.DB {
	return db.conn.Group(query)
}

// Having
func (db *db) Having(query interface{}, values ...interface{}) *gorm.DB {
	return db.conn.Having(query, values...)
}

// Joins
func (db *db) Joins(query string, args ...interface{}) *gorm.DB {
	return db.conn.Joins(query, args)
}

// First
func (db *db) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.conn.Find(out, where)
}

// Take
func (db *db) Take(out interface{}, where ...interface{}) *gorm.DB {
	return db.conn.Take(out, where...)
}

// Last
func (db *db) Last(out interface{}, where ...interface{}) *gorm.DB {
	return db.conn.Last(out, where...)
}

// Find
func (db *db) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.conn.Find(out, where...)
}

// Preloads
func (db *db) Preloads(out interface{}) *gorm.DB {
	return db.conn.Preloads(out)
}

// Scan
func (db *db) Scan(dest interface{}) *gorm.DB {
	return db.conn.Scan(dest)
}

// Row
func (db *db) Row() *sql.Row {
	return db.conn.Row()
}

// Rows
func (db *db) Rows() (*sql.Rows, error) {
	return db.conn.Rows()
}

// ScanRows
func (db *db) ScanRows(rows *sql.Rows, result interface{}) error {
	return db.ScanRows(rows, result)
}

// Pluck
func (db *db) Pluck(column string, value interface{}) *gorm.DB {
	return db.conn.Pluck(column, value)
}

// Count
func (db *db) Count(value interface{}) *gorm.DB {
	return db.conn.Count(value)
}

// Related
func (db *db) Related(value interface{}, foreignKeys ...string) *gorm.DB {
	return db.conn.Related(value, foreignKeys...)
}

// FirstOrInit
func (db *db) FirstOrInit(out interface{}, where ...interface{}) *gorm.DB {
	return db.conn.FirstOrInit(out, where...)
}

// FirstOrCreate
func (db *db) FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB {
	return db.conn.FirstOrCreate(out, where...)
}

// Update
func (db *db) Update(attrs ...interface{}) *gorm.DB {
	return db.conn.Update(attrs)
}

// Updates
func (db *db) Updates(values interface{}, ignoreProtectedAttrs ...bool) *gorm.DB {
	return db.conn.Updates(values, ignoreProtectedAttrs...)
}

// UpdateColumn
func (db *db) UpdateColumn(attrs ...interface{}) *gorm.DB {
	return db.conn.UpdateColumn(attrs...)
}

// UpdateColumns
func (db *db) UpdateColumns(values interface{}) *gorm.DB {
	return db.conn.UpdateColumns(values)
}

// Save
func (db *db) Save(value interface{}) *gorm.DB {
	return db.conn.Save(value)
}

// Create
func (db *db) Create(value interface{}) *gorm.DB {
	return db.conn.Create(value)
}

// Delete
func (db *db) Delete(value interface{}, where ...interface{}) *gorm.DB {
	return db.conn.Delete(value, where...)
}

// Raw
func (db *db) Raw(sql string, values ...interface{}) *gorm.DB {
	return db.conn.Raw(sql, values...)
}

// Exec
func (db *db) Exec(sql string, values ...interface{}) *gorm.DB {
	return db.conn.Exec(sql, values...)
}

// Begin
func (db *db) Begin() *gorm.DB {
	return db.conn.Begin()
}

// BeginTx
func (db *db) BeginTx(ctx context.Context, opts *sql.TxOptions) *gorm.DB {
	return db.conn.BeginTx(ctx, opts)
}

// Commit
func (db *db) Commit() *gorm.DB {
	return db.conn.Commit()
}

// Rollback
func (db *db) Rollback() *gorm.DB {
	return db.conn.Rollback()
}

// RollbackUnlessCommitted
func (db *db) RollbackUnlessCommitted() *gorm.DB {
	return db.conn.RollbackUnlessCommitted()
}

// NewRecord
func (db *db) NewRecord(value interface{}) bool {
	return db.conn.NewRecord(value)
}

// RecordNotFound
func (db *db) RecordNotFound() bool {
	return db.conn.RecordNotFound()
}

// Preload
func (db *db) Preload(column string, conditions ...interface{}) *gorm.DB {
	return db.conn.Preload(column, conditions...)
}

// Get
func (db *db) Get(name string) (value interface{}, ok bool) {
	return db.conn.Get(name)
}

// AddError
func (db *db) AddError(err error) error {
	return db.conn.AddError(err)
}

// GetErrors (
func (db *db) GetErrors() []error {
	return db.conn.GetErrors()
}

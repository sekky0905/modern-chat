package repository

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
)

// DB は、Database を表す。
type DB interface {
	Close() error
	DB() *sql.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Or(query interface{}, args ...interface{}) *gorm.DB
	Not(query interface{}, args ...interface{}) *gorm.DB
	Limit(limit interface{}) *gorm.DB
	Offset(offset interface{}) *gorm.DB
	Order(value interface{}, reorder ...bool) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Omit(columns ...string) *gorm.DB
	Group(query string) *gorm.DB
	Having(query interface{}, values ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Take(out interface{}, where ...interface{}) *gorm.DB
	Last(out interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Preloads(out interface{}) *gorm.DB
	Scan(dest interface{}) *gorm.DB
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) *gorm.DB
	Count(value interface{}) *gorm.DB
	Related(value interface{}, foreignKeys ...string) *gorm.DB
	FirstOrInit(out interface{}, where ...interface{}) *gorm.DB
	FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	Updates(values interface{}, ignoreProtectedAttrs ...bool) *gorm.DB
	UpdateColumn(attrs ...interface{}) *gorm.DB
	UpdateColumns(values interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	Begin() *gorm.DB
	BeginTx(ctx context.Context, opts *sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
	RollbackUnlessCommitted() *gorm.DB
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	Preload(column string, conditions ...interface{}) *gorm.DB
	Get(name string) (value interface{}, ok bool)
	AddError(err error) error
	GetErrors() []error
}

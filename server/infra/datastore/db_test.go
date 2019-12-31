package datastore

import (
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/ory/dockertest"
	"github.com/sekky0905/modern-chat/util"
)

// dbMock は、db の mock。
type dbMock struct {
	conn *gorm.DB
}

func (db dbMock) createTables() {
	db.conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}, &ChatRoom{}, &Comment{}, &Like{})
}

func (db dbMock) truncateTables() {
	tables := []string{"users", "chat_rooms", "comments", "likes"}

	for _, table := range tables {
		db.conn.Exec("TRUNCATE TABLE ?", table)
	}
}

// DBMock は、DB の Mock。
var DBMock *dbMock

func TestMain(m *testing.M) {
	pool, resource := initDB()
	code := m.Run()

	closeDB(pool, resource)
	os.Exit(code)
}

// refs: https://github.com/ory/dockertest
func initDB() (*dockertest.Pool, *dockertest.Resource) {
	defaultEndpoint := ""
	pool, err := dockertest.NewPool(defaultEndpoint)
	if err != nil {
		phrase := fmt.Sprintf("Could not connect to docker: %s", err)
		util.Logger().Error(phrase)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		phrase := fmt.Sprintf("Could not start resource: %s", err)
		util.Logger().Error(phrase)
	}

	db := &dbMock{}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		param := "parseTime=true&loc=Asia%2FTokyo&time_zone=%27%2b9%3a00%27&charset=utf8mb4"
		dataSource := fmt.Sprintf("root:secret@(localhost:%s)/mysql?%s", resource.GetPort("3306/tcp"), param)
		db.conn, err = gorm.Open("mysql", dataSource)
		if err != nil {
			return err
		}
		return db.conn.DB().Ping()
	}); err != nil {
		phrase := fmt.Sprintf("Could not connect to docker: %s", err)
		util.Logger().Error(phrase)
	}

	db.createTables()

	return pool, resource
}

func closeDB(pool *dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		phrase := fmt.Sprintf("Could not purge resource: %s", err)
		util.Logger().Error(phrase)
	}
}

package main

//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

//需要抛给上层，方便溯源

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

type Db struct {
	Dsn      string
	Db       *sql.DB
	UserInfo user
}
type user struct {
	Id   int
	Name sql.NullString
	Age  sql.NullInt64
}

func main() {
	var err error
	db := Db{
		Dsn: "root:123456@tcp(localhost:3306)/sqlx_db?charset=utf8mb4",
	}
	db.Db, err = sql.Open("mysql", db.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Db.Close()

	err = db.insertData()
	if err != nil {
		log.Fatalf("err: %v", errors.Cause(err))
	}
}

func (db *Db) insertData() error {
	stmt, _ := db.Db.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`)
	defer stmt.Close()

	_, err := stmt.Exec("xys", 123)

	return errors.Wrap(err, "insert data failed")

}

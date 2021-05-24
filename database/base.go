package database

import (
	"database/sql"
	"fmt"
	"log"

	//ドライバの登録
	_ "github.com/mattn/go-sqlite3"
)

type DbError struct{
	Cmd string
	Err error
}

func (err *DbError)Error()string{
	return err.Err.Error()
}

const(
	dbName ="database.db"
	dbDriver ="sqlite"
	tableName="phone_book"
)

var dbConn *sql.DB

func init(){
	var err error

	dbConn,err=sql.Open(dbDriver,dbName)
	
	if err!=nil{
		log.Fatalf("dbName=%s, dbDriver=%s\nerr: %#v",dbName,dbDriver,err)
	}

	cmd:=fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
			id INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,
			name VARCHAR(30) NOT NULL
			phone_number VARCHAR(30) NOT NULL
			PRIMARY KEY (id)
		)`,tableName)

	if _,err:=dbConn.Exec(cmd);err!=nil{
		log.Fatalf("dbName=%s, dbDriver=%s\nerr: %#v",dbName,dbDriver,err)		
	}

	GetAllUser()
}
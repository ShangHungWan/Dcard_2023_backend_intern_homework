package db

import (
	"database/sql"
	"fmt"
	_ "key-value-system/env"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const TABLE string = "nodes"
const CLEAR_OLD_NODES_SQL string = "DELETE FROM %[1]s WHERE key IN (SELECT key FROM %[1]s WHERE value IS NOT NULL AND updated_at < now() - interval '1 day')"
const INSERT_HEAD_SQL string = "INSERT INTO %[1]s (key) VALUES ($1)"
const INSERT_NODE_SQL string = "INSERT INTO %[1]s (key, value) VALUES ($1, $2)"
const UPDATE_NODE_SQL string = "UPDATE %[1]s SET next=$1 WHERE key=$2"
const GET_HEAD_SQL string = "SELECT key, next, created_at, updated_at FROM %[1]s WHERE key = $1 and value IS NULL"
const GET_NODE_SQL string = "SELECT key, value, next, created_at, updated_at FROM %[1]s WHERE key = $1 and value IS NOT NULL"
const DELETE_NODE_SQL string = "DELETE FROM %[1]s WHERE key = $1"
const DELETE_NODES_SQL string = "DELETE FROM %[1]s WHERE key IN ($1)"

var DB *sql.DB

func init() {
	connStr := fmt.Sprintf("host= %s port= %s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", connStr)
	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	DB = db
}

func PrepareAndExec(sql string, args ...any) (sql.Result, error) {
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return nil, err
	}

	return stmt.Exec(args...)
}

func GetSql(sql string) string {
	return fmt.Sprintf(sql, TABLE)
}

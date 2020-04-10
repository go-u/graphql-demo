package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(projectID string) *SqlHandler {
	source := getSource(projectID)
	con := createConnection(source)
	sqlHandler := &SqlHandler{
		Conn: con,
	}
	return sqlHandler
}

func getSource(projectID string) string {
	// Local
	const (
		Name = "graph"
		Pass = "pass"
		Host = "127.0.0.1"
		Port = "5555"
	)
	return fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=true&collation=utf8mb4_bin", Pass, Host, Port, Name)
}

func createConnection(Source string) *sql.DB {
	con, err := sql.Open("mysql", Source)
	// https://blog.nownabe.com/2017/01/16/570.html#accessing-the-database
	// defer con.Close()
	if err != nil {
		log.Fatal("DB Open Error: ", err)
	}
	log.Println("DB initialized")
	return con
}

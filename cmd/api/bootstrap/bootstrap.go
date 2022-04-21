package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "demo"
	dbPass = "demo"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "demo"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}

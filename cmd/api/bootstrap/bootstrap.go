package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mooc "github.com/juansecardozo/hexagonal-http-api/internal"
	"github.com/juansecardozo/hexagonal-http-api/internal/creating"
	"github.com/juansecardozo/hexagonal-http-api/internal/increasing"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/bus/inmemory"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/storage/mysql"
	"time"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "demo"
	dbPass    = "demo"
	dbHost    = "localhost"
	dbPort    = "6603"
	dbName    = "demo"
	dbTimeout = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseService := increasing.NewCourseCounterIncreaseService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}

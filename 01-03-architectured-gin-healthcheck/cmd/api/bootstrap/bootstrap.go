package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/creating"
	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/platform/bus/inmemory"
	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/platform/server"
	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "codely"
	dbPass    = "codely"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "codely"
	dbTimeout = 5 * time.Second
)

func Run() error {

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var commandBus = inmemory.NewCommandBus()

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)

	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}

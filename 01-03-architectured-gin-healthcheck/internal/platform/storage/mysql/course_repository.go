package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal"
	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a MYSQL mooc.CourseRepository
type CourseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// CourseRepository initializes a MYSQL-based implementation of mooc.CourseRepository
func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save initializes the mooc.CourseRepository
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}
	return nil
}

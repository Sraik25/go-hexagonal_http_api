package mysql

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal"
	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a MYSQL mooc.CourseRepository
type CourseRepository struct {
	db *sql.DB
}

// CourseRepository initializes a MYSQL-based implementation of mooc.CourseRepository
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
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

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}
	return nil
}

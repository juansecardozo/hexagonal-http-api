package mysql

import (
	"database/sql"
	mooc "github.com/juansecardozo/hexagonal-http-api/internal"
)

func NewCourseRepository(_ *sql.DB) mooc.CourseRepository {
	return nil
}

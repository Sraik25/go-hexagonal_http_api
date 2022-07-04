package creating

import (
	"context"
	"errors"

	mooc "github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal"
	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/increasing"
	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increaserService increasing.CourseCounterIncreaserService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing.CourseCounterIncreaserService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increaserService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increaserService.Increase(courseCreatedEvt.ID())
}

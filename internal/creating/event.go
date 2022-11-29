package creating

import (
	"context"
	"errors"
	mooc "github.com/juansecardozo/hexagonal-http-api/internal"
	"github.com/juansecardozo/hexagonal-http-api/internal/increasing"
	"github.com/juansecardozo/hexagonal-http-api/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increaseService increasing.CourseCounterIncreaseService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaseService increasing.CourseCounterIncreaseService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increaseService: increaseService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increaseService.Increase(courseCreatedEvt.ID())
}

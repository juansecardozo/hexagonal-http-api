package creating

import (
	"context"
	"errors"
	mooc "github.com/juansecardozo/hexagonal-http-api/internal"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/storage/storagemocks"
	"github.com/juansecardozo/hexagonal-http-api/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCourseService_CreateCourse_RepoError(t *testing.T) {
	courseID, courseName, courseDuration := "a2cdff46-9d78-4aa3-8497-c4f1379656d6", "Test Course", "1 hour"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(errors.New("something failed"))

	eventBusMock := new(eventmocks.Bus)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func TestCourseService_CreateCourse_EventBusError(t *testing.T) {
	courseID, courseName, courseDuration := "a2cdff46-9d78-4aa3-8497-c4f1379656d6", "Test Course", "1 hour"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something went wrong"))

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func TestCourseService_CreateCourse_Succeed(t *testing.T) {
	courseID, courseName, courseDuration := "a2cdff46-9d78-4aa3-8497-c4f1379656d6", "Test Course", "1 hour"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}

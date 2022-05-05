package increasing

type CourseCounterIncreaseService struct {
}

func NewCourseCounterIncreaseService() CourseCounterIncreaseService {
	return CourseCounterIncreaseService{}
}

func (s CourseCounterIncreaseService) Increase(id string) error {
	return nil
}

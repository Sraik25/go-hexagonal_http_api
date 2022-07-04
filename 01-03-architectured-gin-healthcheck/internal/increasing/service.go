package increasing

type CourseCounterIncreaserService struct {
}

func NewCourseCounterIncreaserService() CourseCounterIncreaserService {
	return CourseCounterIncreaserService{}
}

func (s CourseCounterIncreaserService) Increase(id string) error {
	return nil
}

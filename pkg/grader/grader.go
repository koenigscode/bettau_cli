package grader

type GradeQuery struct {
	Question string
	Input    string
}
type GradeResult struct {
	Correct  bool
	Solution string
	Feedback string
}

type Grader interface {
	Grade(GradeQuery) GradeResult
}

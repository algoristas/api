package problems

// DataProvider defines data API for problems.
type DataProvider interface {
	GetSets() ([]byte, error)
	FindProblem(userID, problemID string) (*Problem, error)
}

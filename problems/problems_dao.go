package problems

// DAO defines Data Access Object for problems.
type DAO interface {
	GetSets() ([]byte, error)
}

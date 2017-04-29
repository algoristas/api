package results

// DAO defines Data Access Object for results.
type DAO interface {
	GetResults() ([]byte, error)
}

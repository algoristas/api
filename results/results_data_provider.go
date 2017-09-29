package results

// DataProvider defines data API for results.
type DataProvider interface {
	GetResults() ([]byte, error)
}

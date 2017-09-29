package standings

// DataProvider defines data API for standings.
type DataProvider interface {
	GetStandings() ([]byte, error)
}

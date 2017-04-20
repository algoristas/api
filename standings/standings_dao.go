package standings

// DAO defines Data Access Object for standings.
type DAO interface {
	GetStandings() ([]byte, error)
}

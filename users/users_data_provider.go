package users

// DataProvider defines the data API for users.
type DataProvider interface {
	FindUserByID(id uint) (*User, error)
	FindUserByUserName(userName string) (*User, error)
}

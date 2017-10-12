package users

// DataProvider defines the data API for users.
type DataProvider interface {
	FindUser(id string) (*User, error)
	FindUserByID(id uint) (*User, error)
	FindUserByUserName(userName string) (*User, error)
}

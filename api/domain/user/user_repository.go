package user

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	FetchByEmail(email string) error
}

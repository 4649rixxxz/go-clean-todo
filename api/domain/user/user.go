package user

import (
	"regexp"
	"time"
	"unicode/utf8"

	errDomain "go-clean-todo/domain/error"
)

type User struct {
	userID    uint
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

const (
	passwordLengthMin = 10
	passwordLengthMax = 30
)

func NewUser(
	email string,
	password string,
) (*User, errDomain.DomainErrorI) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return nil, errDomain.NewDomainError("email", "メールアドレスが不正です。")
	}
	passwordLength := utf8.RuneCountInString(password)
	if passwordLength < passwordLengthMin || passwordLength > passwordLengthMax {
		return nil, errDomain.NewDomainError("password", "パスワードの長さは、10文字以上30文字以内でお願いします。")
	}

	return &User{
		userID:    0,
		email:     email,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func Reconstruct(
	userID uint,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		userID:    userID,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) UserID() uint {
	return u.userID
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

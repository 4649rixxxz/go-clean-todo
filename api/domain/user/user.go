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
	deletedAt *time.Time
}

const (
	passwordLengthMin = 10
	passwordLengthMax = 30
)

func NewUser(
	email string,
	password string,
) (*User, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return nil, errDomain.NewError("メールアドレスが不正です。")
	}
	passwordLength := utf8.RuneCountInString(password)
	if passwordLength < passwordLengthMin || passwordLength > passwordLengthMax {
		return nil, errDomain.NewError("パスワードの長さは、10文字以上30文字以内でお願いします。")
	}

	return &User{
		userID:    0,
		email:     email,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: nil,
	}, nil
}

func Reconstruct(
	userID uint,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *User {
	return &User{
		userID:    userID,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
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

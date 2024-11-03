package user

import (
	"fmt"
	userDomain "go-clean-todo/domain/user"
)

type SigninUserUsecase struct {
	userRepo userDomain.UserRepository
}

func NewSigninUserUsecase(
	userRepo userDomain.UserRepository,
) *SigninUserUsecase {
	return &SigninUserUsecase{
		userRepo: userRepo,
	}
}

func (uc *SigninUserUsecase) Run() {
	fmt.Println("サインインの処理")
}

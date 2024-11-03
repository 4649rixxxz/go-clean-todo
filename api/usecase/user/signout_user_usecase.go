package user

import (
	"fmt"
	userDomain "go-clean-todo/domain/user"
)

type SignoutUserUsecase struct {
	userRepo userDomain.UserRepository
}

func NewSignoutUserUsecase(
	userRepo userDomain.UserRepository,
) *SignoutUserUsecase {
	return &SignoutUserUsecase{
		userRepo: userRepo,
	}
}

func (uc *SignoutUserUsecase) Run() {
	fmt.Println("サインアウトの処理")
}

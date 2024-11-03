package user

import (
	userDomain "go-clean-todo/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type SignupUserUsecase struct {
	userRepo userDomain.UserRepository
}

func NewSignupUserUsecase(
	userRepo userDomain.UserRepository,
) *SignupUserUsecase {
	return &SignupUserUsecase{
		userRepo: userRepo,
	}
}

func (uc *SignupUserUsecase) Run(dto SignupUserUsecaseDTO) error {
	user, err := userDomain.NewUser(
		dto.Email,
		dto.Password,
	)
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return err
	}
	user.SetPassword(string(hash))
	if _, createdErr := uc.userRepo.CreateUser(user); createdErr != nil {
		return createdErr
	}

	return nil
}

type SignupUserUsecaseDTO struct {
	Email    string
	Password string
}

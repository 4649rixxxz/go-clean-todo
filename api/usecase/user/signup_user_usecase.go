package user

import (
	userDomain "go-clean-todo/domain/user"
	"go-clean-todo/usecase"

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

func (uc *SignupUserUsecase) Run(dto SignupUserUsecaseDTO) usecase.UsecaseErrorI {
	user, err := userDomain.NewUser(
		dto.Email,
		dto.Password,
	)
	if err != nil {
		return usecase.NewInvalidInputError(err.Field(), err.Error())
	}
	record, fetchErr := uc.userRepo.FetchByEmail(user.Email())
	if record != nil {
		return usecase.NewInvalidInputError("email", "このメールアドレスはすでに使用されています。")
	}
	signupErrMsg := "サインアップに失敗しました。お手数ですが、もう一度お試しください。"
	if fetchErr != nil {
		return usecase.NewInternalServerError(signupErrMsg)
	}
	hash, passwordErr := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if passwordErr != nil {
		return usecase.NewInternalServerError(signupErrMsg)
	}
	user.SetPassword(string(hash))
	if _, createdErr := uc.userRepo.CreateUser(user); createdErr != nil {
		return usecase.NewInternalServerError(signupErrMsg)
	}

	return nil
}

type SignupUserUsecaseDTO struct {
	Email    string
	Password string
}

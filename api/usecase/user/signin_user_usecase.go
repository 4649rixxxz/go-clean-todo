package user

import (
	userDomain "go-clean-todo/domain/user"
	"go-clean-todo/usecase"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SigninUserUsecase struct {
	userRepo userDomain.UserRepository
}

type SigninUserUsecaseDTO struct {
	Email    string
	Password string
}

func NewSigninUserUsecase(
	userRepo userDomain.UserRepository,
) *SigninUserUsecase {
	return &SigninUserUsecase{
		userRepo: userRepo,
	}
}

func (uc *SigninUserUsecase) Run(dto SigninUserUsecaseDTO) (string, usecase.UsecaseErrorI) {
	user, err := uc.userRepo.FetchByEmail(dto.Email)
	if err != nil {
		return "", usecase.NewInvalidInputError("email", "メールアドレスまたはパスワードが間違っています。")
	}
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(dto.Password))
	if passErr != nil {
		return "", usecase.NewInvalidInputError("password", "メールアドレスまたはパスワードが間違っています。")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("SECRET")))
	if tokenErr != nil {
		return "", usecase.NewInternalServerError("認証に失敗しました。")
	}

	return tokenString, nil
}

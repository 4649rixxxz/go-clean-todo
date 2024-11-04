package user

import (
	userDomain "go-clean-todo/domain/user"
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

func (uc *SigninUserUsecase) Run(dto SigninUserUsecaseDTO) (string, error) {
	user, err := uc.userRepo.FetchByEmail(dto.Email)
	if err != nil {
		return "", err
	}
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(dto.Password))
	if passErr != nil {
		return "", passErr
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("SECRET")))
	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenString, nil
}

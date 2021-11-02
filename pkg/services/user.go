package services

import (
	"GoRestApi/pkg/config"
	"GoRestApi/pkg/data"
	"GoRestApi/pkg/domain"
	"GoRestApi/pkg/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateAccount(user *domain.User) *models.Error
	Login(user *domain.User) (string, *models.Error)
}

type UserService struct {
	userProvider data.IUserProvider
	cfg          *config.Settings
}

func NewUserService(cfg *config.Settings, userProvider data.IUserProvider) IUserService {
	return &UserService{
		userProvider: userProvider,
		cfg:          cfg,
	}
}

func (u UserService) CreateAccount(user *domain.User) *models.Error {
	userExists, err := u.userProvider.UsernameExists(user.Username)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err,
		}
	}

	if userExists {
		return &models.Error{
			Code:    400,
			Name:    "USERNAME_TAKEN",
			Message: "Username already exists",
		}
	}

	user.ID = primitive.NewObjectID()
	hash, err := hashPassword(user.Password)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err,
		}
	}

	user.Password = hash
	err = u.userProvider.CreateAcount(user)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err,
		}
	}
	return nil

}

func (u UserService) Login(user *domain.User) (string, *models.Error) {
	userFound, err := u.userProvider.FindByUsername(user.Username)
	if err != nil {
		return "", &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err,
		}
	}

	if userFound == nil {
		return "", &models.Error{
			Code:    400,
			Name:    "INVALID_LOGIN",
			Message: "You username or pasword is invalid",
		}
	}

	err = comparePasswordWithHash(user.Password, userFound.Password)

	if err != nil {
		return "", &models.Error{
			Code:    400,
			Name:    "INVALID_LOGIN",
			Message: "Your username or password is invalid.",
		}
	}

	token, err := u.createJwtToken(userFound.ID.Hex())

	if err != nil {
		return "", &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err,
		}
	}

	return token, nil
}

func hashPassword(password string) (string, error) {

	passwordBytes := []byte(password)
	hashPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 12)

	if err != nil {
		return "", errors.Wrap(err, "Error creating password")
	}

	return string(hashPassword), nil
}

func comparePasswordWithHash(password string, hash string) error {

	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)

	return errors.Wrap(err, "error comparing password hash")
}

func (u UserService) createJwtToken(userId string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	expiresIn, err := strconv.ParseInt(u.cfg.JwtExpires, 10, 64)

	if err != nil {
		return "", errors.Wrap(err, "Error parsing int")
	}

	expiration := time.Duration(int64(time.Minute) * expiresIn)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = userId

	claims["exp"] = time.Now().Add(expiration).Unix()

	t, err := token.SignedString([]byte(u.cfg.JwtSecret))

	if err != nil {
		return "", errors.Wrap(err, "Error signing JWT token")
	}

	return t, nil
}

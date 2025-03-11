package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface{
	//function(args)(return)
.	//後半の()は返り値を複数返す時の書き方らしいこの場合はerrorも返す形
	SignUp(user model.User)(model.UserResponse,error)
	Login(user model.User)(string,error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User)(model.UserResponse, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return model.UserResponse{},err
	}
	newUser := model.User{Email:user.Email, Password:string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: newUser.ID
		Email: newUser.Email
	}
	return resUser
}

func (uu *userUsecase)Login(user model.User)(string, error){
	storeUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "",err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storeUser.Password),[]byte(user.Password))
	if err != nil {
		return "",err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims){
		"user_id": storeUser.ID
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}
	tokenString,err := token.SignedString([]byte(os.Getenv("SERECT")))
	return tokenString,nil
}
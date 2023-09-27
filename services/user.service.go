package services

import (
	"context"
	"fmt"
	"strings"
	"tecnotree-go-project/entities"
	"tecnotree-go-project/interfaces"
	"tecnotree-go-project/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	User *mongo.Collection
}

func InitUserService(collection *mongo.Collection) interfaces.IUser {
	return &UserService{User: collection}
}

func (u *UserService) Register(user *entities.Register) (string, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Email = strings.ToLower(user.Email)

	if user.Password != user.ConfirmPassword {
		return "Password not matched", nil
	}

	if hashPassword := utils.EncryptPassword(user); hashPassword != nil {
		user.Password = string(hashPassword)
		user.ConfirmPassword = string(hashPassword)
	} else {
		return "Error in Password Encryption", nil
	}
	_, err := u.User.InsertOne(context.Background(), user)

	if err != nil {
		return "", err
	} else {
		return "User registered Successsfully", nil
	}
}

func (u *UserService) Login(user *entities.Login) (string, error) {
	if user.Email != "" && user.Password != "" {
		ctx := context.Background()
		query := bson.M{"Email": strings.ToLower(user.Email)}
		var loginResult *entities.Register
		err := u.User.FindOne(ctx, query).Decode(&loginResult)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
		err2 := utils.VerifyPassword(loginResult.Password, user)
		if err2 != nil {
			return "", err2
		}
		return "Login Successful", nil
	} else {
		return "Missing required fields", nil
	}
}

func (u *UserService) Logout() string {
	return "Logout Successful"
}

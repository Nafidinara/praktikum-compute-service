package services

import (
	"Praktikum/database"
	"Praktikum/middlewares"
	"Praktikum/models"
	"Praktikum/utils"
)

type UserService struct {
	jwtAuth *middlewares.JWTConfig
}

func InitUserService(jwtAuth *middlewares.JWTConfig) UserService {
	return UserService{
		jwtAuth: jwtAuth,
	}
}

func (service *UserService) Register(input models.RegisterInput) (models.User, error) {
	config := &utils.ArgonConfig{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		KeyLength:   32,
		SaltLength:  16,
	}

	password, err := utils.CreatePassword(input.Password, config)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	err = result.Last(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (service *UserService) Login(input models.LoginInput) (string, error) {
	var user models.User

	err := database.DB.First(&user, "email = ?", input.Email).Error

	if err != nil {
		return "", err
	}

	match, err := utils.ComparePasswordHash(input.Password, user.Password)

	isFailed := err != nil || !match

	if isFailed {
		return "", err
	}

	token, err := service.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserService) GetUser(id string) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
